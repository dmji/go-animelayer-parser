package animelayer

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (p *parser) parseItemMetrics(n *html.Node) (*ItemMetrics, error) {
	clearTexts := make([]string, 0, 10)
	texts := getAllChildTextData(n)
	for _, t := range texts {
		t = cleanStringFromSpecialSymbols(t)
		if len(t) > 0 {
			clearTexts = append(clearTexts, t)
		}
	}

	switch len(clearTexts) {
	case 6:

		return &ItemMetrics{
			Uploads:         clearTexts[0],
			Downloads:       clearTexts[1],
			FilesSize:       clearTexts[2],
			Author:          clearTexts[3],
			VisitorCounter:  clearTexts[4],
			ApprovedCounter: clearTexts[5],

			DebugReadFromElementClass: "info pd20",
		}, nil

	case 5:

		return &ItemMetrics{
			Uploads:         clearTexts[0],
			Downloads:       clearTexts[1],
			FilesSize:       clearTexts[2],
			VisitorCounter:  clearTexts[3],
			ApprovedCounter: clearTexts[4],

			DebugReadFromElementClass: "info pd20",
		}, nil

	default:
		return nil, fmt.Errorf("got unexpected texts count='%d'", len(clearTexts))
	}

}

func (p *parser) parseItemMetricsFromCategoryPage(n *html.Node, item *Item) error {
	clearTexts := make([]string, 0, 10)
	texts := getAllChildTextData(n)
	for _, t := range texts {
		t = cleanStringFromSpecialSymbols(t)
		if len(t) > 0 {
			clearTexts = append(clearTexts, t)
		}
	}

	switch len(clearTexts) {

	case 6:

		item.Metrics = ItemMetrics{
			Uploads:   clearTexts[0],
			Downloads: clearTexts[1],
			FilesSize: clearTexts[2],
			Author:    clearTexts[3],
			//VisitorCounter:  clearTexts[4],
			//ApprovedCounter: clearTexts[5],

			DebugReadFromElementClass: "info pd20",
		}

		updatedType := clearTexts[4]
		updatedData := p.dateFromAnimelayerDate(clearTexts[5])

		if strings.HasPrefix(updatedType, "Добавлен") {
			item.Updated.CreatedDate = updatedData
		} else if strings.HasPrefix(updatedType, "Обновлён") {
			item.Updated.UpdatedDate = updatedData
		} else {
			item.Updated.SeedLastPresenceDate = updatedData
		}

		item.Updated.DebugReadFromElementClass = "info pd20"
	case 5:

		item.Metrics = ItemMetrics{
			Uploads:   clearTexts[0],
			Downloads: clearTexts[1],
			FilesSize: clearTexts[2],
			//VisitorCounter:  clearTexts[3],
			//ApprovedCounter: clearTexts[4],

			DebugReadFromElementClass: "info pd20",
		}

		updatedType := clearTexts[3]
		updatedData := p.dateFromAnimelayerDate(clearTexts[4])

		if strings.HasPrefix(updatedType, "Добавлен") {
			item.Updated.CreatedDate = updatedData
		} else if strings.HasPrefix(updatedType, "Обновлён") {
			item.Updated.UpdatedDate = updatedData
		} else {
			item.Updated.SeedLastPresenceDate = updatedData
		}

		item.Updated.DebugReadFromElementClass = "info pd20"

	default:
		return fmt.Errorf("got unexpected texts count='%d'", len(clearTexts))
	}

	return nil
}
