package animelayer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func nameForTestItemDataFile(p TestGetItemParams) string {
	if p.NoteClass == "" && p.NoteElem == "" {
		return fmt.Sprintf("%s", p.Identifier)
	}

	if p.NoteClass == "" && p.NoteElem != "" {
		return fmt.Sprintf("%s_%s", p.Identifier, p.NoteElem)
	}

	return fmt.Sprintf("%s_%s#%s", p.Identifier, p.NoteElem, p.NoteClass)
}

func TestParseItem(t *testing.T) {

	testParamss := testListIdentifiers()
	ctx := context.Background()

	for _, params := range testParamss {

		testFileExam := fmt.Sprintf("test_data/%s.json", nameForTestItemDataFile(params))
		testFileHtml := fmt.Sprintf("test_data/%s", nameForTestItemDataFile(params))

		if err := GenerateInitialItemExams(testFileHtml, testFileExam, ctx, params); err != nil {
			panic(err)
		}

		// Parse from file
		p := animelayer.New(&ClientHtmlGetFromFile{File: testFileHtml}, animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))

		// test
		item, err := p.GetItemByIdentifier(ctx, params.Identifier)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		data, err := os.ReadFile(testFileExam)
		if err != nil {
			t.Fatalf("error on read exam item: %v", err)
		}
		var expectedItem animelayer.Item
		err = json.Unmarshal(data, &expectedItem)
		if err != nil {
			t.Fatalf("error on unmarshal exam item: %v", err)
		}

		err = isEqualItem(item, &expectedItem)
		if err != nil {
			t.Fatal(testFileExam, err)
		}
	}
}
