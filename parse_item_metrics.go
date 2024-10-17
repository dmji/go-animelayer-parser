package animelayer

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (p *parserHtml) parseItemMetrics(n *html.Node) (*ItemMetrics, error) {
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

func (p *parserHtml) parseItemMetricsFromCategoryPage(n *html.Node, item *Item) error {
	clearTexts := make([]string, 0, 10)
	texts := getAllChildTextData(n)
	for _, t := range texts {
		t = cleanStringFromHtmlSymbols(t)
		if len(t) > 0 {
			clearTexts = append(clearTexts, t)
		}
	}

	if len(clearTexts) != 6 {
		return fmt.Errorf("expected 6 texts, but got %d", len(clearTexts))
	}

	item.Metrics = ItemMetrics{
		Uploads:   clearTexts[0],
		Downloads: clearTexts[1],
		FilesSize: clearTexts[2],
		Author:    clearTexts[3],
		//VisitorCounter:  clearTexts[4],
		//ApprovedCounter: clearTexts[5],

		ReadFromHtmlKey: "info pd20",
	}

	updatedType := clearTexts[4]
	if strings.HasPrefix(updatedType, "Добавлен") {
		item.Updated.CreatedDate = p.dateFromAnimelayerDate(clearTexts[5])
	} else if strings.HasPrefix(updatedType, "Обновлён") {
		item.Updated.UpdatedDate = p.dateFromAnimelayerDate(clearTexts[5])
	} else {
		item.Updated.SeedLastPresenceDate = p.dateFromAnimelayerDate(clearTexts[5])
	}

	item.Updated.ReadFromHtmlKey = "info pd20"

	return nil
}
