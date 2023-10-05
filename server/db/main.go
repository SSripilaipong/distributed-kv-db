package db

import "log/slog"

func New() Func {
	return func(dbPort, peeringPort int, advertisedIp string, peerAddresses []string) error {
		slog.Info("starting server",
			"db-port", dbPort,
			"peer-port", peeringPort,
			"advertised-ip", advertisedIp,
			"peers", peerAddresses,
		)
		return nil
	}
}
