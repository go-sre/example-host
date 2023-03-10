package resource

import (
	"embed"
	"fmt"
	"github.com/gotemplates/core/runtime"
	"io/fs"
)

//go:embed fs/*
var f embed.FS

func ReadFile(name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid argument : file name is empty")
	}
	return fs.ReadFile(f, name)
}

func ReadMap(name string) (map[string]string, error) {
	buf, err := ReadFile(name)
	if err != nil {
		return nil, err
	}
	return runtime.ParseMap(buf)
}
