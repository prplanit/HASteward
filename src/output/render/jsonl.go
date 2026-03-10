package render

import (
	"encoding/json"
	"io"
	"sync"
)

// JSONL renders newline-delimited JSON (one compact JSON object per line).
// Thread-safe for concurrent event emission.
type JSONL struct {
	Out io.Writer
	mu  sync.Mutex
}

// Emit writes a single value as a compact JSON line.
func (jl *JSONL) Emit(v any) error {
	jl.mu.Lock()
	defer jl.mu.Unlock()

	enc := json.NewEncoder(jl.Out)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
