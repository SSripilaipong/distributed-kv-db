package cli

import (
	"distributed-kv-db/client/repl"
	"github.com/urfave/cli/v2"
)

func New(runRepl repl.Func) []*cli.Command {
	return []*cli.Command{
		{
			Name: "connect",
			Action: func(ctx *cli.Context) error {
				serverIp := ctx.Args().First()
				return runRepl(serverIp)
			},
		},
	}
}
