package animelayer

import (
	"bytes"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func removeHTMLTag(s, tag string) string {
	re := regexp.MustCompile(fmt.Sprintf(`<\/?%s([^>]+)?>`, tag))
	s = re.ReplaceAllString(s, "")
	return s
}

func removeHTMLTags(s string, tags ...string) string {
	if len(tags) == 0 {
		s = removeHTMLTag(s, "")
	}

	for _, tag := range tags {
		s = removeHTMLTag(s, tag)
	}

	return strings.TrimRight(s, ": ")
}

type notesSplitted struct {
	Tag  string
	Text string

	// if tag not found put text to others
	Others string
}

func splitTextByTags(text, tag string) []notesSplitted {
	res := make([]notesSplitted, 0)

	reBeg := regexp.MustCompile(fmt.Sprintf(`<%s([^>]+)?>`, tag))
	stringsFromBeg := reBeg.Split(text, -1)

	reEnd := regexp.MustCompile(fmt.Sprintf(`<\/%s([^>]+)?>`, tag))
	for _, s := range stringsFromBeg {
		vals := reEnd.Split(s, -1)
		if len(vals) != 2 {
			res = append(res, notesSplitted{Others: vals[0]})
		} else {
			res = append(res, notesSplitted{Tag: vals[0], Text: vals[1]})
		}
	}
	return res
}

func extractDivSpoiler(text string) (string, NotesSematizied) {
	result := NotesSematizied{}

	text, divs := extractAndRemoveDiv(text)

	for _, s := range divs {

		vals := strings.Split(s, "</div>")
		for i := range len(vals) {
			vals[i] = removeHTMLTags(vals[i], "div", "span")
		}

		vals = slices.DeleteFunc(vals, func(e string) bool { return len(e) == 0 })
		if len(vals) != 2 {
			if len(removeHTMLTags(s)) > 0 {
				result.Untaged = append(result.Untaged, s)
			}
			continue
		}

		result.Taged = append(result.Taged, NotesSematiziedItem{
			Tag:    vals[0],
			Childs: TryGetSomthingSemantizedFromNotes(vals[1]),
		})
	}

	return text, result
}

func extractAndRemoveDiv(input string) (string, []string) {
	rootNode, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	divs := removeDiv(rootNode)

	for _, d := range divs {
		i := strings.Index(input, d)
		input = strings.Join(strings.Split(input, d), "")
		_ = i
	}

	return input, divs
}

func removeDiv(n *html.Node) []string {
	res := make([]string, 0)

	if n.Type == html.ElementNode && n.Data == "div" {

		var buffer bytes.Buffer
		html.Render(&buffer, n)

		n.Parent.RemoveChild(n)
		return []string{strings.ReplaceAll(buffer.String(), "&#34;", "\"")}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, removeDiv(c)...)
	}

	return res
}

func strongLinesToNoteItems(strongLines []notesSplitted) *NotesSematizied {
	result := &NotesSematizied{}

	for _, line := range strongLines {
		if len(line.Others) > 0 {
			if len(removeHTMLTags(line.Others)) > 0 {
				result.Untaged = append(result.Untaged, line.Others)
			}
			continue
		}

		result.Taged = append(result.Taged, NotesSematiziedItem{
			Tag:  removeHTMLTags(line.Tag),
			Text: removeHTMLTags(line.Text),
		})
	}

	return result
}

func TryGetSomthingSemantizedFromNotes(text string) *NotesSematizied {
	result := &NotesSematizied{}

	text, tagItems := extractDivSpoiler(text)
	result.Extend(&tagItems)

	uLines := splitTextByTags(text, "u")
	for _, line := range uLines {
		var s string
		if len(line.Others) > 0 {
			s = line.Others
		} else {
			s = line.Text
		}

		strongLines := splitTextByTags(s, "strong")
		if len(line.Others) > 0 {
			result.Extend(strongLinesToNoteItems(strongLines))
		} else {
			result.Taged = append(result.Taged, NotesSematiziedItem{
				Tag:    removeHTMLTags(line.Tag),
				Childs: strongLinesToNoteItems(strongLines),
			})
		}
	}

	return result
}
