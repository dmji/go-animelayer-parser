package animelayer

import "strings"

type parser struct {
	NotePlaintTextElementInterceptor      string
	NotePlaintTextElementClassInterceptor string
}

var suffixes = []string{
	" Complete",
	"Complete",
	" Сomplete",
	"Сomplete",
}

func (p *parser) grabTitleWithCompletedStatus(name string) (string, bool) {
	title := cleanStringFromSpecialSymbols(name)
	bCompleted := false

	for _, suffix := range suffixes {
		if titleCuted, bFound := strings.CutSuffix(title, suffix); bFound {
			title = titleCuted
			bCompleted = true
		} else if strings.Contains(title, ")"+suffix) {
			strings.ReplaceAll(title, ")"+suffix, ") ")
			bCompleted = true
		}
	}

	return title, bCompleted
}
