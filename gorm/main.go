package main

import (
	"fmt"
	"log"
	"os"

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
}

var (
	RecordNotFound string = "record not found"
)

func NewCli(db *gorm.DB) *cli.App {
	var manufacturerName string
	var vehicleName string

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
					fmt.Printf("Added %v to your garage\n", vehicle.Name)
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
	)

	cli := NewCli(db)
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
