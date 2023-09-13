package rpc

import (
	"context"

	"github.com/muhammadarash1997/movies/model"
	pb "github.com/muhammadarash1997/movies/proto"
)

// Server
func (s *grpcMoviesServer) GetMovieDetail(ctx context.Context, req *pb.Request) (*pb.Movie, error) {
	_, res, err := s.getMovieDetail.ServeGRPC(ctx, req)
	if err != nil {
		return &pb.Movie{}, err
	}
	return res.(*pb.Movie), nil
}

func decodeGetMovieDetailRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Request)
	return req.Id, nil
}

func encodeGetMovieDetailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(model.Movie)
	return &pb.Movie{
		Id:          resp.ID,
		Title:       resp.Title,
		Description: resp.Description,
		Rating:      resp.Rating,
		Image:       resp.Image,
		CreatedAt:   resp.CreatedAt.Unix(),
		UpdatedAt:   resp.UpdatedAt.Unix(),
	}, nil
}
