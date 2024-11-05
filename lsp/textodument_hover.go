package lsp

type HoverRequest struct {
	Request
	Params HoverParams `json:"hoverParams"`
}

type HoverParams struct {
	TextDocumentPositionParams
}
