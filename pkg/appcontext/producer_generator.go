package appcontext

import (
	"fmt"
	"path"

	"github.com/fremantle-industries/tabletop/pkg/generators"
)

type ProducerGenerator struct {
	generators.Project
	generators.ApplicationContext
	Name string
	Type string
}

func (p *ProducerGenerator) Steps() ([]generators.Step, error) {
	tplData := struct {
		App          string
		ProducerName string
		ContextName  string
	}{p.App, p.Name, p.Context}

	steps := []generators.Step{
		&generators.EnsureDirStep{
			Path: path.Join(p.Project.Path, p.AppContextPath()),
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: p.SourceTemplatePath(),
			OutPath:    p.OutTemplatePath(),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: p.SourceTestTemplatePath(),
			OutPath:    p.OutTestTemplatePath(),
			Data:       tplData,
		},
	}

	return steps, nil
}

func (p *ProducerGenerator) SourceTemplatePath() string {
	return path.Join("producer", fmt.Sprintf("%s_producer.go.tpl", p.Type))
}

func (p *ProducerGenerator) OutTemplatePath() string {
	return path.Join(p.Project.Path, p.AppContextPath(), fmt.Sprintf("%s_producer.go", p.Name))
}

func (p *ProducerGenerator) SourceTestTemplatePath() string {
	return path.Join("producer", fmt.Sprintf("%s_producer_test.go.tpl", p.Type))
}

func (p *ProducerGenerator) OutTestTemplatePath() string {
	return path.Join(p.Project.Path, p.AppContextPath(), fmt.Sprintf("%s_producer_test.go", p.Name))
}
