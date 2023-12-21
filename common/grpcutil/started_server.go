package grpcutil

import (
	"google.golang.org/grpc"
)

type StartedServer struct {
	listen ListenInfo
	server *grpc.Server
	done   <-chan error
}

func (s StartedServer) Server() *grpc.Server {
	return s.server
}

func (s StartedServer) Port() int {
	return s.listen.Port
}

func (s StartedServer) Done() <-chan error {
	return s.done
}
