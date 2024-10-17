package animelayer_test

import (
	"bytes"
	"os"

	"github.com/dmji/go-animelayer-parser"
	"golang.org/x/net/html"
)

type ClientHtmlGetFromFile struct {
	File string
}

func (f *ClientHtmlGetFromFile) Get(url string) (*html.Node, error) {

	data, err := os.ReadFile(f.File)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

type ClientHtmlSaveToFile struct {
	File   string
	Client animelayer.DocGetter
}

func (f *ClientHtmlSaveToFile) Get(url string) (*html.Node, error) {

	doc, err := f.Client.Get(url)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = html.Render(&b, doc)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(f.File, b.Bytes(), 0644)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
