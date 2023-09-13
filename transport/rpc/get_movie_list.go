package rpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/muhammadarash1997/movies/model"
	pb "github.com/muhammadarash1997/movies/proto"
)

// Server
func (s *grpcMoviesServer) GetMovieList(ctx context.Context, req *empty.Empty) (*pb.GetMovieListResponse, error) {
	_, res, err := s.getMovieList.ServeGRPC(ctx, req)
	if err != nil {
		return &pb.GetMovieListResponse{}, err
	}
	return res.(*pb.GetMovieListResponse), nil
}

func encodeGetMovieListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.([]model.Movie)

	movies := []*pb.Movie{}
	for _, v := range resp {
		movies = append(movies, &pb.Movie{
			Id:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Rating:      v.Rating,
			Image:       v.Image,
			CreatedAt:   v.CreatedAt.Unix(),
			UpdatedAt:   v.UpdatedAt.Unix(),
		})
	}
	return &pb.GetMovieListResponse{
		Movies: movies,
	}, nil
}
