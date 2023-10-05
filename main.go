package main

import (
	"distributed-kv-db/appcli"
	"distributed-kv-db/client"
	"distributed-kv-db/server"
	"os"
)

func main() {
	run := appcli.New(server.NewCli(), client.NewCli())

	if err := run(os.Args); err != nil {
		panic(err)
	}
}
