package internal

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func copyFiles(src string, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, relPath)
		if d.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}
		return CopyFile(path, targetPath)
	})
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}
	return nil
}

func GenerateHtml(content DynamicContent) error {
	err := copyFiles("static", "public")
	if err != nil {
		return err
	}

	file, err := os.Create("public/index.html")
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(file, "index.gohtml", content)
	if err != nil {
		return err
	}
	return nil
}
