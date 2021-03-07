package appcontext

import (
	"fmt"
	"path"

	"github.com/fremantle-industries/tabletop/pkg/generators"
)

type ProcessorGenerator struct {
	generators.Project
	generators.ApplicationContext
	Name string
	Type string
}

func (p *ProcessorGenerator) Steps() ([]generators.Step, error) {
	tplData := struct {
		App           string
		ProcessorName string
		ContextName   string
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

func (p *ProcessorGenerator) SourceTemplatePath() string {
	return path.Join("processor", fmt.Sprintf("%s_processor.go.tpl", p.Type))
}

func (p *ProcessorGenerator) OutTemplatePath() string {
	return path.Join(p.Project.Path, p.AppContextPath(), fmt.Sprintf("%s_processor.go", p.Name))
}

func (p *ProcessorGenerator) SourceTestTemplatePath() string {
	return path.Join("processor", fmt.Sprintf("%s_processor_test.go.tpl", p.Type))
}

func (p *ProcessorGenerator) OutTestTemplatePath() string {
	return path.Join(p.Project.Path, p.AppContextPath(), fmt.Sprintf("%s_processor_test.go", p.Name))
}
