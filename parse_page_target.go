package animelayer

import (
	"context"
	"errors"

	"golang.org/x/net/html"
)

type loopExitStatus bool

const (
	continueLoopExitStatus = false
	breakLoopExitStatus    = true
)

func (p *parser) tryReadNodeAsDivClass(n *html.Node, item *Item, val string) (loopExitStatus, error) {
	var err error
	switch val {

	case "info pd20": // cart status
		metrics, err := p.parseItemMetrics(n)
		if err != nil {
			return continueLoopExitStatus, err
		}

		item.Metrics = *metrics
		return breakLoopExitStatus, nil
	case "info pd20 b0": // cart status date
		update, err := p.parseItemUpdate(n)
		if err != nil {
			return continueLoopExitStatus, err
		}

		item.Updated = *update
		item.Updated.DebugReadFromElementClass = "info pd20 b0"
		return breakLoopExitStatus, nil
	case "description pd20 panel widget": // cart description
		note, err := p.parseItemNotes(n)
		if err != nil {
			return continueLoopExitStatus, err
		}

		item.Notes = note
		item.NotesSematizied = TryGetSomthingSemantizedFromNotes(note)
		return breakLoopExitStatus, nil
	case "cover": // cart cover image
		href := getFirstChildHrefNode(n)
		if href == nil {
			return continueLoopExitStatus, errors.New("not found href in cover div")
		}
		categoryPresentation, bFound := getFirstChildTextData(href)
		if bFound {
			item.Category, err = categoryFromPresentationString(categoryPresentation)
			if err != nil {
				return continueLoopExitStatus, err
			}
		}

		ref := getFirstChildImgNode(n)
		if ref == nil {
			return continueLoopExitStatus, errors.New("not found image in cover div")
		}
		val, bFound := getAttrByKey(ref, "src")
		if bFound {
			item.RefImageCover = val
			return breakLoopExitStatus, nil
		}
	case "panel widget pd20": // cart additional image
		ref := getFirstChildHrefNode(n)
		val, bFound := getAttrByKey(ref, "href")
		if bFound {
			item.RefImagePreview = val
			return breakLoopExitStatus, nil
		}
	}

	return continueLoopExitStatus, nil
}

func (p *parser) traverseItemNodes(ctx context.Context, n *html.Node, item *Item) error {
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
			p.traverseItemNodes(ctx, c, item)
		}
	}

	return nil
}

func (p *parser) ParseItem(ctx context.Context, doc *html.Node, identifier string) (*Item, error) {
	item := &Item{Identifier: identifier}
	err := p.traverseItemNodes(ctx, doc, item)
	return item, err
}
