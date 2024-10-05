package animelayer

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func dateFromAnimelayerDate(t string) *time.Time {

	t = strings.ReplaceAll(t, "января", "01")
	t = strings.ReplaceAll(t, "февраля", "02")
	t = strings.ReplaceAll(t, "марта", "03")
	t = strings.ReplaceAll(t, "апреля", "04")
	t = strings.ReplaceAll(t, "мая", "05")
	t = strings.ReplaceAll(t, "июня", "06")
	t = strings.ReplaceAll(t, "июля", "07")
	t = strings.ReplaceAll(t, "августа", "08")
	t = strings.ReplaceAll(t, "сентября", "09")
	t = strings.ReplaceAll(t, "октября", "10")
	t = strings.ReplaceAll(t, "ноября", "11")
	t = strings.ReplaceAll(t, "декабря", "12")
	t = strings.ReplaceAll(t, " в ", " ")
	t = strings.ReplaceAll(t, ":", " ")

	numbers := strings.Split(t, " ")
	if len(numbers) == 4 {

		day, _ := strconv.ParseInt(numbers[0], 10, 64)
		month, _ := strconv.ParseInt(numbers[1], 10, 64)
		hour, _ := strconv.ParseInt(numbers[2], 10, 64)
		minute, _ := strconv.ParseInt(numbers[3], 10, 64)

		d := time.Date(time.Now().Year(), time.Month(month), int(day), int(hour), int(minute), 0, 0, time.UTC)
		return &d

	} else if len(numbers) == 5 {

		day, _ := strconv.ParseInt(numbers[0], 10, 64)
		month, _ := strconv.ParseInt(numbers[1], 10, 64)
		year, _ := strconv.ParseInt(numbers[2], 10, 64)
		hour, _ := strconv.ParseInt(numbers[3], 10, 64)
		minute, _ := strconv.ParseInt(numbers[4], 10, 64)

		d := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), 0, 0, time.UTC)
		return &d

	}

	return nil
}

type nodeWithParent struct {
	item *html.Node
	next *html.Node
}

func (p *parserDetailedItems) collectTextWithoudTags(root *html.Node, childsToReplace chan<- nodeWithParent) {

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if c.FirstChild == nil && c.Type == html.TextNode && c.Parent.Data == "div" {
			childsToReplace <- nodeWithParent{
				item: c,
				next: c.NextSibling,
			}
		}

		//p.collectTextWithoudTags(c, childsToReplace)
	}

}

func (p *parserDetailedItems) parseItemNotes(n *html.Node, item *ItemDetailed) {

	if len(p.NotePlaintTextElementInterceptor) > 0 {
		childsToReplaceChan := make(chan nodeWithParent, 10)
		go func() {
			defer close(childsToReplaceChan)
			p.collectTextWithoudTags(n, childsToReplaceChan)
		}()
		childsToReplace := make([]nodeWithParent, 0, 20)
		for i := range childsToReplaceChan {
			childsToReplace = append(childsToReplace, i)
		}

		for i := 0; i < len(childsToReplace); i++ {

			nodeToReplace := childsToReplace[i]

			data := &html.Node{
				Type: html.TextNode,
				Data: nodeToReplace.item.Data,
			}
			div := &html.Node{
				Type: html.ElementNode,
				Data: p.NotePlaintTextElementInterceptor,
			}
			div.AppendChild(data)

			if len(p.NotePlaintTextElementClassInterceptor) > 0 {
				div.Attr = append(div.Attr, html.Attribute{Key: "class", Val: p.NotePlaintTextElementClassInterceptor})
			}

			n.RemoveChild(nodeToReplace.item)
			n.InsertBefore(div, nodeToReplace.next)

			for {
				sib := div.NextSibling
				if sib == nil {
					break
				}
				sib2 := sib.NextSibling

				if isElementNodeData(sib, "br") && sib2.Type == html.TextNode {
					n.RemoveChild(sib)
					n.RemoveChild(sib2)
					//div.AppendChild(sib)
					//div.AppendChild(sib2)
					data.Data = data.Data + "\n" + sib2.Data
					childsToReplace = slices.DeleteFunc(childsToReplace, func(e nodeWithParent) bool { return e.item == sib || e.item == sib2 })
				} else {
					break
				}
			}
		}
	}

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

func (p *parserDetailedItems) tryReadNodeAsDivClass(n *html.Node, item *ItemDetailed, val string) (bool, error) {

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
			item.UpdatedDate = dateFromAnimelayerDate(clearTexts[1])
			item.CreatedDate = dateFromAnimelayerDate(clearTexts[3])
		} else if nText == 4 {
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: created date
			item.UpdatedDate = dateFromAnimelayerDate(clearTexts[1])
			item.CreatedDate = dateFromAnimelayerDate(clearTexts[3])
		} else if nText == 2 {
			// clearTexts[0]: Created
			// clearTexts[1]: created date
			item.CreatedDate = dateFromAnimelayerDate(clearTexts[1])
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

		p.parseItemNotes(n, item)
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

func (p *parserDetailedItems) traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *ItemDetailed) error {

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

func (p *parserDetailedItems) parseItem(ctx context.Context, doc *html.Node, identifier string) *ItemDetailed {

	item := &ItemDetailed{Identifier: identifier}
	p.traverseHtmlItemNodes(ctx, doc, item)
	return item
}

type parserDetailedItems struct {
	NotePlaintTextElementInterceptor      string
	NotePlaintTextElementClassInterceptor string
}
