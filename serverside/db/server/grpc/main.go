package grpc

import (
	grpc2 "distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
	"google.golang.org/grpc"
)

var New = newWithDeps(
	grpcutil.ListenToPort,
	registerWithImplServerFunc,
	grpcutil.StartServerFunc,
	newRunningServerFromGrpcServer,
)

func newWithDeps(
	listenToPort func(int) result.Of[grpcutil.ListenInfo],
	registerImplFunc func(grpc2.ServerServer) func(*grpc.Server),
	startServerFunc func(func(*grpc.Server)) func(grpcutil.ListenInfo) result.Of[grpcutil.StartedServer],
	newResultFromGrpcServer func(grpcutil.StartedServer) RunningServer,
) func(getValue coordinator.GetValueFunc, setValue coordinator.SetValueFunc) Func {

	startServerFuncFromImpl := fn.Compose(startServerFunc, registerImplFunc)
	newResult := result.Fmap(newResultFromGrpcServer)

	return func(getValue coordinator.GetValueFunc, setValue coordinator.SetValueFunc) Func {

		startServer := result.FmapPartial(startServerFuncFromImpl(grpcImpl{
			getValue: getValue,
			setValue: setValue,
		}))

		return fn.Compose3(newResult, startServer, listenToPort)
	}
}

func registerWithImplServerFunc(impl grpc2.ServerServer) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		grpc2.RegisterServerServer(server, impl)
	}
}
