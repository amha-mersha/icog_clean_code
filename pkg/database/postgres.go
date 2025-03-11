package database

import (
	"fmt"
	"log"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(host, user, password, dbname, port string) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&domain.TaskItem{})
	return db
}
