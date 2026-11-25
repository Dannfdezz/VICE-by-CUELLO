package main

import (
	"path/filepath"
	"testing"
)

func TestManagedStorage(t *testing.T) {
	root := t.TempDir()
	for _, folder := range []string{"motors", "ports"} {
		build := filepath.Join(root, folder, "belobug", "versions", "build-123")
		for _, exe := range []string{filepath.Join(build, "release", "victory.exe"), filepath.Join(build, "x86_64-pc-windows-msvc", "release", "victory.exe")} {
			if got := managedBuildDir(exe, "belobug"); got != build {
				t.Fatalf("%s != %s", got, build)
			}
			if got := managedBuildDir(exe, "original"); got != "" {
				t.Fatal("wrong engine")
			}
		}
	}
	if managedBuildDir(filepath.Join(root, "old", "target", "release", "victory.exe"), "belobug") != "" {
		t.Fatal("imported engine incorrectly treated as managed")
	}
}
