package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/dmji/go-animelayer-parser"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	p := animelayer.New(animelayer.NewClientWrapper(&http.Client{}), animelayer.WithNoteClassOverride("i", ""))

	itemId := "60ddae82fd89780039313044"
	item, err := p.GetItemByIdentifier(ctx, itemId)
	if err != nil {
		panic(err)
	}

	os.Mkdir("~dump", 0777)
	outfile := path.Join("~dump", time.Now().Format("2006-01-02T15:04:05")+".json")
	file, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&item)
	if err != nil {
		panic(err)
	}
}
