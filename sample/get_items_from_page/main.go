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

	items, err := p.GetItemsFromCategoryPages(ctx, animelayer.CategoryAnime, 1)
	if err != nil {
		panic(err)
	}

	os.Mkdir("~dump", 0777)
	outfile := path.Join("~dump", time.Now().Format("2006-01-02T15:04:05")+".json")
	file, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}

	for i := range len(items) {
		items[i].Notes = ""
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&items)
	if err != nil {
		panic(err)
	}
}
