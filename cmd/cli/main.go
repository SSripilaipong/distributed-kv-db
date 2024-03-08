package main

import (
	"distributed-kv-db/clientside"
	"distributed-kv-db/serverside"
	"os"
)

func main() {
	run := New(serverside.NewCli(), clientside.NewCli())

	if err := run(os.Args); err != nil {
		panic(err)
	}
}
