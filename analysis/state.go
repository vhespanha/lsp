package analysis

import (
	"fmt"
	"log"

	"git.sr.ht/~vhespanha/lsp/lsp"
)

type State struct {
	// Map filenames to file contents
	Documents map[string]string
	logger    *log.Logger
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) SetLogger(logger *log.Logger) {
	s.logger = logger
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]
	if s.logger != nil {
		s.logger.Printf(
			"Hover requested - URI: %s, Found doc length: %d",
			uri,
			len(document),
		)
		s.logger.Printf("Available documents: %v", s.Documents)
	}
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d.", uri, len(document)),
		},
	}
}

func (s *State) Definition(
	id int,
	uri string,
	position lsp.Position,
) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
