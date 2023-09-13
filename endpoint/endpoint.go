package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	svc "github.com/muhammadarash1997/movies/service"
)

// MoviesEndpoint
type MoviesEndpoint struct {
	GetMovieListEndpoint   endpoint.Endpoint
	GetMovieDetailEndpoint endpoint.Endpoint
	AddMovieEndpoint       endpoint.Endpoint
	UpdateMovieEndpoint    endpoint.Endpoint
	DeleteMovieEndpoint    endpoint.Endpoint
}

func NewMoviesEndpoint(service svc.MoviesService) (MoviesEndpoint, error) {
	var getMovieListEp endpoint.Endpoint
	{
		getMovieListEp = makeGetMovieListEndpoint(service)
	}
	var getMovieDetailEp endpoint.Endpoint
	{
		getMovieDetailEp = makeGetMovieDetailEndpoint(service)
	}
	var addMovieEp endpoint.Endpoint
	{
		addMovieEp = makeAddMovieEndpoint(service)
	}
	var updateMovieEp endpoint.Endpoint
	{
		updateMovieEp = makeUpdateMovieEndpoint(service)
	}
	var deleteMovieEp endpoint.Endpoint
	{
		deleteMovieEp = makeDeleteMovieEndpoint(service)
	}

	return MoviesEndpoint{
		GetMovieListEndpoint:   getMovieListEp,
		GetMovieDetailEndpoint: getMovieDetailEp,
		AddMovieEndpoint:       addMovieEp,
		UpdateMovieEndpoint:    updateMovieEp,
		DeleteMovieEndpoint:    deleteMovieEp,
	}, nil

}
