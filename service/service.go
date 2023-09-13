package service

import (
	"context"

	"github.com/muhammadarash1997/movies/model"
)

type MoviesService interface {
	GetMovieList(context.Context) ([]model.Movie, error)
	GetMovieDetail(context.Context, int64) (model.Movie, error)
	AddMovie(context.Context, model.Movie) error
	UpdateMovie(context.Context, model.Movie) error
	DeleteMovie(context.Context, int64) error
}
