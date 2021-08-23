package readme

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	filename string
	f        *os.File
	staging  *os.File
}

func Open(filename string) (*File, error) {
	if ext := filepath.Ext(filename); ext != ".md" {
		return nil, errors.New("file must be markdown with ext .md")
	}

	final, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	staging, err := os.Create(fmt.Sprintf("%s_preview.md", filename))
	if err != nil {
		return nil, err
	}

	return &File{filename, final, staging}, nil
}

func (f *File) Write(content []byte) (int, error) {
	return f.staging.Write(content)
}

func (f *File) Close() error {
	f.f.Close()
	f.staging.Close()
	return nil
}
