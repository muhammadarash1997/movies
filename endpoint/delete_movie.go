package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/movies/service"
)

// DeleteMovie Endpoint
func makeDeleteMovieEndpoint(service service.MoviesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, service.DeleteMovie(ctx, request.(int64))
	}
}
