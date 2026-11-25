package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscovery(t *testing.T) {
	d := t.TempDir()
	for _, p := range []string{"sources/VIctory-main/target/release/victory.exe", "sources/VIctory-Belobug-AMD-main/target/release/victory.exe", "sources/ztrx-main/target/x86_64-pc-windows-msvc/release/victory.exe", "sources/victory-fork-main/target/release/victory.exe", "sources/VItoriaBr-main/target/release/victory.exe", "xcl/xcl.exe", "other/victory.exe", "sources/VIctory-main/target/release/deps/victory.exe"} {
		p = filepath.Join(d, filepath.FromSlash(p))
		if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
			t.Fatal(e)
		}
		if e := os.WriteFile(p, []byte("test only"), 0600); e != nil {
			t.Fatal(e)
		}
	}
	r, e := discoverPorts(d)
	if e != nil {
		t.Fatal(e)
	}
	for i, p := range r {
		if len(p) != 1 {
			t.Fatalf("engine %d: %v", i, p)
		}
	}
}
