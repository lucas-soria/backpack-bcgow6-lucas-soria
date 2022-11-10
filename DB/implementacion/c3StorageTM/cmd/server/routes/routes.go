package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/user"
)

type Router interface {
	MapRoutes()
}

type router struct {
	en *gin.Engine
	rg *gin.RouterGroup
	db *dynamodb.DynamoDB
}

func NewRouter(en *gin.Engine, db *dynamodb.DynamoDB) Router {
	return &router{en: en, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.en.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	userRepository := user.NewDynamoRepository(r.db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUser(userService)
	r.rg.GET("/users/:id", userHandler.Get())
	r.rg.GET("/users/", userHandler.GetAll())
	r.rg.POST("/users/", userHandler.Store())
	r.rg.PUT("/users/:id", userHandler.Update())
	r.rg.DELETE("/users/:id", userHandler.Delete())
}
