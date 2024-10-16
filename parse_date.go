package animelayer

import (
	"strconv"
	"strings"
	"time"
)

func (p *parserDetailedItems) dateFromAnimelayerDate(t string) *time.Time {

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
