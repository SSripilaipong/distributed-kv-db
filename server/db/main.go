package db

import (
	"distributed-kv-db/server/db/server"
	"log/slog"
)

func New() Func {
	return build(server.New())

}

func build(serverFunc server.Func) Func {
	return func(dbPort, peeringPort int, advertisedIp string, peerAddresses []string) error {
		slog.Info("starting server",
			"db-port", dbPort,
			"peer-port", peeringPort,
			"advertised-ip", advertisedIp,
			"peers", peerAddresses,
		)
		return serverFunc(dbPort)
	}
}
