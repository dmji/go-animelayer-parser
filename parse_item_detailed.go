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

				if isElementNodeData(sib, "br") && sib2.Type == html.TextNode && len(cleanStringFromHtmlSymbols(sib2.Data)) > 0 {
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
				UpdatedDate:          dateFromAnimelayerDate(clearTexts[1]),
				CreatedDate:          dateFromAnimelayerDate(clearTexts[3]),
				SeedLastPresenceDate: dateFromAnimelayerDate(clearTexts[5]),
			}
		case 4:
			// clearTexts[0]: Updated
			// clearTexts[1]: Updated date
			// clearTexts[2]: Created
			// clearTexts[3]: created date
			item.Updated = ItemUpdate{
				UpdatedDate: dateFromAnimelayerDate(clearTexts[1]),
				CreatedDate: dateFromAnimelayerDate(clearTexts[3]),
			}
		case 2:
			// clearTexts[0]: Created
			// clearTexts[1]: created date
			item.Updated = ItemUpdate{
				UpdatedDate: dateFromAnimelayerDate(clearTexts[1]),
			}
		default:
			return false, errors.New("unexpected info in pd20 b0")
		}

		item.Updated.ReadFromHtmlKey = "info pd20 b0"
		return true, nil
	case "pd20": // cart title
		parseHeaderTitle(n, item)
		return true, nil
	case "description pd20 panel widget": // cart description
		p.parseItemNotes(n, item)
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

func (p *parserDetailedItems) parseItem(ctx context.Context, doc *html.Node, identifier string) (*ItemDetailed, error) {

	item := &ItemDetailed{Identifier: identifier}
	err := p.traverseHtmlItemNodes(ctx, doc, item)
	return item, err
}

type parserDetailedItems struct {
	NotePlaintTextElementInterceptor      string
	NotePlaintTextElementClassInterceptor string
}
