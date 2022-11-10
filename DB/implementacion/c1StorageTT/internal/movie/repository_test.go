package movie

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

var nilInt *int

var movieTest = domain.Movie{
	ID:      1,
	Title:   "Cars 1",
	Rating:  4,
	Awards:  2,
	Length:  nilInt,
	GenreID: nilInt,
}

/*
var movieTest = domain.Movie{
	ID:          1,
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	Title:       "Cars 1",
	Rating:      4,
	Awards:      2,
	ReleaseDate: time.Layout,
	Length:      nilInt,
	GenreID:     nilInt,
}
*/

var query = "DELETE FROM DUAL WHERE id = 1;"

/* TESTS USANDO SQLMOCK */

func TestRepository_GetAll_OK(t *testing.T) {
	db, mock, errSql := sqlmock.New()
	defer db.Close()
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(movieTest.ID, movieTest.Title, movieTest.Rating, movieTest.Awards, movieTest.Length, movieTest.GenreID)
	mock.ExpectQuery(regexp.QuoteMeta(GetAllMovies)).WillReturnRows(rows)
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	movieResult, errGetAll := repo.GetAll(ctx)
	assert.NoError(t, errSql)
	assert.NoError(t, errGetAll)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, []domain.Movie{movieTest}, movieResult)
}

func TestRepository_GetAll_TimeOut(t *testing.T) {
	db, _, errSql := sqlmock.New()
	defer db.Close()
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	moviesResult, errGetAll := repo.GetAll(ctx)
	assert.NoError(t, errSql)
	assert.Error(t, errGetAll, context.DeadlineExceeded)
	assert.Nil(t, moviesResult)
}

func TestRepository_Delete_OK(t *testing.T) {
	searchID := 1
	db, mock, errSql := sqlmock.New()
	defer db.Close()
	repo := NewRepository(db)
	mock.ExpectPrepare(regexp.QuoteMeta(DeleteMovie))
	mock.ExpectExec(regexp.QuoteMeta(DeleteMovie)).WithArgs(searchID).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errDelete := repo.Delete(ctx, searchID)
	assert.NoError(t, errSql)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NoError(t, errDelete)
}

func TestRepository_Delete_TimeOut(t *testing.T) {
	searchID := 1
	db, mock, errSql := sqlmock.New()
	defer db.Close()
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	mock.ExpectPrepare(regexp.QuoteMeta(DeleteMovie)).ExpectExec().WithArgs(searchID).WillDelayFor(2 * time.Millisecond)
	errDelete := repo.Delete(ctx, searchID)
	assert.NoError(t, errSql)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Error(t, errDelete, context.DeadlineExceeded)
}
