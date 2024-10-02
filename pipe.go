package animelayer

import (
	"net/http"
)

type logger interface {
	Infow(msg string, keys ...interface{})
	Errorw(msg string, keys ...interface{})
}

type loggerPlaceholder struct{}

func (l *loggerPlaceholder) Infow(msg string, keys ...interface{})  {}
func (l *loggerPlaceholder) Errorw(msg string, keys ...interface{}) {}

func (s *pipe) SetLogger(l logger) {
	s.logger = l
}

type pipe struct {
	client *http.Client
	logger logger
}

func New(client *http.Client) *pipe {
	return &pipe{
		client: client,
		logger: &loggerPlaceholder{},
	}

}
