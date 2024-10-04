package animelayer

import (
	"bytes"
	"context"
	"errors"
	"slices"

	"golang.org/x/net/html"
)

func parseItemNotes(n *html.Node, item *ItemDetailed) {

	var b bytes.Buffer

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		err := html.Render(&b, c)
		if err != nil {
			panic(err)
		}

	}

	item.Notes = cleanStringFromHtmlSymbols(b.String())
}

func parseHeaderTitle(n *html.Node, item *ItemDetailed) bool {

	if isElementNodeData(n, "h1") {

		title, bCompleted, err := parseTitle(n)
		if err == nil {
			item.Title = title
			item.IsCompleted = bCompleted
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if bOk := parseHeaderTitle(c, item); bOk {
			return bOk
		}

	}

	return false
}

func tryReadNodeAsDivClass(n *html.Node, item *ItemDetailed, val string) (bool, error) {

	if !slices.Contains([]string{"info pd20",
		"info pd20 b0",
		"pd20",
		"description pd20 panel widget",
		"cover",
		"panel widget pd20",
	}, val) {
		return false, nil
	}

	// cart status
	if val == "info pd20" {

		clearTexts := make([]string, 0, 10)
		texts := getAllChildTextData(n)
		for _, t := range texts {

			t = cleanStringFromHtmlSymbols(t)
			if len(t) > 0 {
				clearTexts = append(clearTexts, t)
			}
		}

		if len(clearTexts) != 6 {
			return true, nil
		}

		// clearTexts[0]: Uploads
		// clearTexts[1]: Downloads
		// clearTexts[2]: Files size
		// clearTexts[3]: Author
		// clearTexts[4]: Visitor counter
		// clearTexts[5]: Approved counter
		item.TorrentFilesSize = clearTexts[2]

		return true, nil
	}

	// cart status date
	if val == "info pd20 b0" {

		clearTexts := make([]string, 0, 10)
		texts := getAllChildTextData(n)
		for _, t := range texts {

			t = cleanStringFromHtmlSymbols(t)
			if len(t) > 0 {
				clearTexts = append(clearTexts, t)
			}
		}

		nText := len(clearTexts)
		if nText != 2 && nText != 4 && nText != 6 {
			return false, errors.New("unexpected info in pd20 b0")
		}

		if nText == 6 {
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: Created date
			// clearTexts[4]: Seeder last presence
			// clearTexts[5]: Seed last presence date
			item.UpdatedDate = clearTexts[1]
			item.CreatedDate = clearTexts[3]
		} else if nText == 4 {
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: created date
			item.UpdatedDate = clearTexts[1]
			item.CreatedDate = clearTexts[3]
		} else if nText == 2 {
			// clearTexts[0]: Created
			// clearTexts[1]: created date
			item.CreatedDate = clearTexts[1]
		}

		return true, nil
	}

	// cart title
	if val == "pd20" {

		parseHeaderTitle(n, item)
		return true, nil
	}

	// cart description
	if val == "description pd20 panel widget" {

		parseItemNotes(n, item)
		return true, nil
	}

	// cart cover image
	if val == "cover" {

		ref := getFirstChildImgNode(n)
		val, bFound := getAttrByKey(ref, "src")
		if bFound {
			item.RefImageCover = val
			return true, nil
		}
	}

	// cart additional image
	if val == "panel widget pd20" {

		ref := getFirstChildHrefNode(n)
		val, bFound := getAttrByKey(ref, "href")
		if bFound {
			item.RefImagePreview = val
			return true, nil
		}
	}
	return false, nil
}

func traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *ItemDetailed) error {

	if isElementNodeData(n, "div") {
		divClassValue, bFound := getAttrByKey(n, "class")

		if bFound {
			bFinish, err := tryReadNodeAsDivClass(n, item, divClassValue)
			if err != nil {
				return err
			}
			if bFinish {
				return nil
			}
		}
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			traverseHtmlItemNodes(ctx, c, item)
		}
	}

	return nil
}

func parseItem(ctx context.Context, doc *html.Node, identifier string) *ItemDetailed {

	item := &ItemDetailed{Identifier: identifier}
	traverseHtmlItemNodes(ctx, doc, item)
	return item
}
