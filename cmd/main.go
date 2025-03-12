package main

import (
	"github.com/amha-mersha/icog_clean_code/config"
	delivery "github.com/amha-mersha/icog_clean_code/internal/delivery/http"
	"github.com/amha-mersha/icog_clean_code/pkg/database"
)

func main() {
	configs := config.LoadConfig("../.env")
	db := database.Init(configs.PostgresHost, configs.PostgresUser, configs.PostgresPassword, configs.PostgresDB, configs.PostgresPort)
	delivery.InitRouter(db, configs.APIVersion, configs.Port)
}
