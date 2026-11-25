package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

// Recognizes only known engines by directory names and expected executables.
// No executable is launched, copied, or modified by this scan.
func identifyPort(path string) int {
	base := strings.ToLower(filepath.Base(path))
	if base != "victory.exe" && base != "xcl.exe" && base != "ztrx.exe" && base != "hyphen.exe" && base != "vitoriabr.exe" {
		return -1
	}
	if base == "xcl.exe" {
		return 3
	}
	if base == "ztrx.exe" {
		return 2
	}
	if base == "hyphen.exe" {
		return 1
	}
	if base == "vitoriabr.exe" {
		return 5
	}
	for dir := filepath.Dir(path); ; dir = filepath.Dir(dir) {
		n := strings.ToLower(filepath.Base(dir))
		n = strings.NewReplacer("_", "-", " ", "-").Replace(n)
		switch {
		case strings.Contains(n, "belobug"):
			return 4
		case strings.Contains(n, "vitoriabr"):
			return 5
		case n == "ztrx" || strings.HasPrefix(n, "ztrx-"):
			return 2
		case n == "hyphen" || strings.HasPrefix(n, "hyphen-") || strings.HasPrefix(n, "victory-fork"):
			return 1
		case n == "xcl" || strings.HasPrefix(n, "xcl-"):
			return 3
		case n == "victory" || n == "original" || strings.HasPrefix(n, "victory-main") || strings.HasPrefix(n, "victory-original"):
			return 0
		}
		if filepath.Dir(dir) == dir {
			break
		}
	}
	return -1
}
func discoverPorts(root string) ([6][]string, error) {
	var result [6][]string
	count := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if path == root {
				return err
			}
			return nil
		}
		count++
		if count > 60000 {
			return fmt.Errorf("Limite de exploracion alcanzado; elige una carpeta mas concreta")
		}
		if d.Type()&fs.ModeSymlink != 0 {
			return nil
		}
		if d.IsDir() && path != root {
			rel, _ := filepath.Rel(root, path)
			if strings.Count(rel, string(filepath.Separator)) >= 9 {
				return filepath.SkipDir
			}
			switch strings.ToLower(d.Name()) {
			case ".git", ".cargo", "node_modules", "deps", "incremental", "build", ".fingerprint", "cache", "pack", "data":
				return filepath.SkipDir
			}
		}
		if d.IsDir() {
			return nil
		}
		i := identifyPort(path)
		if i >= 0 {
			result[i] = append(result[i], path)
		}
		return nil
	})
	for i := range result {
		sort.Strings(result[i])
	}
	return result, err
}
