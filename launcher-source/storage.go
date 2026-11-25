package main

import (
	"path/filepath"
	"strings"
)

// Preserve the exact cache location of managed v10 builds and recognize builds
// installed by v11, even if the application is subsequently installed elsewhere.
func managedBuildDir(exe, key string) string {
	if !filepath.IsAbs(exe) {
		return ""
	}
	release := filepath.Dir(exe)
	build := filepath.Dir(release)
	// Cargo may use an explicit target triple between build and release.
	if !strings.HasPrefix(filepath.Base(build), "build-") {
		build = filepath.Dir(build)
	}
	versions := filepath.Dir(build)
	engine := filepath.Dir(versions)
	container := filepath.Base(filepath.Dir(engine))
	if strings.EqualFold(filepath.Base(release), "release") && strings.HasPrefix(filepath.Base(build), "build-") && strings.EqualFold(filepath.Base(versions), "versions") && strings.EqualFold(filepath.Base(engine), key) && (strings.EqualFold(container, "motors") || strings.EqualFold(container, "ports")) {
		return build
	}
	return ""
}
