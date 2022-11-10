package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/movie"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	movieRepository := movie.NewRepository(r.db)
	movieService := movie.NewService(movieRepository)
	movieHandler := handler.NewMovie(movieService)
	r.rg.GET("/movies", movieHandler.GetAll())
	r.rg.GET("/movies/:id", movieHandler.Get())
	r.rg.GET("/movies/title/:title", movieHandler.GetByTitle())
	r.rg.POST("/movies", movieHandler.Create())
	r.rg.DELETE("/movies/:id", movieHandler.Delete())
	r.rg.PUT("/movies/:id", movieHandler.Update())
}
