package main

import (
	"dynamodb_go/cmd/server/handler"
	"dynamodb_go/users"
	"dynamodb_go/util"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	dDB, err := util.InitDynamo()
	if err != nil {
		panic(err)
	}

	repo := users.NewDynamoRepository(dDB)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	ur := engine.Group("/api/v1/users")
	ur.POST("/", u.Store())
	ur.GET("/:id", u.GetOne())
	ur.DELETE("/:id", u.Delete())

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
