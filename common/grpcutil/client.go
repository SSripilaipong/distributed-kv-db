package grpcutil

import "google.golang.org/grpc"

func ClientToInterface(conn *grpc.ClientConn) grpc.ClientConnInterface {
	return conn
}

func CloseClient(conn *grpc.ClientConn) error {
	return conn.Close()
}
