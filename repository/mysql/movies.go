package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/muhammadarash1997/movies/model"
)

const (
	insertMovie = `
		INSERT INTO movies(
			id, title, description, rating, image, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	selectAllMovies = `
		SELECT id, title, description, rating, image, created_at, updated_at
		FROM movies
	`
	selectMovie = `
		SELECT id, title, description, rating, image, created_at, updated_at
		FROM movies
		WHERE id = ?
	`
	checkMovieExist = `
		SELECT id
		FROM movies
		WHERE id = ?
	`
	updateMovie = `
		UPDATE movies
		SET title = ?, description = ?, rating = ?, image = ?, updated_at = ?
		WHERE id = ?
	`
	deleteMovie = `
		DELETE FROM movies
		WHERE id = ?
	`
)

func (rw *dbReadWriter) WriteMovie(movie model.Movie) error {
	const (
		fileName = `movies.go`
		funcName = `WriteMovie`
		dbName   = `movie`
		sqlName  = `insertMovie`
	)

	timeNow := time.Now()
	movie.CreatedAt = timeNow
	movie.UpdatedAt = timeNow

	_, err := rw.db.Exec(insertMovie, movie.ID, movie.Title, movie.Description, movie.Rating, movie.Image, movie.CreatedAt, movie.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.Exec", err.Error())
	}

	return nil
}

func (rw *dbReadWriter) ReadAllMovies() ([]model.Movie, error) {
	const (
		fileName = `movies.go`
		funcName = `ReadAllMovies`
		dbName   = `movie`
		sqlName  = `selectAllMovies`
	)

	var movies []model.Movie

	rows, err := rw.db.Query(selectAllMovies)
	if err == sql.ErrNoRows {
		return movies, nil
	}
	if err != nil {
		return movies, fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.Query", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		movie := model.Movie{}
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return movies, fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "rows.Scan", err.Error())
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (rw *dbReadWriter) ReadMovie(id int64) (model.Movie, error) {
	const (
		fileName = `movies.go`
		funcName = `ReadMovie`
		dbName   = `movie`
		sqlName  = `selectMovie`
	)

	var movie model.Movie
	
	err := rw.db.QueryRow(selectMovie, id).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)
	if err == sql.ErrNoRows {
		return movie, nil
	}
	if err != nil {
		return movie, fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.QueryRow", err.Error())
	}

	return movie, nil
}

func (rw *dbReadWriter) CheckMovieExist(id int64) (bool, error) {
	const (
		fileName = `movies.go`
		funcName = `CheckMovieExist`
		dbName   = `movie`
		sqlName  = `checkMovieExist`
	)

	var movie model.Movie

	err := rw.db.QueryRow(checkMovieExist, id).Scan(&movie.ID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.QueryRow", err.Error())
	}

	return true, nil
}

func (rw *dbReadWriter) UpdateMovie(movie model.Movie) error {
	const (
		fileName = `movies.go`
		funcName = `UpdateMovie`
		dbName   = `movie`
		sqlName  = `updateMovie`
	)

	movie.UpdatedAt = time.Now()

	_, err := rw.db.Exec(updateMovie, movie.Title, movie.Description, movie.Rating, movie.Image, movie.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.Exec", err.Error())
	}

	return nil
}

func (rw *dbReadWriter) DeleteMovie(id int64) error {
	const (
		fileName = `movies.go`
		funcName = `DeleteMovie`
		dbName   = `movie`
		sqlName  = `deleteMovie`
	)

	_, err := rw.db.Exec(deleteMovie, id)
	if err != nil {
		return fmt.Errorf("%s | %s | %s | %s", fileName, funcName, "db.Exec", err.Error())
	}

	return nil
}
