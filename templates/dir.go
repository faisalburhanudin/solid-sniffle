package templates

import (
	"path"
	"runtime"
)

// TemplateDir get template directory
func TemplateDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}
