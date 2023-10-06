package db

import (
	"distributed-kv-db/serverside/db/server"
)

func New(interrupt func() <-chan struct{}) Func {
	return build(interrupt, server.New())
}

func build(interrupt func() <-chan struct{}, serverFunc server.Func) Func {
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
