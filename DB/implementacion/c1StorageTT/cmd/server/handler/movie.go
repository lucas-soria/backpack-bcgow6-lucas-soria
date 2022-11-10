package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/movie"
	"net/http"
	"strconv"
)

type Movie struct {
	service movie.Service
}

func NewMovie(service movie.Service) *Movie {
	return &Movie{
		service: service,
	}
}

func (m *Movie) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movies, errGetAll := m.service.GetAll(ctx)
		if errGetAll != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errGetAll.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusOK, movies)
	}
}

func (m *Movie) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, errStrConv := strconv.Atoi(ctx.Param("id"))
		if errStrConv != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": errStrConv.Error(),
				},
			)
		}
		movieObtained, errGet := m.service.Get(ctx, id)
		if errGet != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errGet.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusOK, movieObtained)
	}
}

func (m *Movie) GetByTitle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		moviesObtained, errGetByTitle := m.service.GetByTitle(ctx, title)
		if errGetByTitle != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errGetByTitle.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusOK, moviesObtained)
	}
}

func (m *Movie) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var moviePOST domain.Movie
		errBind := ctx.ShouldBindJSON(&moviePOST)
		if errBind != nil {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"error": errBind.Error(),
				},
			)
		}
		movieSaved, errSave := m.service.Save(ctx, moviePOST)
		if errSave != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errSave.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusCreated, movieSaved)
	}
}

func (m *Movie) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, errStrConv := strconv.Atoi(ctx.Param("id"))
		if errStrConv != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "invalid ID",
				},
			)
			return
		}
		var moviePUT domain.Movie
		errBind := ctx.ShouldBindJSON(&moviePUT)
		if errBind != nil {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"error": errBind.Error(),
				},
			)
			return
		}
		movieUpdated, errUpdate := m.service.Update(ctx, moviePUT, id)
		if errUpdate != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errUpdate.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusOK, movieUpdated)
	}
}

func (m *Movie) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, errStrConv := strconv.Atoi(ctx.Param("id"))
		if errStrConv != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": errStrConv.Error(),
				},
			)
			return
		}
		errDelete := m.service.Delete(ctx, id)
		if errDelete != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": errDelete.Error(),
				},
			)
			return
		}
		ctx.JSON(http.StatusNoContent, "")
	}
}
