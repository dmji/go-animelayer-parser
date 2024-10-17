package animelayer

import (
	"context"
	"errors"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func (p *parserDetailedItems) parseItemTitle(n *html.Node, item *Item) error {

	ref := getFirstChildHrefNode(n)
	if ref == nil {
		return errors.New("href not found")
	}

	identifier, bFound := getAttrByKey(ref, "href")
	if !bFound {
		return errors.New("href attr not found")
	}

	identifier, bFound = strings.CutPrefix(identifier, "/torrent/")
	if !bFound {
		return errors.New("got unexpected url prefix")
	}
	identifier, bFound = strings.CutSuffix(identifier, "/")
	if !bFound {
		return errors.New("got unexpected url suffix")
	}
	item.Identifier = identifier

	name, bOk := getFirstChildTextData(ref)
	if !bOk {
		return errors.New("failed to get title")
	}

	title, bCompleted := p.grabTitleWithCompletedStatus(name)
	item.Title = title
	item.IsCompleted = bCompleted
	return nil
}

func (p *parserDetailedItems) tryReadCardNodeAsDivClass(n *html.Node, item *Item, val string) (bool, error) {

	switch val {

	case "info pd20": // cart status
		err := p.parseItemMetricsFromCategoryPage(n, item)
		if err != nil {
			return false, err
		}

		return true, nil
	case "description": // cart description
		note, err := p.parseItemNotes(n)
		if err != nil {
			return false, err
		}

		item.Notes = note
		return true, nil
	case "pd20": // cart cover image
		ref := getFirstChildImgNode(n)
		val, bFound := getAttrByKey(ref, "data-original")
		if bFound {
			item.RefImageCover = val
			return false, nil
		}
	}

	return false, nil
}

func (p *parserDetailedItems) traverseHtmlCardNodes(ctx context.Context, n *html.Node, item *Item) error {

	// cart title
	if isExistAttrWithTargetKeyValue(n, "h3", "class", "h2 m0") {

		err := p.parseItemTitle(n, item)
		if err != nil {
			return err
		}
		return nil
	}

	if isElementNodeData(n, "div") {
		divClassValue, bFound := getAttrByKey(n, "class")

		if bFound {
			bFinish, err := p.tryReadCardNodeAsDivClass(n, item, divClassValue)
			if err != nil {
				return err
			}
			if bFinish {
				return nil
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := p.traverseHtmlCardNodes(ctx, c, item)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *parserDetailedItems) parseCategoryPage(ctx context.Context, n *html.Node, chItems chan<- ItemPartialWithError, wg *sync.WaitGroup) {

	if isExistAttrWithTargetKeyValue(n, "li", "class", "torrent-item torrent-item-medium panel") {

		wg.Add(1)
		go func() {
			defer wg.Done()

			item := &Item{}
			err := p.traverseHtmlCardNodes(ctx, n, item)

			if err != nil {
				chItems <- ItemPartialWithError{
					Item:  nil,
					Error: err,
				}
			}

			chItems <- ItemPartialWithError{
				Item:  item,
				Error: nil,
			}
		}()

		return
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return
		default:
			p.parseCategoryPage(ctx, c, chItems, wg)
		}
	}
}

func (p *parserDetailedItems) ParseCategoryPageToChan(ctx context.Context, category *CategoryHtml, chItems chan<- ItemPartialWithError) {

	wg := &sync.WaitGroup{}
	p.parseCategoryPage(ctx, category.Node, chItems, wg)
	wg.Wait()

}