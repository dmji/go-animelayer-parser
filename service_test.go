package animelayer_test

import (
	"net/http"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func toParser(p animelayer.Parser) animelayer.Parser {
	return p
}

func toParserPipeline(p animelayer.ParserPipeline) animelayer.ParserPipeline {
	return p
}

func TestServiceInterfaces(t *testing.T) {

	client := animelayer.New(animelayer.NewHttpClientWrapper(&http.Client{}))

	_ = toParser(client)
	_ = toParserPipeline(client)

}
