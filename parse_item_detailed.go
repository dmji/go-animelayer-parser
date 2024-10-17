package animelayer

/*
import (
	"context"
	"errors"

	"golang.org/x/net/html"
)

func (p *parserDetailedItems) tryReadNodeAsDivClass(n *html.Node, item *Item, val string) (bool, error) {

	switch val {

	case "info pd20": // cart status
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

		item.Metrics = ItemMetrics{
			Uploads:         clearTexts[0],
			Downloads:       clearTexts[1],
			FilesSize:       clearTexts[2],
			Author:          clearTexts[3],
			VisitorCounter:  clearTexts[4],
			ApprovedCounter: clearTexts[5],

			ReadFromHtmlKey: "info pd20",
		}

		return true, nil
	case "info pd20 b0": // cart status date
		clearTexts := make([]string, 0, 10)
		texts := getAllChildTextData(n)
		for _, t := range texts {

			t = cleanStringFromHtmlSymbols(t)
			if len(t) > 0 {
				clearTexts = append(clearTexts, t)
			}
		}

		nText := len(clearTexts)
		switch nText {
		case 6:
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: Created date
			// clearTexts[4]: Seeder last presence
			// clearTexts[5]: Seed last presence date
			item.Updated = ItemUpdate{
				UpdatedDate:          p.dateFromAnimelayerDate(clearTexts[1]),
				CreatedDate:          p.dateFromAnimelayerDate(clearTexts[3]),
				SeedLastPresenceDate: p.dateFromAnimelayerDate(clearTexts[5]),
			}
		case 4:
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: created date
			item.Updated = ItemUpdate{
				UpdatedDate: p.dateFromAnimelayerDate(clearTexts[1]),
				CreatedDate: p.dateFromAnimelayerDate(clearTexts[3]),
			}
		case 2:
			// clearTexts[0]: Created
			// clearTexts[1]: created date
			item.Updated = ItemUpdate{
				CreatedDate: p.dateFromAnimelayerDate(clearTexts[1]),
			}
		default:
			return false, errors.New("unexpected info in pd20 b0")
		}

		item.Updated.ReadFromHtmlKey = "info pd20 b0"
		return true, nil
	case "description pd20 panel widget": // cart description
		note, _ := p.parseItemNotes(n)
		item.Notes = note
		return true, nil
	case "cover": // cart cover image
		ref := getFirstChildImgNode(n)
		val, bFound := getAttrByKey(ref, "src")
		if bFound {
			item.RefImageCover = val
			return true, nil
		}
	case "panel widget pd20": // cart additional image
		ref := getFirstChildHrefNode(n)
		val, bFound := getAttrByKey(ref, "href")
		if bFound {
			item.RefImagePreview = val
			return true, nil
		}
	}

	return false, nil
}

func (p *parserDetailedItems) traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *Item) error {

	if isElementNodeData(n, "div") {
		divClassValue, bFound := getAttrByKey(n, "class")

		if bFound {
			bFinish, err := p.tryReadNodeAsDivClass(n, item, divClassValue)
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
			p.traverseHtmlItemNodes(ctx, c, item)
		}
	}

	return nil
}

func (p *parserDetailedItems) parseItem(ctx context.Context, doc *html.Node, identifier string) (*Item, error) {

	item := &Item{Identifier: identifier}
	err := p.traverseHtmlItemNodes(ctx, doc, item)
	return item, err
}

type parserDetailedItems struct {
	NotePlaintTextElementInterceptor      string
	NotePlaintTextElementClassInterceptor string
}
*/
