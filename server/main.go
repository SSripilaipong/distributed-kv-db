package server

import (
	serverCli "distributed-kv-db/server/cli"
	"distributed-kv-db/server/db"
	"github.com/urfave/cli/v2"
)

func NewCli() []*cli.Command {
	return serverCli.New(db.New())
}
