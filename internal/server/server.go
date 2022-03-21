package server

import (
	"errors"
	"github.com/go-pg/pg/v10"
	handler2 "github.com/rowdyroad/grpc-hw/internal/server/handler"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"github.com/rowdyroad/grpc-hw/internal/storage/csv"
	db "github.com/rowdyroad/grpc-hw/internal/storage/db"
	"net"
)
import "google.golang.org/grpc"
import proto "github.com/rowdyroad/grpc-hw/internal/grpc"

type Config struct {
	Listen string
	CSV *string
	DB *pg.Options
}

type Server struct {
	config Config
	listener net.Listener
	server *grpc.Server
}

func NewServer(config Config) (*Server, error) {
	var storage storage.IStorage
	var err error
	if config.CSV != nil {
		storage, err = csv.NewCSV(*config.CSV)
	} else if config.DB != nil {
		storage, err = db.NewDB(*config.DB)
	}
	if err != nil {
		return nil, err
	}
	if storage == nil {
		return nil, errors.New("incorrect config")
	}
	listener, err := net.Listen("tcp", config.Listen)
	if err != nil {
		return nil, err
	}
	server := grpc.NewServer()
	proto.RegisterStorageServer(server, handler2.NewHandler(storage))
	return &Server{
		config: config,
		listener: listener,
		server: server,
	},nil
}

func (s *Server) Run() error {
	if err := s.server.Serve(s.listener); err != nil && err != net.ErrClosed {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	return s.listener.Close()
}