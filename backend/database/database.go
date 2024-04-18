package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() error {
	var err error
	p := config.Config("PSQL_DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("PSQL_DB_HOST"), port, config.Config("PSQL_DB_USER"), config.Config("PSQL_DB_PASSWORD"), config.Config("PSQL_DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return err
	}

	fmt.Println("Postgres Connnected!")
	db.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{}, &models.Option{})

	DB = Dbinstance{
		Db: db,
	}
	return nil
}
