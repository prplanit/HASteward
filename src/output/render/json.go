package render

import (
	"encoding/json"
	"io"
)

// JSON renders a single JSON document to the writer.
type JSON struct {
	Out io.Writer
}

// Render writes v as a single pretty-printed JSON document.
func (j *JSON) Render(v any) error {
	enc := json.NewEncoder(j.Out)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// RenderCompact writes v as a single compact JSON document (no indentation).
func (j *JSON) RenderCompact(v any) error {
	enc := json.NewEncoder(j.Out)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
