package movie

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"log"
)

type Repository interface {
	Exists(ctx context.Context, id int) bool
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	Save(ctx context.Context, movie domain.Movie) (int, error)
	Update(ctx context.Context, movie domain.Movie, id int) error
	Delete(ctx context.Context, id int) error
	GetByTitle(ctx context.Context, title string) ([]domain.Movie, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	SaveMovie = "INSERT INTO movies (title, rating, awards, length, genre_id, release_date) VALUES (?, ?, ?, ?, ?, ?);"

	GetAllMovies = "SELECT id, title, rating, awards, length, genre_id FROM movies;"

	GetMovie = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE id = ?;"

	UpdateMovie = "UPDATE movies SET title = ?, rating = ?, awards = ?, length = ?, genre_id = ? WHERE id = ?;"

	DeleteMovie = "DELETE FROM movies WHERE id = ?;"

	ExistMovie = "SELECT id FROM movies WHERE id = ?;"

	GetMoviesByTitle = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE title = ?;"
)

func CloseStmt(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(ExistMovie, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetAll(ctx context.Context) (movies []domain.Movie, err error) {
	rows, err := r.db.QueryContext(ctx, GetAllMovies)
	if err != nil {
		return
	}
	for rows.Next() {
		var movie domain.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.GenreID)
		if err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return
}

func (r *repository) Get(ctx context.Context, id int) (domain.Movie, error) {
	row := r.db.QueryRow(GetMovie, id)
	var movie domain.Movie
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.GenreID); err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (r *repository) GetByTitle(ctx context.Context, title string) (movies []domain.Movie, err error) {
	rows, err := r.db.Query(GetMoviesByTitle, title)
	if err != nil {
		return
	}
	for rows.Next() {
		var movie domain.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.GenreID)
		if err != nil {
			return
		}
		movies = append(movies, movie)
	}
	return
}

func (r *repository) Save(ctx context.Context, movie domain.Movie) (int, error) {
	stmt, err := r.db.Prepare(SaveMovie)
	if err != nil {
		return 0, err
	}
	defer CloseStmt(stmt)
	result, err := stmt.Exec(movie.Title, movie.Rating, movie.Awards, movie.Length, movie.GenreID, movie.ReleaseDate)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) Update(ctx context.Context, movie domain.Movie, id int) error {
	stmt, err := r.db.Prepare(UpdateMovie)
	if err != nil {
		return err
	}
	defer CloseStmt(stmt)
	result, err := stmt.Exec(movie.Title, movie.Rating, movie.Awards, movie.Length, movie.GenreID, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("no affected rows")
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, DeleteMovie)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("no affected rows")
	}
	return nil
}
