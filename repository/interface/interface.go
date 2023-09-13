package _interface

import (
	"io"

	"github.com/muhammadarash1997/movies/model"
)

type ReadWriter interface {
	io.Closer
	ReadAllMovies() ([]model.Movie, error)
	ReadMovie(int64) (model.Movie, error)
	CheckMovieExist(int64) (bool, error)
	WriteMovie(model.Movie) error
	UpdateMovie(model.Movie) error
	DeleteMovie(int64) error
}
