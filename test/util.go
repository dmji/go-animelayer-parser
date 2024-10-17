package animelayer_test

import (
	"fmt"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/joho/godotenv"
)

func init() {

	path := ".env"
	for i := range 10 {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}

}

func isSameError(got, expected error) bool {
	if got == nil && expected == nil {
		return true
	}

	if got != nil && expected != nil {
		return got.Error() == expected.Error()
	}

	return false
}

func isEqualItemMetrics(got, expected *animelayer.ItemMetrics) error {

	if got.Uploads != expected.Uploads {
		return fmt.Errorf("expected Metrics.Uploads='%s', but got='%s", expected.Uploads, got.Uploads)
	}

	if got.Downloads != expected.Downloads {
		return fmt.Errorf("expected Metrics.Downloads='%s', but got='%s", expected.Downloads, got.Downloads)
	}

	if got.FilesSize != expected.FilesSize {
		return fmt.Errorf("expected Metrics.FilesSize='%s', but got='%s", expected.FilesSize, got.FilesSize)
	}

	if got.Author != expected.Author {
		return fmt.Errorf("expected Metrics.Author='%s', but got='%s", expected.Author, got.Author)
	}

	if got.VisitorCounter != expected.VisitorCounter {
		return fmt.Errorf("expected Metrics.VisitorCounter='%s', but got='%s", expected.VisitorCounter, got.VisitorCounter)
	}

	if got.ApprovedCounter != expected.ApprovedCounter {
		return fmt.Errorf("expected Metrics.ApprovedCounter='%s', but got='%s", expected.ApprovedCounter, got.ApprovedCounter)
	}

	if got.DebugReadFromElementClass != expected.DebugReadFromElementClass {
		return fmt.Errorf("expected Metrics.ReadFromHtmlKey='%s', but got='%s", expected.DebugReadFromElementClass, got.DebugReadFromElementClass)
	}

	return nil
}

func isEquaTime(got, expected *time.Time) bool {

	if got == nil && expected == nil {
		return true
	}

	if got == nil || expected == nil {
		return false
	}

	return got.Equal(*expected)
}

func isEqualItemUpdate(got, expected *animelayer.ItemUpdate) error {

	if !isEquaTime(got.UpdatedDate, expected.UpdatedDate) {
		return fmt.Errorf("expected Updated.UpdatedDate='%s', but got='%s", expected.UpdatedDate, got.UpdatedDate)
	}

	if !isEquaTime(got.CreatedDate, expected.CreatedDate) {
		return fmt.Errorf("expected Updated.CreatedDate='%s', but got='%s", expected.CreatedDate, got.CreatedDate)
	}

	if !isEquaTime(got.SeedLastPresenceDate, expected.SeedLastPresenceDate) {
		return fmt.Errorf("expected Updated.SeedLastPresenceDate='%s', but got='%s", expected.SeedLastPresenceDate, got.SeedLastPresenceDate)
	}

	if got.DebugReadFromElementClass != expected.DebugReadFromElementClass {
		return fmt.Errorf("expected Updated.ReadFromHtmlKey='%s', but got='%s", expected.DebugReadFromElementClass, got.DebugReadFromElementClass)
	}

	return nil
}

func isEqualItem(got, expected *animelayer.Item) error {

	if got.Identifier != expected.Identifier {
		return fmt.Errorf("expected Identifier='%s', but got='%s'", expected.Identifier, got.Identifier)
	}

	if got.Title != expected.Title {
		return fmt.Errorf("expected Title='%s', but got='%s", expected.Title, got.Title)
	}

	if got.IsCompleted != expected.IsCompleted {
		return fmt.Errorf("expected IsCompleted='%v', but got='%v", expected.IsCompleted, got.IsCompleted)
	}

	if got.RefImagePreview != expected.RefImagePreview {
		return fmt.Errorf("expected RefImagePreview='%s', but got='%s", expected.RefImagePreview, got.RefImagePreview)
	}

	if got.RefImageCover != expected.RefImageCover {
		return fmt.Errorf("expected RefImageCover='%s', but got='%s", expected.RefImageCover, got.RefImageCover)
	}

	if err := isEqualItemMetrics(&got.Metrics, &expected.Metrics); err != nil {
		return err
	}

	if err := isEqualItemUpdate(&got.Updated, &expected.Updated); err != nil {
		return err
	}

	if got.Notes != expected.Notes {
		return fmt.Errorf("expected Notes='%s', but got='%s", expected.Notes, got.Notes)
	}

	return nil
}

func isValidItem(got *animelayer.Item, bTarget bool) error {

	if len(got.Identifier) == 0 {
		return fmt.Errorf("expected not empty Identifier")
	}

	if len(got.Title) == 0 {
		return fmt.Errorf("expected not empty Title")
	}

	if len(got.Metrics.DebugReadFromElementClass) == 0 {
		return fmt.Errorf("expected not empty Metrics.ReadFromHtmlKey")
	}

	if len(got.Updated.DebugReadFromElementClass) == 0 {
		return fmt.Errorf("expected not empty Updated.ReadFromHtmlKey")
	}

	if bTarget && len(got.RefImagePreview) == 0 {
		return fmt.Errorf("expected not empty RefImagePreview")
	}

	if len(got.RefImageCover) == 0 {
		return fmt.Errorf("expected not empty RefImageCover")
	}

	if len(got.Notes) == 0 {
		return fmt.Errorf("expected not empty Notes")
	}

	return nil
}
