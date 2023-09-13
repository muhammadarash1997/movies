package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/movies/service"
)

// GetMovieDetail Endpoint
func makeGetMovieDetailEndpoint(service service.MoviesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMovieDetail(ctx, request.(int64))
	}
}
