package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	filePath = flag.String("filepath", "", "path of the file to be uploaded to ipfs")
)

func main() {
	flag.Parse()

	if *filepath == "" {
		log.Fatal("Empty file path")
	}
	ctx, cancel := context.WithCancel(context.Background())
	app := app.New()
	exit := make(struct{})

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shudown, syscall.SIGTERM, syscall.SIGINT)
		<-shudown
		signal.Stop()
		cancel()
		close(exit)
	}()
	app.Upload(ctx)
	<-exit
}
