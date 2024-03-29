package grpc

import (
	grpc2 "distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/usecase/getvalue"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
	"google.golang.org/grpc"
)

var New = newWithDeps(
	grpcutil.ListenToPort,
	registerWithImplServerFunc,
	grpcutil.StartServerFunc,
	newRunningServerFromGrpcServer,
)

func newWithDeps(
	listenToPort func(int) rslt.Of[grpcutil.ListenInfo],
	registerImplFunc func(grpc2.ServerServer) func(*grpc.Server),
	startServerFunc func(func(*grpc.Server)) func(grpcutil.ListenInfo) rslt.Of[grpcutil.StartedServer],
	newResultFromGrpcServer func(grpcutil.StartedServer) RunningServer,
) func(getValue getvalue.Func, setValue setvalue.Func) Func {

	startServerFuncFromImpl := fn.Compose(startServerFunc, registerImplFunc)
	newResult := rslt.Fmap(newResultFromGrpcServer)

	return func(getValue getvalue.Func, setValue setvalue.Func) Func {

		startServer := rslt.FmapPartial(startServerFuncFromImpl(grpcImpl{
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
