package fs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	filename string
	final    *os.File
}

func Open(filename string) (*File, error) {
	ext := filepath.Ext(filename)
	if ext != ".md" {
		return nil, errors.New("file must be markdown with ext .md")
	}

	final, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &File{filename, final}, nil
}

func (f *File) Write(content []byte) (int, error) {
	return f.final.Write(content)
}

func (f *File) Flush() error {
	preview := previewFile(f.filename, ".md")
	if _, err := os.Stat(preview); err != nil {
		return err
	}

	old, err := ioutil.ReadAll(f.final)
	if err != nil {
		return err
	}

	fmt.Println("README.md >>>>>", string(old))

	if err := os.Remove("README.md"); err != nil {
		return err
	}

	if err := os.Rename(preview, "README.md"); err != nil {
		return err
	}

	return nil
}

func (f *File) Close() error {
	return f.final.Close()
}

func previewFile(filename, ext string) string {
	return fmt.Sprintf("%s_preview.md", strings.Replace(filename, ext, "", -1))
}
