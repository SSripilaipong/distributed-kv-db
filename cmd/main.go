package cmd

import (
	"github.com/urfave/cli/v2"
)

func New(serverCommands, clientCommands []*cli.Command) Func {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "server",
				Subcommands: serverCommands,
			},
			{
				Name:        "client",
				Subcommands: clientCommands,
			},
		},
	}
	return (app).Run
}
