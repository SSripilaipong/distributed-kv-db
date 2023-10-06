package db

import "log/slog"

func logStartDb(dbPort int, peeringPort int, advertisedIp string, peerAddresses []string) {
	slog.Info("starting db",
		"db-port", dbPort,
		"peer-port", peeringPort,
		"advertised-ip", advertisedIp,
		"peers", peerAddresses,
	)
}
