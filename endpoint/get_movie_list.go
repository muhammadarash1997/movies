package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/movies/service"
)

// GetMovieList Endpoint
func makeGetMovieListEndpoint(service service.MoviesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMovieList(ctx)
	}
}
