package types

type EditorChange struct {
	Type   string     `json:"type"`
	Path   string     `json:"path"`
	Change ChangeData `json:"change"`
}

type ChangeData struct {
	Type      string      `json:"type"`
	Range     ChangeRange `json:"range"`
	Text      string      `json:"text"`
	Timestamp int64       `json:"timestamp"`
}

type ChangeRange struct {
	StartLineNumber int `json:"startLineNumber"`
	StartColumn     int `json:"startColumn"`
	EndLineNumber   int `json:"endLineNumber"`
	EndColumn       int `json:"endColumn"`
}
