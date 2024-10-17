package animelayer

// ServiceOptionsApplier - type of options applier
type ServiceOptionsApplier func(s *service)

// WithNoteClassOverride - Provide customize for notes elements, wrap plaint text into noteElem
// if need to specify class of elem add 'noteClass' as well else left empty
func WithNoteClassOverride(noteElem, noteClass string) ServiceOptionsApplier {
	return func(s *service) {
		s.parser.NotePlaintTextElementInterceptor = noteElem
		s.parser.NotePlaintTextElementClassInterceptor = noteClass
	}
}
