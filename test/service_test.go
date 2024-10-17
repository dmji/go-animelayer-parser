package animelayer_test

import (
	"net/http"
	"testing"

	"github.com/dmji/go-animelayer-parser"
)

func toParser(p animelayer.ItemProvider) animelayer.ItemProvider {
	return p
}

func TestServiceInterfaces(t *testing.T) {

	client := animelayer.New(animelayer.NewClientWrapper(&http.Client{}))

	_ = toParser(client)

}
