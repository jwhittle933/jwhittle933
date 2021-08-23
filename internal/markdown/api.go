package markdown

import (
	"fmt"
	"io"
	"strings"
)

func H1(w io.Writer, s string) error {
	return writeError(w.Write(heading(1, s)))
}

func H2(w io.Writer, s string) error {
	return writeError(w.Write(heading(2, s)))
}

func H3(w io.Writer, s string) error {
	return writeError(w.Write(heading(3, s)))
}

func H5(w io.Writer, s string) error {
	return writeError(w.Write(heading(4, s)))
}

func Anchor(w io.Writer, href, children string) error {
	return writeError(w.Write([]byte(fmt.Sprintf(`<a href="%s">
  %s
</a>
`, href, children))))
}

func Link(w io.Writer, text, href string) error {
	return writeError(w.Write([]byte(fmt.Sprintf(`![%s](%s)
`, text, href))))
}

func Img(src, alt string) string {
	return fmt.Sprintf(`<img src="%s" alt="%s">`, src, alt)
}

func HR(w io.Writer) error {
	return writeError(w.Write([]byte("<hr>\n")))
}

func BR(w io.Writer) error {
	return writeError(w.Write([]byte("<br>\n")))
}

func NewLine(w io.Writer) error {
	return writeError(w.Write([]byte("\n")))
}

func LI(w io.Writer, content string) error {
	return writeError(w.Write([]byte(fmt.Sprintf("- %s\n", content))))
}

func Italics(s string) string {
	return "_" + s + "_"
}

func Bold(s string) string {
	return "__" + s + "__"
}


func heading(level int, s string) []byte {
	return []byte(fmt.Sprintf("%s %s\n", strings.Repeat("#", level), s))
}

func writeError(_ int, err error) error {
	return err
}
