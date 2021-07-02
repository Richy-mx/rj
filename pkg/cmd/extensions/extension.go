package extensions

import (
	"path/filepath"
	"strings"
)

type Extension struct {
	path      string
	url       string
	updatable bool
}

func (e *Extension) Name() string {
	return strings.TrimPrefix(filepath.Base(e.path), "gh-")
}

func (e *Extension) Path() string {
	return e.path
}

func (e *Extension) URL() string {
	return e.url
}

func (e *Extension) Updatable() bool {
	return e.updatable
}
