package main

import (
	"abanku/configs"
	"abanku/infra"
	"abanku/service"
	"fmt"
)

func main() {
	fmt.Println("Start ABanku")

	// init config
	configs := configs.InitConfigs()

	// init infra
	db := infra.ConnectDB(&configs.DBConfig)
	defer db.Close()

	// init service
	services := service.NewService(db)
	infra.NewRest(services)
}
