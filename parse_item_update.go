package animelayer

import (
	"errors"

	"golang.org/x/net/html"
)

func (p *parserHtml) parseItemUpdate(n *html.Node) (*ItemUpdate, error) {
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
		return &ItemUpdate{
			UpdatedDate:          p.dateFromAnimelayerDate(clearTexts[1]),
			CreatedDate:          p.dateFromAnimelayerDate(clearTexts[3]),
			SeedLastPresenceDate: p.dateFromAnimelayerDate(clearTexts[5]),
		}, nil
	case 4:
		// clearTexts[0]: Updated
		// clearTexts[1]: Updated date
		// clearTexts[2]: Created
		// clearTexts[3]: created date
		return &ItemUpdate{
			UpdatedDate: p.dateFromAnimelayerDate(clearTexts[1]),
			CreatedDate: p.dateFromAnimelayerDate(clearTexts[3]),
		}, nil
	case 2:
		// clearTexts[0]: Created
		// clearTexts[1]: created date
		return &ItemUpdate{
			CreatedDate: p.dateFromAnimelayerDate(clearTexts[1]),
		}, nil
	default:
		return nil, errors.New("unexpected info in pd20 b0")
	}
}
