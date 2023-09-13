package rpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/muhammadarash1997/movies/proto"
)

// Server
func (s *grpcMoviesServer) DeleteMovie(ctx context.Context, req *pb.Request) (*empty.Empty, error) {
	_, res, err := s.deleteMovie.ServeGRPC(ctx, req)
	if err != nil {
		return &empty.Empty{}, err
	}
	return res.(*empty.Empty), nil
}

func decodeDeleteMovieRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Request)
	return req.Id, nil
}
