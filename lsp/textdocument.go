package lsp

type TextDocumentItem struct {
	// +---------------------------------------------------------+
	// | The text document's URI.                                |
	// +---------------------------------------------------------+
	URI string `json:"uri"`

	// +---------------------------------------------------------+
	// | The text document's language identifier.                |
	// +---------------------------------------------------------+
	LanguageID string `json:"languageId"`

	// +---------------------------------------------------------+
	// | The version number of this document.                    |
	// +---------------------------------------------------------+
	Version int `json:"version"`

	// +---------------------------------------------------------+
	// | The content of the opened text document.                |
	// +---------------------------------------------------------+
	Text string `json:"text"`
}
