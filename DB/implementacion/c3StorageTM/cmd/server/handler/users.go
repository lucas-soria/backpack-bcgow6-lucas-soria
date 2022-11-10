package handler

import (
	"dynamodb_go/models"
	"dynamodb_go/pkg/web"
	"dynamodb_go/users"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type requestUser struct {
	Id         string `json:"id,omitempty"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	MacAddress string `json:"macAddress"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{service: s}
}

func (s *User) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestUser
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		user := models.User(req)
		err := s.service.Store(c, &user)
		if err != nil {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		c.JSON(http.StatusCreated, web.NewResponse(user, "", http.StatusCreated))
	}
}

func (s *User) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, errors.New("id es inválido").Error(), http.StatusBadRequest))
			return
		}

		user, err := s.service.GetOne(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(user, "", http.StatusOK))
	}
}

func (s *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, errors.New("id es inválido").Error(), http.StatusBadRequest))
			return
		}

		err := s.service.Delete(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusNoContent, web.NewResponse(nil, "", http.StatusNoContent))
	}
}
