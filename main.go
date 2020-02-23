package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// Option represents application options
type Option struct {
	Number bool `short:"n" long:"number" description:"Show contents with line numbers"`
}

// Config represents the settings for this application
type Config struct {
	Theme string `json:"theme"`
}

// CLI represents this application itself
type CLI struct {
	Config Config
}

// Cat formats file with syntax highlighting
func (c CLI) Cat(opt Option, path string) (string, error) {
	lexer := lexers.Match(path)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(c.Config.Theme)
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	r, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	iterator, err := lexer.Tokenise(nil, string(r))
	if err != nil {
		return "", err
	}

	w := new(bytes.Buffer)
	if err := formatter.Format(w, style, iterator); err != nil {
		return "", err
	}

	s := w.String()
	ss := strings.Split(s, "\n")

	contents := ""
	for i, s := range ss {
		if opt.Number {
			contents += fmt.Sprintf("%6d  ", i+1)
		}

		contents += fmt.Sprintf("%s\n", s)
	}

	return contents, nil
}
