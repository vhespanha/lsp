package analysis

type State struct {
	// Map filenames to file contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(document, text string) {
	s.Documents[document] = text
}
