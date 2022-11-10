package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/routes"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/util"
	"log"
)

func main() {
	engine := gin.Default()
	dataBase, errInit := util.InitDynamo()
	if errInit != nil {
		log.Fatal(errInit)
	}
	router := routes.NewRouter(engine, dataBase)
	router.MapRoutes()
	if errRun := engine.Run(); errRun != nil {
		panic(errRun)
	}
}
