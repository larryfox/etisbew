package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/ini.v1"
)

type Page struct {
	Title    string
	Template string
	Body     template.HTML
	Date     *time.Time
}

const delimiter = "~~~\n"

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.DefinitionList,
		extension.Footnote,
		extension.Strikethrough,
		extension.Table,
		extension.Typographer,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

func NewPage(file string) (page Page, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("could not open file: %v", err)
	}

	head, body := split(data)

	var buf bytes.Buffer
	if err = md.Convert(body, &buf); err != nil {
		err = fmt.Errorf("could not parse page body: %v", err)
	}

	page.Body = template.HTML(buf.String())
	page.Template = "default.html"

	err = ini.MapTo(&page, head)
	if err != nil {
		err = fmt.Errorf("could not parse page header: %v", err)
	}
	return
}

func split(input []byte) ([]byte, []byte) {
	if !bytes.HasPrefix(input, []byte(delimiter)) {
		return nil, input
	}
	s := bytes.SplitN(input, []byte(delimiter), 3)
	if len(s) == 3 {
		return s[1], s[2]
	}
	return nil, input
}
