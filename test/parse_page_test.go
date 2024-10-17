package animelayer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func nameForTestPageDataFile(p TestParseFirstPageParams) string {
	if p.NoteClass == "" && p.NoteElem == "" {
		return strings.ReplaceAll(fmt.Sprintf("page_%s%d", p.Category, p.Page), "/", "_")
	}

	if p.NoteClass == "" && p.NoteElem != "" {
		return strings.ReplaceAll(fmt.Sprintf("page_%s%d#%s", p.Category, p.Page, p.NoteElem), "/", "_")
	}

	return strings.ReplaceAll(fmt.Sprintf("page_%s%d#%s_%s", p.Category, p.Page, p.NoteElem, p.NoteClass), "/", "_")
}

func TestParsePage(t *testing.T) {

	testParamss := testCategoryPages()
	ctx := context.Background()

	for _, params := range testParamss {

		testFileExam := fmt.Sprintf("test_data/%s.json", nameForTestPageDataFile(params))
		testFileHtml := fmt.Sprintf("test_data/%s.html", nameForTestPageDataFile(params))

		GenerateInitialPageExams(testFileHtml, testFileExam, ctx, params)

		// Parse from file
		p := animelayer.New(&ClientHtmlGetFromFile{File: testFileHtml}, animelayer.WithNoteClassOverride(params.NoteElem, params.NoteClass))

		// test
		items, err := p.GetItemsFromCategoryPages(ctx, params.Category, params.Page)
		if !isSameError(err, params.ExpectedError) {
			t.Fatalf("expected error='%v', but got error='%v'", params.ExpectedError, err)
		}
		data, err := os.ReadFile(testFileExam)
		if err != nil {
			t.Fatalf("error on read exam item: %v", err)
		}
		var expectedItems []animelayer.Item
		err = json.Unmarshal(data, &expectedItems)
		if err != nil {
			t.Fatalf("error on unmarshal exam item: %v", err)
		}

		if len(items) != len(expectedItems) {
			t.Fatal("length of items slices not equal")
		}

		for i := range len(items) {

			iExam := slices.IndexFunc(expectedItems, func(item animelayer.Item) bool { return item.Identifier == items[i].Identifier })
			if iExam == -1 {
				t.Fatalf("not found item identifier '%s' in exam items", items[i].Identifier)
			}

			err = isEqualItem(&items[i], &expectedItems[iExam])
			if err != nil {
				t.Fatal(testFileExam, err)
			}

		}
	}
}
