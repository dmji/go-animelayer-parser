package animelayer

import "strings"

type parserDetailedItems struct {
	NotePlaintTextElementInterceptor      string
	NotePlaintTextElementClassInterceptor string
}

func (p *parserDetailedItems) grabTitleWithCompletedStatus(name string) (string, bool) {
	title := cleanStringFromHtmlSymbols(name)
	bCompleted := false

	if titleCuted, bFound := strings.CutSuffix(title, " Complete"); bFound {
		title = titleCuted
		bCompleted = true
	} else {
		bFound := strings.Contains(title, ") Complete")
		if bFound {
			strings.ReplaceAll(title, ") Complete", ") ")
			bCompleted = true
		}
	}
	return title, bCompleted
}
