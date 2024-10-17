package animelayer_test

import (
	"context"
	"os"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func TestGetItemsFromCategoryPages(t *testing.T) {

	testParamss := testCategoryPages()

	ctx := context.Background()

	client, err := animelayer.HttpClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		t.Fatal(err)
	}

	p := animelayer.New(animelayer.NewHttpClientWrapper(client))

	for _, params := range testParamss {

		items, err := p.GetItemsFromCategoryPages(ctx, params.Category, params.Page)
		if gotErr := err; !isSameError(gotErr, params.ExpectedError) {
			t.Fatalf("got error='%v', expected error='%v'", gotErr, params.ExpectedError)
		}

		for _, item := range items {
			err := isValidItem(&item, false)
			if err != nil {
				t.Fatal(err)
			}

		}
	}

}
