package analysis

import (
	"fmt"
	"log"
	"strings"

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

func (s *State) TextDocumentCodeAction(
	id int,
	uri string,
) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "Foo")
		if idx >= 0 {
			fooToBar := map[string][]lsp.TextEdit{}
			fooToBar[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("Foo")),
					NewText: "Bar",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Foo to Bar",
				Edit:  &lsp.WorkspaceEdit{Changes: fooToBar},
			})
			fooToBaz := map[string][]lsp.TextEdit{}
			fooToBaz[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("Foo")),
					NewText: "Baz",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Foo to Baz",
				Edit:  &lsp.WorkspaceEdit{Changes: fooToBaz},
			})
		}
	}
	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
	return response
}

func (s *State) TextDocumentCompletion(
	id int,
	uri string,
) lsp.CompletionResponse {
	items := []lsp.CompletionItem{{
		Label:         "Melange",
		Detail:        "It must flow...",
		Documentation: "Melange, often referred to as 'the spice', is the fictional psychedelic drug central to the Dune series of science fiction novels by Frank Herbert and derivative works.",
	}}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
