package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	ep "github.com/muhammadarash1997/movies/endpoint"
	pb "github.com/muhammadarash1997/movies/proto"
	"github.com/muhammadarash1997/movies/repository"
	svc "github.com/muhammadarash1997/movies/server"
	"github.com/muhammadarash1997/movies/transport/rpc"
	shv "github.com/muhammadarash1997/movies/utils/sharevar"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type RegisterHTTPHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

type Config struct {
	ServiceIP   string `json:"serviceip"`
	ServicePort string `json:"serviceport"`
	DBHost      string `json:"dbhost"`
	DBName      string `json:"dbname"`
	DBUser      string `json:"dbusr"`
	DBPwd       string `json:"dbpwd"`
}

type conf struct {
	IP      string
	Port    string
	Address string
	DBHost  string
	DBName  string
	DBUser  string
	DBPwd   string
}

func main() {
	shv.Logger = log.NewLogfmtLogger(os.Stdout)
	shv.Logger = log.With(shv.Logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	config, err := parseAndLoadConfigs()
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return
	}

	serviceServer, repo, err := config.moviesServer()
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return
	}
	defer repo.Close()

	grpcServer := config.initGRPCServer(serviceServer)
	go func() {
		ServeGRPCAndHTTP(config.Address, config.Port, grpcServer, pb.RegisterMoviesServiceHandlerFromEndpoint, shv.Logger)
	}()

	level.Info(shv.Logger).Log("[RUNNING]", "Server started")

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(shv.Logger).Log("exit", <-errs)
}

func parseAndLoadConfigs() (*conf, error) {
	godotenv.Load(".env")

	ip := os.Getenv("SERVICE_IP")
	port := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	address := fmt.Sprintf("%s:%s", ip, port)
	dbAddress := fmt.Sprintf("%s:%s", dbHost, dbPort)

	return &conf{
		IP:      ip,
		Port:    port,
		Address: address,
		DBHost:  dbAddress,
		DBName:  dbName,
		DBUser:  dbUser,
		DBPwd:   dbPwd,
	}, nil
}

func (cnf *conf) moviesServer() (*pb.MoviesServiceServer, *repository.Repo, error) {
	repoConf := repository.RepoConfigs{
		DBConf: repository.DBConf{
			URL:      cnf.DBHost,
			Schema:   cnf.DBName,
			User:     cnf.DBUser,
			Password: cnf.DBPwd,
		},
	}

	repo, err := repository.NewMoviesRepo(repoConf)
	if err != nil {
		return nil, nil, err
	}

	moviesService := svc.NewMoviesService(*repo)

	endpoint, err := ep.NewMoviesEndpoint(moviesService)
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return nil, nil, err
	}

	server := rpc.NewGRPCMoviesServer(endpoint, shv.Logger)

	return &server, repo, nil
}

func (cnf *conf) initGRPCServer(serviceServer *pb.MoviesServiceServer) *grpc.Server {
	var grpcServer *grpc.Server
	grpcServer = grpc.NewServer()
	pb.RegisterMoviesServiceServer(grpcServer, *serviceServer)
	return grpcServer
}

func ServeGRPCAndHTTP(address, port string, grpcServer *grpc.Server, register RegisterHTTPHandler, logger log.Logger) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	ctx := context.Background()
	mux := runtime.NewServeMux()

	err := register(ctx, mux, address, opts)
	if err != nil {
		logger.Log("[ERROR]", err.Error())
		return
	}

	var handler http.Handler
	handler = mux

	var httpServer *http.Server
	httpServer = &http.Server{
		Addr:    address,
		Handler: handler,
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Log("[ERROR]", err.Error())
		return
	}

	cMux := cmux.New(conn)
	grpcL := cMux.Match(cmux.HTTP2(), cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := cMux.Match(cmux.HTTP1Fast())

	go grpcServer.Serve(grpcL)
	go httpServer.Serve(httpL)

	cMux.Serve()
}
