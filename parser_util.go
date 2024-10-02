package animelayer

import (
	"strings"

	"golang.org/x/net/html"
)

func cleanStringFromHtmlSymbols(t string) string {
	t = strings.ReplaceAll(t, "\n", "")
	t = strings.ReplaceAll(t, "\t", "")
	t = strings.ReplaceAll(t, "\u00a0", " ")
	t = strings.ReplaceAll(t, "|", "")
	t = strings.TrimSpace(t)
	return t
}

func getFirstChildHrefNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isElementNodeData(c, "a") {
			return c
		}
	}
	return nil
}

func getFirstChildImgNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isElementNodeData(c, "img") {
			return c
		}
	}
	return nil
}

func getFirstChildTextData(n *html.Node) (string, bool) {
	if n.Type == html.TextNode {
		return n.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		text, ok := getFirstChildTextData(c)
		if ok {
			return text, true
		}
	}
	return "", false
}

func getAllChildTextData(n *html.Node) []string {
	res := make([]string, 0, 10)

	if n.Type == html.TextNode {
		res = append(res, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, getAllChildTextData(c)...)
	}

	return res
}

func isExistAttrWithTargetKeyValue(n *html.Node, key, value string) bool {
	for _, a := range n.Attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}

	return false
}

func isElementNodeData(n *html.Node, data string) bool {
	return n.Type == html.ElementNode && n.Data == data
}
