package logsx

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestRotatingWriterRotatesEvery500Lines(t *testing.T) {
	dir := t.TempDir()
	w, err := NewRotatingWriter(dir, "app.log", 500, 5)
	if err != nil {
		t.Fatalf("new writer: %v", err)
	}
	defer w.Close()

	for i := range 1200 {
		if _, err := fmt.Fprintf(w, "line %d\n", i); err != nil {
			t.Fatalf("write: %v", err)
		}
	}

	rotated, _ := filepath.Glob(filepath.Join(dir, "app-*.log"))
	if len(rotated) < 2 {
		t.Fatalf("expected >=2 rotated files, got %d", len(rotated))
	}
	if got := countLines(filepath.Join(dir, "app.log")); got > 500 {
		t.Fatalf("active file has %d lines, want <=500", got)
	}
}

func TestRotatingWriterPrunesToKeep(t *testing.T) {
	dir := t.TempDir()
	w, err := NewRotatingWriter(dir, "app.log", 10, 3)
	if err != nil {
		t.Fatalf("new writer: %v", err)
	}
	defer w.Close()

	for i := range 100 {
		fmt.Fprintf(w, "line %d\n", i)
	}

	rotated, _ := filepath.Glob(filepath.Join(dir, "app-*.log"))
	if len(rotated) > 3 {
		t.Fatalf("expected <=3 retained rotated files, got %d", len(rotated))
	}
}
