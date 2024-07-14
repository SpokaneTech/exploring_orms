package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manufacturer struct {
	gorm.Model
	Name     string
	Vehicles []Vehicle
}

type Vehicle struct {
	gorm.Model
	Name           string
	ManufacturerID uint
	Parts          []*Part `gorm:"many2many:vehicle_parts;"`
}

type Part struct {
	gorm.Model
	Name     string
	Cost     decimal.Decimal `gorm:"type:decimal(7,2);"`
	Vehicles []*Vehicle      `gorm:"many2many:vehicle_parts;"`
}

var (
	RecordNotFound string = "record not found"
)

func NewCli(db *gorm.DB) *cli.App {
	var (
		manufacturerName string
		vehicleName      string
		partName         string
		partCostStr      string
	)

	return &cli.App{
		Name:  "garage",
		Usage: "Manage your garage",
		Action: func(*cli.Context) error {
			vehicles := []Vehicle{}
			result := db.Find(&vehicles)
			if result.Error != nil {
				return result.Error
			}
			if len(vehicles) == 0 {
				fmt.Println("No vehicles found in your garage")
			}
			for _, vehicle := range vehicles {
				fmt.Printf("%v\n", vehicle.Name)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a car to your garage",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "manufacturer",
						Required:    true,
						Destination: &manufacturerName,
					},
					&cli.StringFlag{
						Name:        "name",
						Required:    true,
						Destination: &vehicleName,
					},
				},

				Action: func(ctx *cli.Context) error {
					manufacturer := &Manufacturer{}
					if result := db.Where("name = ?", manufacturerName).First(&manufacturer); result.Error != nil {
						if result.Error.Error() != RecordNotFound {
							return result.Error
						}
						manufacturer.Name = manufacturerName
						db.Save(manufacturer)
					}

					vehicle := &Vehicle{Name: vehicleName}
					result := db.Save(vehicle)
					if result.Error != nil {
						return result.Error
					}
					fmt.Printf("Added a %v %v to your garage\n", manufacturer.Name, vehicle.Name)
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
		&Vehicle{},
		&Manufacturer{},
		&Part{},
	)

	cli := NewCli(db)
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
