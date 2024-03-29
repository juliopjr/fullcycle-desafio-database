package database

import (
	"log"

	"github.com/juliopjr/fullcycle-desafio-database/server/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	log.Println("-> Iniciando BD")
	db, err := gorm.Open(sqlite.Open("quotations.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("\tRealizando migrations")
	db.AutoMigrate(&entity.Quotation{})
	log.Println("\tBD pronto")
	return db
}
