//go:build examgenerator

package animelayer_test

import (
	"context"
	"encoding/json"
	"os"

	"github.com/dmji/go-animelayer-parser"
)

func GenerateInitialItemExams(testFileHtml, testFileExam string, ctx context.Context, params TestGetItemParams) error {

	client, err := animelayer.HttpClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		return err
	}

	pReal := animelayer.New(&ClientHtmlSaveToFile{File: testFileHtml, Client: animelayer.NewHttpClientWrapper(client)}, animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))

	item, err := pReal.GetItemByIdentifier(ctx, params.Identifier)
	if err != nil {
		return err
	}
	data, err := json.Marshal(&item)
	if err != nil {
		return err
	}

	return os.WriteFile(testFileExam, data, 0644)
}

func GenerateInitialPageExams(testFileHtml, testFileExam string, ctx context.Context, params TestParseFirstPageParams) error {

	client, err := animelayer.HttpClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		return err
	}

	p := animelayer.New(&ClientHtmlSaveToFile{File: testFileHtml, Client: animelayer.NewHttpClientWrapper(client)}, animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))

	items, err := p.GetItemsFromCategoryPages(ctx, params.Category, params.Page)
	if err != nil {
		return err
	}
	data, err := json.Marshal(&items)
	if err != nil {
		return err
	}

	return os.WriteFile(testFileExam, data, 0644)
}
