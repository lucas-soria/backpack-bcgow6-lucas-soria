package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/router"
	"log"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if err := r.Run(":18085"); err != nil {
		log.Fatalf("Error running engine")
	}

}
