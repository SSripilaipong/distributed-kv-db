package clientside

import (
	clientCli "distributed-kv-db/clientside/cli"
	"distributed-kv-db/clientside/repl"
	"github.com/urfave/cli/v2"
)

func NewCli() []*cli.Command {
	return clientCli.New(repl.New())
}
