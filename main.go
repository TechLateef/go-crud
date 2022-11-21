package main

import (
	"github.com/techlateef/tech-lateef-gol/config"
	"github.com/techlateef/tech-lateef-gol/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabasec(db)

	routes.Routes()

}
