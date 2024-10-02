package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dmji/go-animelayer-parser"
)

type loggerBasic struct{}

func (l *loggerBasic) kts(keys ...interface{}) string {
	args := ""
	for _, key := range keys {

		switch key.(type) {

		case string:
			args += " " + key.(string)
		case int:
			args += " " + strconv.FormatInt(int64(key.(int)), 10)
		case []interface{}:
			args += l.kts(key.([]interface{})...)
		}

	}
	return args
}

func (l *loggerBasic) Infow(msg string, keys ...interface{}) {

	log.Print("Info  | ", msg, l.kts(keys))
}
func (l *loggerBasic) Errorw(msg string, keys ...interface{}) {
	log.Print("Error | ", msg, l.kts(keys))
}

func main() {
	var login, password string
	flag.StringVar(&login, "l", "", "login for credentials")
	flag.StringVar(&password, "p", "", "password for credentials")
	flag.Parse()

	log.Print(login, nil, password)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	cred := animelayer.Credentials{
		Login:    login,
		Password: password,
	}

	client, err := animelayer.HttpClientWithAuth(cred)
	if err != nil {
		panic(err)
	}

	p := animelayer.New(client)
	p.SetLogger(&loggerBasic{})
	pages := []int{1, 2, 3}

	// get first 3 anime pages
	pageNodes := p.PipeTargetPagessFromCategoryToPageNode(ctx, animelayer.Categories.Anime(), pages)

	// parse partial items
	partialItems := p.PipePageNodesToPartialItems(ctx, pageNodes)

	// got from partial items url to detailed html nodes
	itemNodes := p.PipePartialItemToItemNode(ctx, partialItems)

	// get detailed items form item nodes
	detailedItems := p.PipeItemNodesToDetailedItems(ctx, itemNodes)

	for {

		select {
		case <-ctx.Done():
			return
		case item, bOpen := <-detailedItems:

			if !bOpen && len(detailedItems) == 0 {
				return
			}

			log.Printf("Title: %s (%s)", item.Title, item.Identifier)
		}

	}

}
