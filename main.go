package main

import (
	"distributed-kv-db/clientside"
	"distributed-kv-db/cmd"
	"distributed-kv-db/serverside"
	"os"
)

func main() {
	run := cmd.New(serverside.NewCli(), clientside.NewCli())

	if err := run(os.Args); err != nil {
		panic(err)
	}
}
