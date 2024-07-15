package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// A vehicle manufacturer, like Chevrolet
type Manufacturer struct {
	gorm.Model
	Name     string
	Vehicles []Model
}

// A vehicle model, like a Chevrolet Silverado
type Model struct {
	gorm.Model
	Name           string
	ManufacturerID uint
	Manufacturer   *Manufacturer
	Parts          []*Part `gorm:"many2many:vehicle_parts;"`
}

// An individual of a model, like Joe's Chevrolet Silverado
type Vehicle struct {
	gorm.Model
	Vin            string
	VehicleModelID uint
	VehicleModel   Model
	PersonID       *int
	Person         *Person
}

// A vehicle part for one or more models, like a muffler for all Chevrolet pickups
type Part struct {
	gorm.Model
	Name   string
	Cost   decimal.Decimal `gorm:"type:decimal(7,2);"`
	Models []*Model        `gorm:"many2many:model_parts;"`
}

// A person, who may drive a vehicle
type Person struct {
	gorm.Model
	Name string
}

var (
	RecordNotFound string = "record not found"
)

func NewCli(db *gorm.DB) *cli.App {
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
			vehicles := []Model{}
			result := db.Preload("Manufacturer").Find(&vehicles)
			if result.Error != nil {
				return result.Error
			}
			if len(vehicles) == 0 {
				fmt.Println("No vehicles found in your garage")
			}
			for _, vehicle := range vehicles {
				fmt.Printf("(%v) %v %v\n", vehicle.ID, vehicle.Manufacturer.Name, vehicle.Name)
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
					manufacturer := &Manufacturer{Name: manufacturerName}
					if result := db.Where("name = ?", manufacturerName).First(&manufacturer); result.Error != nil {
						if result.Error.Error() != RecordNotFound {
							return result.Error
						}
						if result = db.Save(manufacturer); result.Error != nil {
							return result.Error
						}
					}

					model := &Model{Name: modelName, Manufacturer: manufacturer}
					if result := db.Where("name = ?", modelName).First(&model); result.Error != nil {
						if result.Error.Error() != RecordNotFound {
							return result.Error
						}
						if result = db.Save(model); result.Error != nil {
							return result.Error
						}
					}

					vehicle := &Vehicle{Vin: vehicleVin, VehicleModel: *model}
					if result := db.Save(vehicle); result.Error != nil {
						return result.Error
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
					part := &Part{Name: partName, Cost: partCost}
					result := db.Save(part)
					if result.Error != nil {
						return result.Error
					}
					fmt.Printf("Added a new part %v to your garage\n", part.Name)
					return nil
				},
			},
			{
				Name: "list-parts",
				Action: func(*cli.Context) error {
					parts := []Part{}
					result := db.Find(&parts)
					if result.Error != nil {
						return result.Error
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
					vehicle := &Vehicle{}
					if result := db.
						Joins("VehicleModel").
						Joins("VehicleModel.Manufacturer").
						First(vehicle, vehicleID); result.Error != nil {
						return result.Error
					}

					person := &Person{Name: personName}
					if err := createOrUpdate(db, person, db.Where("name = ?", modelName)); err != nil {
						return err
					}
					vehicle.Person = person
					if result := db.Save(vehicle); result.Error != nil {
						return result.Error
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
		&Manufacturer{},
		&Model{},
		&Vehicle{},
		&Part{},
		&Person{},
	)

	cli := NewCli(db)
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createOrUpdate(db *gorm.DB, dest interface{}, query interface{}) error {
	if result := db.Where(query).First(&dest); result.Error != nil {
		if result.Error.Error() != RecordNotFound {
			return result.Error
		}
		if result = db.Save(dest); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
