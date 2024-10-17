package animelayer

type ServiceOptionsApplier func(s *service)

func WithNoteClassOverride(noteElem, noteClass string) ServiceOptionsApplier {
	return func(s *service) {
		s.parserDetailedItems.NotePlaintTextElementInterceptor = noteElem
		s.parserDetailedItems.NotePlaintTextElementClassInterceptor = noteClass
	}
}
