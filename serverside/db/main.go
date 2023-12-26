package db

import (
	"distributed-kv-db/serverside/db/coordinator/getvalue"
	"distributed-kv-db/serverside/db/coordinator/setvalue"
	"distributed-kv-db/serverside/db/server"
)

func Builder(interrupt func() <-chan struct{}) Func {
	serverFunc := server.New(
		getvalue.New(nil),
		setvalue.New(),
	)
	return builder(interrupt, serverFunc)
}

func builder(interrupt func() <-chan struct{}, serverFunc server.Func) Func {
	return func(dbPort, peeringPort int, advertisedIp string, peerAddresses []string) error {
		logStartDb(dbPort, peeringPort, advertisedIp, peerAddresses)

		serverCtrl := serverFunc(dbPort)

		select {
		case err := <-serverCtrl.Error():
			return err
		case <-interrupt():
			serverCtrl.ForceStop()
		}
		return nil
	}
}
