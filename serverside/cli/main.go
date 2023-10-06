package cli

import (
	"distributed-kv-db/serverside/db"
	"github.com/urfave/cli/v2"
)

func New(runDb db.Func) []*cli.Command {
	return []*cli.Command{
		{
			Name: "start",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  dbPortFlagName + ",p",
					Value: 5555,
					Usage: "port for db client to connect",
				},
				&cli.IntFlag{
					Name:  peeringPortFlagName + ",q",
					Value: 5556,
					Usage: "port for other peers to join network",
				},
				&cli.StringFlag{
					Name:  advertisedIpFlagName + ",h",
					Value: "localhost",
					Usage: "ip that other peers can reach to",
				},
				&cli.StringSliceFlag{
					Name:  peerAddressesFlagName,
					Usage: "addresses of other peers that can be connected to",
				},
			},
			Action: func(cCtx *cli.Context) error {
				var (
					dbPort        = cCtx.Int(dbPortFlagName)
					peeringPort   = cCtx.Int(peeringPortFlagName)
					advertisedIp  = cCtx.String(advertisedIpFlagName)
					peerAddresses = cCtx.StringSlice(peerAddressesFlagName)
				)
				return runDb(dbPort, peeringPort, advertisedIp, peerAddresses)
			},
		},
	}
}
