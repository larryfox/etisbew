package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var logger = log.New(os.Stdout, "", 0)

const outputDir = "output"
const templateDir = "templates"

func main() {
	logger.Println("┌╴building website")

	err := os.RemoveAll(outputDir)
	if err != nil {
		logger.Fatalf("└╴could not clean output dir: %v", err)
	}

	files, err := findFiles(".", ".md")
	if err != nil {
		logger.Fatalf("└╴could not read files: %v", err)
	}

	templates, err := findFiles(templateDir, ".html")
	if err != nil {
		logger.Fatalf("└╴could not read templates: %v", err)
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		logger.Fatalf("└╴could not parse templates: %v", err)
	}

	for _, file := range files {
		dest := "/" + removeExt(file)
		if file == "index.md" {
			dest = "/"
		}
		log := fmt.Sprintf("%s (%s)\n", dest, file)

		err := processFile(tmpl, file)
		if err != nil {
			logger.Printf("└┬╴%s", log)
			logger.Fatalf(" └╴%v", err)
		}
		logger.Printf("├─╴%s", log)
	}

	logger.Println("└╴all done")
}

func processFile(tmpl *template.Template, src string) error {
	var dest string
	if src == "index.md" {
		dest = filepath.Join(outputDir, replaceExt(src, ".html"))
	} else {
		dest = filepath.Join(outputDir, removeExt(src), "index.html")
	}

	page, err := NewPage(src)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = os.MkdirAll(filepath.Dir(dest), 0777)
	if err != nil {
		return fmt.Errorf("could not make dir: %v", err)
	}

	file, err := os.Create(dest)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("could not make file: %v", err)
	}
	err = tmpl.ExecuteTemplate(file, page.Template, page)
	if err != nil {
		return fmt.Errorf("could not write template: %v", err)
	}
	return nil
}

func removeExt(path string) string {
	ext := filepath.Ext(path)
	return path[0 : len(path)-len(ext)]
}

func replaceExt(path, ext string) string {
	return removeExt(path) + ext
}

func findFiles(dir string, ext string) (files []string, err error) {
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != "." && d.IsDir() && strings.HasPrefix(path, ".") {
			return fs.SkipDir
		}
		if !d.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	return
}
