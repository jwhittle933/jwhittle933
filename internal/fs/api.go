package fs

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	filename string
	f        *os.File
	staging  *os.File
}

func Open(filename string) (*File, error) {
	ext := filepath.Ext(filename)
	if ext != ".md" {
		return nil, errors.New("file must be markdown with ext .md")
	}

	final, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	staging, err := os.Create(previewFile(filename, ext))
	if err != nil {
		return nil, err
	}

	return &File{filename, final, staging}, nil
}

func (f *File) Write(content []byte) (int, error) {
	return f.staging.Write(content)
}

func (f *File) Flush() error {
	if _, err := os.Stat(f.filename); err != nil {
		return err
	}

	if _, err := io.Copy(f.staging, f.f); err != nil {
		return err
	}

	if err := os.Remove(previewFile(f.filename, ".md")); err != nil {
		return err
	}

	return nil
}

func (f *File) Close() error {
	f.f.Close()
	f.staging.Close()
	return nil
}

func previewFile(filename, ext string) string {
	return fmt.Sprintf("%s_preview.md", strings.Replace(filename, ext, "", -1))
}