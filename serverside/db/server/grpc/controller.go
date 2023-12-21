package grpc

import (
	"distributed-kv-db/common/grpcutil"
	"google.golang.org/grpc"
)

type RunningServer struct {
	port   int
	done   <-chan error
	server *grpc.Server
}

func newRunningServerFromGrpcServer(server grpcutil.StartedServer) RunningServer {
	return RunningServer{
		port:   server.Port(),
		done:   server.Done(),
		server: server.Server(),
	}
}

func (c RunningServer) Port() int {
	return c.port
}

func (c RunningServer) Done() <-chan error {
	return c.done
}

func (c RunningServer) ForceStop() {
	c.server.GracefulStop()
}
