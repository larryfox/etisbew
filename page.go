package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/ini.v1"
)

type Page struct {
	Title    string
	Template string
	Body     template.HTML
}

const delimiter = "~~~\n"

func NewPage(file string) (page Page, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("could not open file: %v", err)
	}

	head, body := split(data)

	page.Template = "default.html"
	page.Body = template.HTML(string(blackfriday.Run(body)))

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
