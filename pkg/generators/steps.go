package generators

import (
	"embed"
	"fmt"
	"os"
	"path"
	"strings"

	"text/template"
)

var (
	//go:embed all:templates
	embeddedTpls embed.FS

	// Directory containing the template files
	tplDir = "templates"

	// Pipeline functions available to generator templates
	fMap = template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
		"Title":   strings.Title,
	}
)

// Hydrates a template with data and writes the output to a file
type EmbeddedTemplateStep struct {
	SourcePath string
	OutPath    string
	Data       interface{}
}

func (s *EmbeddedTemplateStep) Execute() error {
	tplPath := path.Join(tplDir, s.SourcePath)
	tpl, err := template.New(path.Base(s.SourcePath)).Funcs(fMap).ParseFS(embeddedTpls, tplPath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(s.OutPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fmt.Printf("creating template %s\n", s.OutPath)
	return tpl.Execute(outFile, s.Data)
}

type EnsureDirStep struct {
	Path string
}

func (s *EnsureDirStep) Execute() error {
	fmt.Printf("creating directory %s\n", s.Path)
	return os.MkdirAll(s.Path, 0755)
}
