package animelayer_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/dmji/go-animelayer-parser"
	"github.com/joho/godotenv"
)

func init() {

	path := ".env"
	for i := range 10 {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}
}

type TestParseFirstPageParams struct {
	Category      animelayer.Category
	Page          int
	ExpectedError error
}

func isSameError(got, expected error) bool {
	if got == nil && expected == nil {
		return true
	}

	if got != nil && expected != nil {
		return got.Error() == expected.Error()
	}

	return false
}

func isValidPartialItem(item *animelayer.ItemPartial) error {

	if len(item.Identifier) == 0 {
		return errors.New("got empty identifier")
	}

	if len(item.Title) == 0 {
		return errors.New("got empty title")
	}

	return nil
}

func TestGetFirstPageOfCategory(t *testing.T) {

	testParamss := []TestParseFirstPageParams{
		{
			Category: animelayer.Categories.Anime(),
			Page:     1,
		},
		{
			Category: animelayer.Categories.Manga(),
			Page:     2,
		},
		{
			Category:      animelayer.Categories.Anime(),
			Page:          1500,
			ExpectedError: errors.New("empty document"),
		},
	}

	ctx := context.Background()

	client, err := animelayer.HttpClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		t.Fatal(err)
	}

	p := animelayer.New(client)

	for _, params := range testParamss {

		items, err := p.CategoryPageToPartialItems(ctx, params.Category, params.Page)
		if gotErr := err; !isSameError(gotErr, params.ExpectedError) {
			t.Fatalf("got error='%v', expected error='%v'", gotErr, params.ExpectedError)
		}

		for _, item := range items {
			err := isValidPartialItem(&item)
			if err != nil {
				t.Fatal(err)
			}

		}
	}

}

type TestGetDetailedItemParams struct {
	Identifier string

	NoteElem  string
	NoteClass string
}

func nameForTestDataFile(p TestGetDetailedItemParams) string {
	if p.NoteClass == "" && p.NoteElem == "" {
		return fmt.Sprintf("%s", p.Identifier)
	}

	if p.NoteClass == "" && p.NoteElem != "" {
		return fmt.Sprintf("%s_%s", p.Identifier, p.NoteElem)
	}

	return fmt.Sprintf("%s_%s#%s", p.Identifier, p.NoteElem, p.NoteClass)
}

func isEqualDetailedItem(got, expected *animelayer.ItemDetailed) error {

	if got.Identifier != expected.Identifier {
		return fmt.Errorf("expected Identifier='%s', but got='%s'", expected.Identifier, got.Identifier)
	}

	if got.Title != expected.Title {
		return fmt.Errorf("expected Title='%s', but got='%s", expected.Title, got.Title)
	}

	if got.IsCompleted != expected.IsCompleted {
		return fmt.Errorf("expected IsCompleted='%v', but got='%v", expected.IsCompleted, got.IsCompleted)
	}

	if got.Metrics.ReadFromHtmlKey != expected.Metrics.ReadFromHtmlKey {
		return fmt.Errorf("expected Metrics.ReadFromHtmlKey='%s', but got='%s", expected.Metrics.ReadFromHtmlKey, got.Metrics.ReadFromHtmlKey)
	}

	if got.Updated.ReadFromHtmlKey != expected.Updated.ReadFromHtmlKey {
		return fmt.Errorf("expected Updated.ReadFromHtmlKey='%s', but got='%s", expected.Updated.ReadFromHtmlKey, got.Updated.ReadFromHtmlKey)
	}

	if got.RefImagePreview != expected.RefImagePreview {
		return fmt.Errorf("expected RefImagePreview='%s', but got='%s", expected.RefImagePreview, got.RefImagePreview)
	}

	if got.RefImageCover != expected.RefImageCover {
		return fmt.Errorf("expected RefImageCover='%s', but got='%s", expected.RefImageCover, got.RefImageCover)
	}

	if got.Notes != expected.Notes {
		return fmt.Errorf("expected Notes='%s', but got='%s", expected.Notes, got.Notes)
	}

	return nil
}

func TestGetDetailedItem(t *testing.T) {

	testParamss := []TestGetDetailedItemParams{
		{
			Identifier: "670ad89197553b37b913cea8",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
		},
		{
			Identifier: "670ad89197553b37b913cea8",
			NoteElem:   "i",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
			NoteElem:   "i",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
			NoteElem:   "i",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
			NoteElem:   "i",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
			NoteElem:   "i",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
			NoteElem:   "i",
		},
		{
			Identifier: "670ad89197553b37b913cea8",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
			NoteElem:   "i",
			NoteClass:  "test",
		},
	}

	ctx := context.Background()

	client, err := animelayer.HttpClientWithAuth(animelayer.Credentials{
		Login:    os.Getenv("ANIME_LAYER_LOGIN"),
		Password: os.Getenv("ANIME_LAYER_PASSWORD"),
	})

	if err != nil {
		t.Fatal(err)
	}

	for _, params := range testParamss {

		p := animelayer.New(client, animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))
		// generate initial result
		/*
			   {
					itemDetailed, err := p.PartialItemToDetailedItem(ctx, params.Identifier)
					if err != nil {
						panic(err)
					}
					data, err := json.Marshal(&itemDetailed)
					if err != nil {
						panic(err)
					}
					os.WriteFile(fmt.Sprintf("test_data/%s.json", nameForTestDataFile(params)), data, 0644)
				}
		*/

		// test
		item, err := p.PartialItemToDetailedItem(ctx, params.Identifier)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}

		data, err := os.ReadFile(fmt.Sprintf("test_data/%s.json", nameForTestDataFile(params)))
		if err != nil {
			t.Fatalf("error on read exam item: %v", err)
		}
		var expectedItem animelayer.ItemDetailed
		err = json.Unmarshal(data, &expectedItem)
		if err != nil {
			t.Fatalf("error on unmarshal exam item: %v", err)
		}

		err = isEqualDetailedItem(item, &expectedItem)
		if err != nil {
			t.Fatal(err)
		}
	}

}
