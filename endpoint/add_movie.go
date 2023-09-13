package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/movies/model"
	"github.com/muhammadarash1997/movies/service"
)

// AddMovie Endpoint
func makeAddMovieEndpoint(service service.MoviesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.Movie)
		return nil, service.AddMovie(ctx, req)
	}
}
