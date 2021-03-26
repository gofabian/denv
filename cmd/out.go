package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/mattn/go-colorable"
)

type PrefixWriter struct {
	prefix    string
	writer    io.Writer
	atNewLine bool
}

func NewPrefixWriter(prefix string, file *os.File) *PrefixWriter {
	return &PrefixWriter{
		prefix:    prefix,
		writer:    colorable.NewColorable(file),
		atNewLine: true,
	}
}

func (w *PrefixWriter) Write(p []byte) (int, error) {
	var err error
	buf := bytes.NewBuffer(p)

	for {
		line, err := buf.ReadBytes('\n')

		if len(line) > 0 {
			if w.atNewLine {
				w.writer.Write([]byte(w.prefix))
			}
			w.writer.Write(line)
			w.atNewLine = (err == nil)
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}

	unread := len(p) - buf.Len()
	return unread, err
}
