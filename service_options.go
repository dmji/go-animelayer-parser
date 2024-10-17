package animelayer

type ServiceOptionsApplier func(s *service)

func WithNoteClassOverride(noteElem, noteClass string) ServiceOptionsApplier {
	return func(s *service) {
		s.parser.NotePlaintTextElementInterceptor = noteElem
		s.parser.NotePlaintTextElementClassInterceptor = noteClass
	}
}
