package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Service interface {
	Get(ctx *gin.Context, id string) (*domain.User, error)
	GetAll(ctx *gin.Context) ([]*domain.User, error)
	Store(ctx *gin.Context, user *domain.User) error
	Update(ctx *gin.Context, user *domain.User, id string) error
	Delete(ctx *gin.Context, id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service *service) Get(ctx *gin.Context, id string) (*domain.User, error) {
	return service.repository.Get(ctx, id)
}

func (service *service) GetAll(ctx *gin.Context) ([]*domain.User, error) {
	return service.repository.GetAll(ctx)
}

func (service *service) Store(ctx *gin.Context, user *domain.User) error {
	return service.repository.Store(ctx, user)
}

func (service *service) Update(ctx *gin.Context, user *domain.User, id string) error {
	userToUpdate, errGet := service.Get(ctx, id)
	if errGet != nil {
		return errGet
	}
	user.ID = userToUpdate.ID
	return service.repository.Update(ctx, user)
}

func (service *service) Delete(ctx *gin.Context, id string) error {
	return service.repository.Delete(ctx, id)
}
