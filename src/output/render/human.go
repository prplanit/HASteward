package render

import (
	"fmt"
	"io"
	"strings"
	"time"
)

const bannerWidth = 60

// Human renders output in a human-readable format with box-drawing and status icons.
type Human struct {
	Out io.Writer
}

// Banner prints a prominent section header.
func (h *Human) Banner(title string) {
	line := strings.Repeat("=", bannerWidth)
	fmt.Fprintln(h.Out)
	fmt.Fprintln(h.Out, line)
	fmt.Fprintf(h.Out, "  %s\n", title)
	fmt.Fprintln(h.Out, line)
	fmt.Fprintln(h.Out)
}

// Header prints a formatted header line for HASteward startup.
func (h *Human) Header(engine, mode, clusterName, namespace string) {
	fmt.Fprintf(h.Out, "=== HASteward / %s (%s) ===\n", strings.ToUpper(engine), mode)
	if engine == "galera" {
		fmt.Fprintf(h.Out, "MariaDB: %s\n", clusterName)
	} else {
		fmt.Fprintf(h.Out, "Cluster: %s\n", clusterName)
	}
	fmt.Fprintf(h.Out, "Namespace: %s\n", namespace)
	fmt.Fprintf(h.Out, "Timestamp: %s\n", time.Now().UTC().Format(time.RFC3339))
}

// Section prints a subsection divider.
func (h *Human) Section(title string) {
	fmt.Fprintf(h.Out, "--- %s ---\n", title)
}

// Field prints a labeled value.
func (h *Human) Field(label, value string) {
	fmt.Fprintf(h.Out, "%s: %s\n", label, value)
}

// Bullet prints a bulleted item with optional indentation.
func (h *Human) Bullet(indent int, format string, args ...any) {
	prefix := strings.Repeat("  ", indent)
	fmt.Fprintf(h.Out, "%s- %s\n", prefix, fmt.Sprintf(format, args...))
}

// Info prints an informational line prefixed with >> (recommendation style).
func (h *Human) Info(format string, args ...any) {
	fmt.Fprintf(h.Out, "  >> %s\n", fmt.Sprintf(format, args...))
}

// Success prints a success message.
func (h *Human) Success(format string, args ...any) {
	fmt.Fprintf(h.Out, "[OK] %s\n", fmt.Sprintf(format, args...))
}

// Warn prints a warning message.
func (h *Human) Warn(format string, args ...any) {
	fmt.Fprintf(h.Out, "[WARN] %s\n", fmt.Sprintf(format, args...))
}

// Fail prints a failure message.
func (h *Human) Fail(format string, args ...any) {
	fmt.Fprintf(h.Out, "[FAIL] %s\n", fmt.Sprintf(format, args...))
}

// Complete prints a completion message.
func (h *Human) Complete(msg string) {
	fmt.Fprintf(h.Out, "=== %s ===\n", msg)
}

// SuggestedCommands prints suggested repair commands for triage output.
func (h *Human) SuggestedCommands(engine, clusterName, namespace string) {
	fmt.Fprintln(h.Out)
	h.Section("Suggested Commands")
	fmt.Fprintln(h.Out, "Repair all unhealthy instances:")
	fmt.Fprintf(h.Out, "  hasteward repair -e %s -c %s -n %s --backups-path /backups\n", engine, clusterName, namespace)
	fmt.Fprintln(h.Out)
	fmt.Fprintln(h.Out, "Repair a specific instance:")
	fmt.Fprintf(h.Out, "  hasteward repair -e %s -c %s -n %s --instance <N> --backups-path /backups\n", engine, clusterName, namespace)
}

// FormatBytes returns a human-readable byte size.
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
