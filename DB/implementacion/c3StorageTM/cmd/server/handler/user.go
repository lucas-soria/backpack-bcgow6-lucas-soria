package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler/request"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/user"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/web"
	"net/http"
)

var (
	ErrInvalidID = errors.New("invalid id")
)

type User struct {
	service user.Service
}

func NewUser(service user.Service) *User {
	return &User{service: service}
}

func (user *User) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		userObtained, errGet := user.service.Get(ctx, id)
		if errGet != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, errGet.Error(), http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(userObtained, "", http.StatusOK))
	}
}

func (user *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usersObtained, errGet := user.service.GetAll(ctx)
		if errGet != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, errGet.Error(), http.StatusNotFound))
			return
		}
		if usersObtained == nil {
			ctx.JSON(http.StatusOK, web.NewResponse([]domain.User{}, "", http.StatusOK))
		}
		ctx.JSON(http.StatusOK, web.NewResponse(usersObtained, "", http.StatusOK))
	}
}

func (user *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequest request.UserRequest
		if errBind := ctx.ShouldBindJSON(&userRequest); errBind != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, errBind.Error(), http.StatusBadRequest))
			return
		}
		userMapped := userRequest.MapToDomain()
		errStore := user.service.Store(ctx, &userMapped)
		if errStore != nil {
			ctx.JSON(http.StatusConflict, web.NewResponse(nil, errStore.Error(), http.StatusConflict))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(userMapped, "", http.StatusCreated))
	}
}

func (user *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var userRequest request.UserRequest
		if errBind := ctx.ShouldBindJSON(&userRequest); errBind != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, errBind.Error(), http.StatusBadRequest))
			return
		}
		userMapped := userRequest.MapToDomain()
		errUpdate := user.service.Update(ctx, &userMapped, id)
		if errUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(nil, errUpdate.Error(), http.StatusInternalServerError))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(userMapped, "", http.StatusOK))
	}
}

func (user *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, ErrInvalidID.Error(), http.StatusBadRequest))
			return
		}
		errDelete := user.service.Delete(ctx, id)
		if errDelete != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, errDelete.Error(), http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(nil, "", http.StatusNoContent))
	}
}
