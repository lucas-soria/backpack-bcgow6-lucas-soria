package main

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/routes"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/db"
	"log"
)

func main() {
	engine, dataBase := db.ConnectDatabase()
	router := routes.NewRouter(engine, dataBase)
	router.MapRoutes()
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Couldn't run engine")
	}
}
