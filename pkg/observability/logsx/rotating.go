package logsx

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// RotatingWriter is an io.Writer that rotates its backing file every maxLines
// written lines, keeping the last `keep` rotated files.
//
// Line-count rotation is custom because the stdlib and lumberjack only rotate
// by byte size; the target policy here is a fixed line count.
type RotatingWriter struct {
	mu       sync.Mutex
	dir      string
	name     string // base file name, e.g. "app.log"
	maxLines int
	keep     int
	file     *os.File
	lines    int
}

// NewRotatingWriter opens (or creates) dir/name for appending and rotates it
// every maxLines lines, retaining keep rotated files. Existing lines in the
// file count toward the threshold so a restart cannot grow it without bound.
func NewRotatingWriter(dir, name string, maxLines, keep int) (*RotatingWriter, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create log dir %q: %w", dir, err)
	}
	w := &RotatingWriter{dir: dir, name: name, maxLines: maxLines, keep: keep}
	if err := w.open(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *RotatingWriter) open() error {
	path := filepath.Join(w.dir, w.name)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open log file %q: %w", path, err)
	}
	w.file = f
	w.lines = countLines(path)
	return nil
}

func (w *RotatingWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	n, err := w.file.Write(p)
	w.lines += strings.Count(string(p[:n]), "\n")
	if err != nil {
		return n, err
	}
	if w.lines >= w.maxLines {
		if rerr := w.rotate(); rerr != nil {
			return n, rerr
		}
	}
	return n, nil
}

func (w *RotatingWriter) rotate() error {
	if err := w.file.Close(); err != nil {
		return fmt.Errorf("close log file: %w", err)
	}
	ext := filepath.Ext(w.name)
	base := strings.TrimSuffix(w.name, ext)
	stamp := time.Now().Format("20060102-150405")
	// Second-granularity stamps collide under load (>maxLines lines in one
	// second); disambiguate so a rotation never overwrites a prior one.
	rotated := filepath.Join(w.dir, fmt.Sprintf("%s-%s%s", base, stamp, ext))
	for i := 1; fileExists(rotated); i++ {
		rotated = filepath.Join(w.dir, fmt.Sprintf("%s-%s-%d%s", base, stamp, i, ext))
	}
	if err := os.Rename(filepath.Join(w.dir, w.name), rotated); err != nil {
		return fmt.Errorf("rotate log file: %w", err)
	}
	w.prune(base, ext)
	w.lines = 0
	return w.open()
}

func (w *RotatingWriter) prune(base, ext string) {
	matches, err := filepath.Glob(filepath.Join(w.dir, base+"-*"+ext))
	if err != nil || len(matches) <= w.keep {
		return
	}
	sort.Strings(matches) // timestamp names sort chronologically
	for _, old := range matches[:len(matches)-w.keep] {
		_ = os.Remove(old)
	}
}

func (w *RotatingWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.file.Close()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func countLines(path string) int {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	return strings.Count(string(b), "\n")
}
