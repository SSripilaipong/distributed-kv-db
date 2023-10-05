package client

import (
	clientCli "distributed-kv-db/client/cli"
	"distributed-kv-db/client/repl"
	"github.com/urfave/cli/v2"
)

func NewCli() []*cli.Command {
	return clientCli.New(repl.New())
}
