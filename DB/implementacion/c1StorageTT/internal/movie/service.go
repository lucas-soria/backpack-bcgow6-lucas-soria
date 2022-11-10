package movie

import (
	"context"
	"errors"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	Save(ctx context.Context, movie domain.Movie) (domain.Movie, error)
	Update(ctx context.Context, movie domain.Movie, id int) (domain.Movie, error)
	Delete(ctx context.Context, id int) error
	GetByTitle(ctx context.Context, title string) ([]domain.Movie, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) (movies []domain.Movie, errGetAll error) {
	movies, errGetAll = s.repo.GetAll(ctx)
	return
}

func (s *service) Get(ctx context.Context, id int) (movie domain.Movie, errGet error) {
	movie, errGet = s.repo.Get(ctx, id)
	return
}

func (s *service) GetByTitle(ctx context.Context, title string) (movies []domain.Movie, errGetByName error) {
	movies, errGetByName = s.repo.GetByTitle(ctx, title)
	return
}

func (s *service) Save(ctx context.Context, movie domain.Movie) (savedMovie domain.Movie, errSave error) {
	if s.repo.Exists(ctx, movie.ID) {
		return domain.Movie{}, errors.New("error: movie id already exists")
	}
	movieID, errSave := s.repo.Save(ctx, movie)
	if errSave != nil {
		return
	}
	savedMovie = movie
	savedMovie.ID = movieID
	return
}

func (s *service) Update(ctx context.Context, movie domain.Movie, id int) (movieUpdated domain.Movie, errUpdate error) {
	errUpdate = s.repo.Update(ctx, movie, id)
	if errUpdate != nil {
		return
	}
	movieUpdated, errUpdate = s.repo.Get(ctx, id)
	return
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
