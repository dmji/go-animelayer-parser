//go:build !examgenerator

package animelayer_test

import "context"

func GenerateInitialItemExams(testFileHtml, testFileExam string, ctx context.Context, params TestGetItemParams) error {
	return nil
}
func GenerateInitialPageExams(testFileHtml, testFileExam string, ctx context.Context, params TestParseFirstPageParams) error {
	return nil
}
