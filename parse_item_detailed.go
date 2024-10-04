package animelayer

import (
	"context"
	"errors"

	"golang.org/x/net/html"
)

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

func traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *ItemDetailed) error {

	// cart status
	if isExistAttrWithTargetKeyValue(n, "div", "class", "info pd20") {

		clearTexts := make([]string, 0, 10)
		texts := getAllChildTextData(n)
		for _, t := range texts {

			t = cleanStringFromHtmlSymbols(t)
			if len(t) > 0 {
				clearTexts = append(clearTexts, t)
			}
		}

		if len(clearTexts) != 6 {
			return nil
		}

		// clearTexts[0]: Uploads
		// clearTexts[1]: Downloads
		// clearTexts[2]: Files size
		// clearTexts[3]: Author
		// clearTexts[4]: Visitor counter
		// clearTexts[5]: Approved counter
		item.TorrentFilesSize = clearTexts[2]

		return nil
	}

	// cart status date
	if isExistAttrWithTargetKeyValue(n, "div", "class", "info pd20 b0") {

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
			return errors.New("unexpected info in pd20 b0")
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

		return nil
	}

	// cart title
	if isExistAttrWithTargetKeyValue(n, "div", "class", "torrent-item torrent-info") {

		parseHeaderTitle(n, item)
		return nil
	}

	// cart description
	if isExistAttrWithTargetKeyValue(n, "div", "class", "description pd20 panel widget") {

		for c := n.FirstChild; c != nil; c = c.NextSibling {

			parseItemNotes(c, item)

		}

		return nil
	}

	// cart cover image
	if isExistAttrWithTargetKeyValue(n, "div", "class", "cover") {

		ref := getFirstChildImgNode(n)
		for _, a := range ref.Attr {
			if a.Key == "src" {
				item.RefImageCover = a.Val
				return nil
			}
		}

	}

	// cart additional image
	if isExistAttrWithTargetKeyValue(n, "div", "class", "panel widget pd20") {

		ref := getFirstChildHrefNode(n)
		for _, a := range ref.Attr {
			if a.Key == "href" {
				item.RefImagePreview = a.Val
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

func parseItem(ctx context.Context, doc *html.Node) *ItemDetailed {

	item := &ItemDetailed{}
	traverseHtmlItemNodes(ctx, doc, item)
	return item
}
