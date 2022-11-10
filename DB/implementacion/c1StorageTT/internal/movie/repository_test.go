package movie

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/util"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

var nilInt *int

var movieTest = domain.Movie{
	ID:          1,
	Title:       "Cars 1",
	Rating:      4,
	Awards:      2,
	Length:      nilInt,
	GenreID:     nilInt,
	ReleaseDate: "2022-10-11 00:00:00",
}

/* TESTS USANDO TXDB */

func TestRepository_TXDB_Save_Ok(t *testing.T) {
	db, errInit := util.InitDb()
	assert.NoError(t, errInit)
	movieRepository := NewRepository(db)
	ctx := context.TODO()
	createdID, errSave := movieRepository.Save(ctx, movieTest)
	assert.NoError(t, errSave)
	movieTest.ID = createdID
	getResult, errGet := movieRepository.Get(ctx, createdID)
	assert.NoError(t, errGet)
	assert.NotNil(t, getResult)
	assert.Equal(t, movieTest, getResult)
}

func TestRepository_TXDB_Get_NotFound(t *testing.T) {
	db, errInit := util.InitDb()
	assert.NoError(t, errInit)
	movieRepository := NewRepository(db)
	ctx := context.TODO()
	searchID := 3493873
	getResult, errGet := movieRepository.Get(ctx, searchID)
	assert.NotNil(t, errGet)
	assert.Empty(t, getResult)
}

func TestRepository_TXDB_Update_OK(t *testing.T) {
	// Get a movie
	db, errInit := util.InitDb()
	assert.NoError(t, errInit)
	movieRepository := NewRepository(db)
	ctx := context.TODO()
	searchID := 3
	getResult, errGet := movieRepository.Get(ctx, searchID)
	assert.Nil(t, errGet)
	assert.NotEmpty(t, getResult)
	// Update the movie
	getResult.Title = "Updated title"
	errUpdate := movieRepository.Update(ctx, getResult, searchID)
	assert.Nil(t, errUpdate)
	// Check updated
	updateResult, errGet := movieRepository.Get(ctx, searchID)
	assert.Nil(t, errGet)
	assert.Equal(t, getResult, updateResult)
}

func TestRepository_TXDB_Delete_OK(t *testing.T) {
	db, errInit := util.InitDb()
	assert.NoError(t, errInit)
	movieRepository := NewRepository(db)
	ctx := context.TODO()
	// Get a movie
	searchID := 46 // Non bounded movie
	getResult, errGet := movieRepository.Get(ctx, searchID)
	assert.Nil(t, errGet)
	assert.NotEmpty(t, getResult)
	// Delete the movie
	errDelete := movieRepository.Delete(ctx, searchID)
	assert.Nil(t, errDelete)
	// Check deleted by id
	deleteResult, errGet := movieRepository.Get(ctx, searchID)
	assert.NotNil(t, errGet)
	assert.Empty(t, deleteResult)
	// Check deleted iterating all
	getAllResult, errGetAll := movieRepository.GetAll(ctx)
	assert.NoError(t, errGetAll)
	for _, movie := range getAllResult {
		assert.NotEqual(t, searchID, movie.ID)
	}
}

/* TESTS USANDO SQLMOCK */

func TestRepository_SQLMOCK_GetAll_OK(t *testing.T) {
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id", "release_date"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(movieTest.ID, movieTest.Title, movieTest.Rating, movieTest.Awards, movieTest.Length, movieTest.GenreID, movieTest.ReleaseDate)
	mock.ExpectQuery(regexp.QuoteMeta(GetAllMovies)).WillReturnRows(rows)
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	movieResult, errGetAll := repo.GetAll(ctx)
	assert.NoError(t, errGetAll)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, []domain.Movie{movieTest}, movieResult)
}

func TestRepository_SQLMOCK_GetAll_TimeOut(t *testing.T) {
	db, _, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	moviesResult, errGetAll := repo.GetAll(ctx)
	assert.Error(t, errGetAll, context.DeadlineExceeded)
	assert.Nil(t, moviesResult)
}

func TestRepository_SQLMOCK_Delete_OK(t *testing.T) {
	searchID := 1
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	repo := NewRepository(db)
	mock.ExpectPrepare(regexp.QuoteMeta(DeleteMovie)).ExpectExec().WithArgs(searchID).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errDelete := repo.Delete(ctx, searchID)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NoError(t, errDelete)
}

func TestRepository_SQLMOCK_Delete_TimeOut(t *testing.T) {
	searchID := 1
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	mock.ExpectPrepare(regexp.QuoteMeta(DeleteMovie)).ExpectExec().WithArgs(searchID).WillDelayFor(2 * time.Millisecond)
	errDelete := repo.Delete(ctx, searchID)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Error(t, errDelete, context.DeadlineExceeded)
}

// Replicación de los ejercicios de TXDB (Delete ya está hecho,)

func TestRepository_SQLMOCK_Save_OK(t *testing.T) {
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id", "release_date"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(movieTest.ID, movieTest.Title, movieTest.Rating, movieTest.Awards, movieTest.Length, movieTest.GenreID, movieTest.ReleaseDate)
	mock.ExpectPrepare(regexp.QuoteMeta(SaveMovie)).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	saveResultID, errGetAll := repo.Save(ctx, movieTest)
	assert.NoError(t, errGetAll)
	assert.NoError(t, mock.ExpectationsWereMet())
	mock.ExpectQuery(regexp.QuoteMeta(GetMovie)).WithArgs(saveResultID).WillReturnRows(rows)
	getResult, errGet := repo.Get(ctx, saveResultID)
	assert.NoError(t, errGet)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, movieTest, getResult)
}

func TestRepository_SQLMOCK_Get_NotFound(t *testing.T) {
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	searchID := 9739823
	mock.ExpectQuery(regexp.QuoteMeta(GetMovie)).WithArgs(searchID)
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	getResult, errGet := repo.Get(ctx, searchID)
	assert.Error(t, errGet)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Empty(t, getResult)
}

func TestRepository_SQLMOCK_Update_OK(t *testing.T) {
	db, mock, errSql := sqlmock.New()
	assert.NoError(t, errSql)
	defer db.Close()
	update := domain.Movie{
		Title:   "Updated",
		Rating:  10.0,
		Awards:  1,
		Length:  nilInt,
		GenreID: nilInt,
	}
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id", "release_date"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(movieTest.ID, update.Title, update.Rating, update.Awards, update.Length, update.GenreID, movieTest.ReleaseDate)
	mock.ExpectPrepare(regexp.QuoteMeta(UpdateMovie)).ExpectExec().WithArgs(update.Title, update.Rating, update.Awards, update.Length, update.GenreID, movieTest.ID).WillReturnResult(sqlmock.NewResult(int64(movieTest.ID), 1))
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errUpdate := repo.Update(ctx, update, movieTest.ID)
	assert.NoError(t, errUpdate)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(regexp.QuoteMeta(GetMovie)).WithArgs(movieTest.ID).WillReturnRows(rows)
	getResult, errGet := repo.Get(ctx, movieTest.ID)
	assert.NoError(t, errGet)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotEqual(t, movieTest, getResult)
}
