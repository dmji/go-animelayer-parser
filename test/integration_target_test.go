package animelayer_test

import (
	"context"
	"os"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func TestGetTargetItem(t *testing.T) {

	testParamss := testListIdentifiers()

	ctx := context.Background()

	client, err := animelayer.DefaultClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		t.Fatal(err)
	}

	for _, params := range testParamss {

		p := animelayer.New(animelayer.NewClientWrapper(client), animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))

		item, err := p.GetItemByIdentifier(ctx, params.Identifier)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}

		err = isValidItem(item, true)
		if err != nil {
			t.Fatal(err)
		}
	}

}
