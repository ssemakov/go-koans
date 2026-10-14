package main

// Write and read a file
// Implement storeText and loadText with os.WriteFile and os.ReadFile. Overwrite existing
// contents, use mode 0600 when creating a file, and return filesystem errors.

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func storeText(filename, text string) error {
	return nil // TODO: write the file and return its error.
}
func loadText(filename string) (string, error) {
	return "", nil // TODO: read the file and preserve its error.
}

func TestKoan(t *testing.T) {
	directory := t.TempDir()
	filename := filepath.Join(directory, "sample.txt")
	for _, text := range []string{"Go", "", "new contents", "é世"} {
		if err := storeText(filename, text); err != nil {
			t.Fatalf("storeText: %v", err)
		}
		got, err := loadText(filename)
		if err != nil || got != text {
			t.Errorf("loadText=(%q,%v); want (%q,nil)", got, err, text)
		}
	}
	if _, err := loadText(filepath.Join(directory, "missing.txt")); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("missing file error=%v", err)
	}
	if err := storeText(filepath.Join(directory, "missing-directory", "file.txt"), "data"); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("missing directory error=%v", err)
	}
}
