package animelayer

import (
	"context"
	"time"

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

func (p *parser) traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *ItemDetailed) error {

	// cart status
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "info pd20") {

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

	}

	// cart status date
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "info pd20 b0") {

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
				p.logger.Errorw("unexpected info in pd20 b0")
				return nil
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

	}

	// cart title
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "torrent-item torrent-info") {

			parseHeaderTitle(n, item)

		}

	}

	// cart description
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "description pd20 panel widget") {

			for c := n.FirstChild; c != nil; c = c.NextSibling {

				p.parseItemNotes(c, item)

			}

			return nil
		}

	}

	// cart cover image
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "cover") {

			ref := getFirstChildImgNode(n)
			for _, a := range ref.Attr {
				if a.Key == "src" {
					item.RefImageCover = a.Val
					return nil
				}
			}

		}

	}

	// cart additional image
	if isElementNodeData(n, "div") {

		if isExistAttrWithTargetKeyValue(n, "class", "panel widget pd20") {

			ref := getFirstChildHrefNode(n)
			for _, a := range ref.Attr {
				if a.Key == "href" {
					item.RefImagePreview = a.Val
					return nil
				}
			}

		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			p.traverseHtmlItemNodes(ctx, c, item)
		}
	}

	return nil
}

func (p *parser) ParseItem(ctx context.Context, doc *html.Node) *ItemDetailed {

	item := &ItemDetailed{}
	p.traverseHtmlItemNodes(ctx, doc, item)
	item.LastCheckedDate = time.Now().Format("2006-01-02 15:04")
	return item
}
