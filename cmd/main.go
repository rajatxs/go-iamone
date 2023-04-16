package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rajatxs/go-iamone/db"
	"github.com/rajatxs/go-iamone/server"
	"github.com/rajatxs/go-iamone/util"
)

func run() {
	util.Attempt(db.Open())
	util.Attempt(server.Start())
}

func terminate() {
	util.Attempt(db.Close())
	util.Attempt(server.Stop())
}

func main() {
	stop := make(chan os.Signal, 1)

	go run()

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	defer terminate()
}
