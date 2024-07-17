package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SpokaneTech/exploring_orms/internal/query"
	"github.com/SpokaneTech/exploring_orms/pkg/models"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	RecordNotFound string = "record not found"
)

func NewCli(q *query.Query) *cli.App {
	var (
		manufacturerName string
		modelName        string
		partName         string
		partCostStr      string
		vehicleVin       string
		personName       string
		vehicleIDStr     string
	)

	return &cli.App{
		Name:  "garage",
		Usage: "Manage your garage",
		Action: func(*cli.Context) error {
			var vehicles []*models.Vehicle
			var err error
			if vehicles, err = q.Vehicle.
				Joins(q.Vehicle.VehicleModel).
				Joins(q.Vehicle.VehicleModel.Manufacturer).
				Find(); err != nil {
				return err
			}
			if len(vehicles) == 0 {
				fmt.Println("No vehicles found in your garage")
			}
			for _, vehicle := range vehicles {
				fmt.Printf(
					"(%v) %v %v\n",
					vehicle.ID,
					vehicle.VehicleModel.Manufacturer.Name,
					vehicle.VehicleModel.Name,
				)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a vehicle to your garage",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "manufacturer",
						Required:    true,
						Destination: &manufacturerName,
					},
					&cli.StringFlag{
						Name:        "model",
						Required:    true,
						Destination: &modelName,
					},
					&cli.StringFlag{
						Name:        "vin",
						Destination: &vehicleVin,
					},
				},
				Action: func(ctx *cli.Context) error {
					var manufacturer *models.Manufacturer
					var err error
					if manufacturer, err = q.Manufacturer.
						Where(q.Manufacturer.Name.Eq(manufacturerName)).
						Attrs(q.Manufacturer.Name.Value(manufacturerName)).
						FirstOrCreate(); err != nil {
						if err.Error() != RecordNotFound {
							return err
						}
					}

					var model *models.Model
					if model, err = q.Model.
						Where(q.Model.Name.Eq(modelName)).
						Attrs(q.Model.Name.Value(modelName)).
						Attrs(q.Model.ManufacturerID.Value(manufacturer.ID)).
						FirstOrCreate(); err != nil {
						if err.Error() != RecordNotFound {
							return err
						}
					}

					vehicle := &models.Vehicle{Vin: vehicleVin, VehicleModel: *model}
					if err := q.Vehicle.Save(vehicle); err != nil {
						return err
					}

					fmt.Printf("Added a %v %v to your garage with the VIN %v\n", manufacturer.Name, model.Name, vehicle.Vin)
					return nil
				},
			},
			{
				Name:  "add-part",
				Usage: "Add a vehicle part to your garage",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Required:    true,
						Destination: &partName,
					},
					&cli.StringFlag{
						Name:        "cost",
						Required:    true,
						Destination: &partCostStr,
					},
				},

				Action: func(ctx *cli.Context) error {
					partCost, err := decimal.NewFromString(partCostStr)
					if err != nil {
						return err
					}

					part := &models.Part{Name: partName, Cost: partCost}
					if err := q.Part.Save(part); err != nil {
						return err
					}

					fmt.Printf("Added a new part %v to your garage\n", part.Name)
					return nil
				},
			},
			{
				Name: "list-parts",
				Action: func(*cli.Context) error {
					parts, err := q.Part.Find()
					if err != nil {
						return err
					}
					if len(parts) == 0 {
						fmt.Println("No parts found in your garage")
					}
					for _, part := range parts {
						fmt.Printf("%v ($%v)\n", part.Name, part.Cost)
					}
					return nil
				},
			},
			{
				Name:  "add-driver",
				Usage: "Add a driver to a vehicle in your garage",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "vehicle_id",
						Required:    true,
						Destination: &vehicleIDStr,
					},
					&cli.StringFlag{
						Name:        "name",
						Required:    true,
						Destination: &personName,
					},
				},
				Action: func(ctx *cli.Context) error {
					vehicleID, err := strconv.Atoi(vehicleIDStr)
					if err != nil {
						return fmt.Errorf("could not parse vehicle ID from %v", vehicleIDStr)
					}

					var vehicle *models.Vehicle
					if vehicle, err = q.Vehicle.
						Joins(q.Vehicle.VehicleModel).
						Joins(q.Vehicle.VehicleModel.Manufacturer).
						Where(q.Vehicle.ID.Eq(uint(vehicleID))).
						First(); err != nil {
						return err
					}

					var person *models.Person
					if person, err = q.Person.
						Where(q.Person.Name.Eq(personName)).
						Attrs(q.Person.Name.Value(personName)).
						FirstOrCreate(); err != nil {
						return err
					}

					q.Person.Limit(10)

					vehicle.Person = person
					if err := q.Vehicle.Save(vehicle); err != nil {
						return err
					}

					fmt.Printf(
						"Added a %v as the owner of (%v) %v %v\n",
						person.Name,
						vehicle.ID,
						vehicle.VehicleModel.Name,
						vehicle.VehicleModel.Manufacturer.Name,
					)

					return nil
				},
			},
		},
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.Manufacturer{},
		&models.Model{},
		&models.Vehicle{},
		&models.Part{},
		&models.Person{},
	)

	query := query.Use(db)

	cli := NewCli(query)
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
