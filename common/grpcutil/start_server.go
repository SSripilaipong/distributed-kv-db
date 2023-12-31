package grpcutil

import (
	"distributed-kv-db/common/rslt"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

type ListenInfo struct {
	Port     int
	listener net.Listener
}

func ListenToPort(port int) rslt.Of[ListenInfo] {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return rslt.Error[ListenInfo](fmt.Errorf("cannot listen to port: %w", err))
	}

	tokens := strings.Split(lis.Addr().String(), ":")
	port, err = strconv.Atoi(tokens[len(tokens)-1])
	if err != nil {
		return rslt.Error[ListenInfo](fmt.Errorf("not recognized port: %w", err))
	}
	return rslt.Value(ListenInfo{
		Port:     port,
		listener: lis,
	})
}

func StartServerFunc(register func(*grpc.Server)) func(listen ListenInfo) rslt.Of[StartedServer] {
	return func(listen ListenInfo) rslt.Of[StartedServer] {
		server := grpc.NewServer()
		register(server)

		done := make(chan error)
		go func() {
			defer close(done)
			if err := server.Serve(listen.listener); err != nil {
				done <- fmt.Errorf("failure occurred while serving grpc server: %w", err)
			} else {
				done <- nil
			}
		}()

		return rslt.Value(StartedServer{listen: listen, server: server, done: done})
	}
}
