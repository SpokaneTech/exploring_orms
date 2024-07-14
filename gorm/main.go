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
	Name string
}

type Vehicle struct {
	gorm.Model
	Name string
}

func NewCli(db *gorm.DB) *cli.App {
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
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Manufacturer{}, &Vehicle{})

	cli := NewCli(db)
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
