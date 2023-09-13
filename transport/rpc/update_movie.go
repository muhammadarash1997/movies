package rpc

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/muhammadarash1997/movies/model"
	pb "github.com/muhammadarash1997/movies/proto"
)

// Server
func (s *grpcMoviesServer) UpdateMovie(ctx context.Context, req *pb.Movie) (*empty.Empty, error) {
	_, res, err := s.updateMovie.ServeGRPC(ctx, req)
	if err != nil {
		return &empty.Empty{}, err
	}
	return res.(*empty.Empty), nil
}

func decodeUpdateMovieRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Movie)
	return model.Movie{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Unix(req.CreatedAt, 0),
		UpdatedAt:   time.Unix(req.UpdatedAt, 0),
	}, nil
}
