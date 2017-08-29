package templates

import (
	"runtime"
	"path"
)

// TemplateDir get template directory
func TemplateDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}
