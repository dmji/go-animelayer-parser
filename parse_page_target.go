package animelayer

import (
	"context"

	"golang.org/x/net/html"
)

func (p *parserHtml) tryReadNodeAsDivClass(n *html.Node, item *Item, val string) (bool, error) {

	switch val {

	case "info pd20": // cart status
		metrics, err := p.parseItemMetrics(n)
		if err != nil {
			return false, err
		}

		item.Metrics = *metrics
		return true, nil
	case "info pd20 b0": // cart status date
		update, err := p.parseItemUpdate(n)
		if err != nil {
			return false, err
		}

		item.Updated = *update
		item.Updated.ReadFromHtmlKey = "info pd20 b0"
		return true, nil
	case "description pd20 panel widget": // cart description
		note, err := p.parseItemNotes(n)
		if err != nil {
			return false, err
		}

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

func (p *parserHtml) traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *Item) error {

	if isExistAttrWithTargetKeyValue(n, "meta", "property", "og:title") {
		val, bFound := getAttrByKey(n, "content")
		if bFound {
			title, bCompleted := p.grabTitleWithCompletedStatus(val)
			item.Title = title
			item.IsCompleted = bCompleted
			return nil
		}
	}

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

func (p *parserHtml) ParseItem(ctx context.Context, doc *html.Node, identifier string) (*Item, error) {

	item := &Item{Identifier: identifier}
	err := p.traverseHtmlItemNodes(ctx, doc, item)
	return item, err
}
