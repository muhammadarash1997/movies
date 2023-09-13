package rpc

import (
	"context"

	"github.com/muhammadarash1997/movies/endpoint"
	pb "github.com/muhammadarash1997/movies/proto"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type grpcMoviesServer struct {
	getMovieList   grpctransport.Handler
	getMovieDetail grpctransport.Handler
	addMovie       grpctransport.Handler
	updateMovie    grpctransport.Handler
	deleteMovie    grpctransport.Handler
}

func NewGRPCMoviesServer(endpoints endpoint.MoviesEndpoint, logger log.Logger) pb.MoviesServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcMoviesServer{
		getMovieList: grpctransport.NewServer(
			endpoints.GetMovieListEndpoint,
			decodeEmptyRequest,
			encodeGetMovieListResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		getMovieDetail: grpctransport.NewServer(
			endpoints.GetMovieDetailEndpoint,
			decodeGetMovieDetailRequest,
			encodeGetMovieDetailResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		addMovie: grpctransport.NewServer(
			endpoints.AddMovieEndpoint,
			decodeAddMovieRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		updateMovie: grpctransport.NewServer(
			endpoints.UpdateMovieEndpoint,
			decodeUpdateMovieRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		deleteMovie: grpctransport.NewServer(
			endpoints.DeleteMovieEndpoint,
			decodeDeleteMovieRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
	}
}

// for empty
func encodeEmptyRequest(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeEmptyRequest(_ context.Context, response interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func decodeEmptyResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}
