package animelayer

type parser struct {
	logger logger
}

func newParser(logger logger) *parser {
	return &parser{logger: logger}
}
