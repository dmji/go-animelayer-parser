package animelayer

import (
	"bytes"
	"slices"

	"golang.org/x/net/html"
)

type nodeWithParent struct {
	item *html.Node
	next *html.Node
}

func (p *parser) collectTextWithoudTags(root *html.Node, childsToReplace chan<- nodeWithParent) {

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if c.FirstChild == nil && c.Type == html.TextNode && c.Parent.Data == "div" {
			childsToReplace <- nodeWithParent{
				item: c,
				next: c.NextSibling,
			}
		}
	}

}

func (p *parser) parseItemNotes(n *html.Node) (string, error) {

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

				if isElementNodeData(sib, "br") && sib2.Type == html.TextNode && len(cleanStringFromSpecialSymbols(sib2.Data)) > 0 {
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
			return "", err
		}

	}

	return cleanStringFromSpecialSymbols(b.String()), nil
}
