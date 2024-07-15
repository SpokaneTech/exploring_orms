package main

import (
	"github.com/SpokaneTech/exploring_orms/pkg/models"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.ApplyBasic(
		&models.Manufacturer{},
		&models.Manufacturer{},
		&models.Model{},
		&models.Vehicle{},
		&models.Part{},
		&models.Person{},
	)

	// Generate the code
	g.Execute()
}
