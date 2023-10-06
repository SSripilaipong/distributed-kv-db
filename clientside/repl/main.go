package repl

import "log/slog"

func New() Func {
	return func(serverIp string) error {

		slog.Info("connecting",
			"server-ip", serverIp,
		)
		return nil
	}
}
