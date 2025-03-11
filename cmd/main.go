package main

import (
	"github.com/amha-mersha/icog_clean_code/config"
	"github.com/amha-mersha/icog_clean_code/pkg/database"
)

func main() {
	configs := config.LoadConfig("../.env")
	database.Init(configs.PostgresHost, configs.PostgresUser, configs.PostgresPassword, configs.PostgresDB, configs.PostgresPort)
}
