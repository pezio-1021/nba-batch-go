package main

import (
	"log"
	"nba-batch-go/pkg/config"
	"nba-batch-go/pkg/handlers"
	"nba-batch-go/service/interfaces"
	"nba-batch-go/service/usecase"
)

func InitializeController() *interfaces.Controller {
	cnf := config.Conf
	d, err := handlers.NewSQLHandler(cnf.Databases.NbaApp)
	if err != nil {
		log.Fatal(err)
	}

	playersRepo := interfaces.NewPlayerRepository(d)
	u := usecase.new(type)
}
