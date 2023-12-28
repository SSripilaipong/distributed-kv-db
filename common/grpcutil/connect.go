package grpcutil

import (
	"distributed-kv-db/common/rslt"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(address string) rslt.Of[*grpc.ClientConn] {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return rslt.Error[*grpc.ClientConn](err)
	}
	return rslt.Value(conn)
}

func LocalAddress(port int) string {
	return fmt.Sprintf(":%d", port)
}
