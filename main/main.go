package main

import (
	"github.com/gursheyss/relay"
)

func main() {
	server, err := relay.Init()
	if err != nil {
		panic(err)
	}
	server()
}
