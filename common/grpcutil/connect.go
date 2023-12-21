package grpcutil

import (
	"distributed-kv-db/common/result"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(address string) result.Of[*grpc.ClientConn] {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return result.Error[*grpc.ClientConn](err)
	}
	return result.Value(conn)
}

func LocalAddress(port int) string {
	return fmt.Sprintf(":%d", port)
}
