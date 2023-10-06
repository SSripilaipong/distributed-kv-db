package serverside

import (
	serverCli "distributed-kv-db/serverside/cli"
	"distributed-kv-db/serverside/db"
	"distributed-kv-db/serverside/db/machine"
	"github.com/urfave/cli/v2"
)

func NewCli() []*cli.Command {
	return serverCli.New(db.New(machine.WaitForInterrupt))
}
