package harvester

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSourceRegistry(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "allowlist.txt")
	content := "# allowed prefixes\nhttps://raw.githubusercontent.com/\nhttps://example.com/sub/\n"
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatalf("write file: %v", err)
	}

	r, err := LoadSourceRegistry(file)
	if err != nil {
		t.Fatalf("load registry: %v", err)
	}
	if !r.Allowed("https://raw.githubusercontent.com/a/b") {
		t.Fatalf("expected source allowed")
	}
	if r.Allowed("https://evil.example/path") {
		t.Fatalf("expected source blocked")
	}
}
