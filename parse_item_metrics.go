package animelayer

import (
	"fmt"

	"golang.org/x/net/html"
)

func (p *parserDetailedItems) parseItemMetrics(n *html.Node) (*ItemMetrics, error) {
	clearTexts := make([]string, 0, 10)
	texts := getAllChildTextData(n)
	for _, t := range texts {
		t = cleanStringFromHtmlSymbols(t)
		if len(t) > 0 {
			clearTexts = append(clearTexts, t)
		}
	}

	if len(clearTexts) != 6 {
		return nil, fmt.Errorf("expected 6 texts, but got %d", len(clearTexts))
	}

	return &ItemMetrics{
		Uploads:         clearTexts[0],
		Downloads:       clearTexts[1],
		FilesSize:       clearTexts[2],
		Author:          clearTexts[3],
		VisitorCounter:  clearTexts[4],
		ApprovedCounter: clearTexts[5],

		ReadFromHtmlKey: "info pd20",
	}, nil
}

func (p *parserDetailedItems) parseItemMetricsFromCategoryPage(n *html.Node) (*ItemMetrics, error) {
	clearTexts := make([]string, 0, 10)
	texts := getAllChildTextData(n)
	for _, t := range texts {
		t = cleanStringFromHtmlSymbols(t)
		if len(t) > 0 {
			clearTexts = append(clearTexts, t)
		}
	}

	if len(clearTexts) != 6 {
		return nil, fmt.Errorf("expected 6 texts, but got %d", len(clearTexts))
	}

	return &ItemMetrics{
		Uploads:         clearTexts[0],
		Downloads:       clearTexts[1],
		FilesSize:       clearTexts[2],
		Author:          clearTexts[3],
		VisitorCounter:  clearTexts[4],
		ApprovedCounter: clearTexts[5],

		ReadFromHtmlKey: "info pd20",
	}, nil
}
