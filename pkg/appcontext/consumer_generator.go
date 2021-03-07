package appcontext

import (
	"fmt"
	"path"

	"github.com/fremantle-industries/tabletop/pkg/generators"
)

type ConsumerGenerator struct {
	generators.Project
	generators.ApplicationContext
	Name string
	Type string
}

func (c *ConsumerGenerator) Steps() ([]generators.Step, error) {
	tplData := struct {
		App          string
		ConsumerName string
		ContextName  string
	}{c.App, c.Name, c.Context}

	steps := []generators.Step{
		&generators.EnsureDirStep{
			Path: path.Join(c.Project.Path, c.AppContextPath()),
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: c.SourceTemplatePath(),
			OutPath:    c.OutTemplatePath(),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: c.SourceTestTemplatePath(),
			OutPath:    c.OutTestTemplatePath(),
			Data:       tplData,
		},
	}

	return steps, nil
}

func (c *ConsumerGenerator) SourceTemplatePath() string {
	return path.Join("consumer", fmt.Sprintf("%s_consumer.go.tpl", c.Type))
}

func (c *ConsumerGenerator) OutTemplatePath() string {
	return path.Join(c.Project.Path, c.AppContextPath(), fmt.Sprintf("%s_consumer.go", c.Name))
}

func (c *ConsumerGenerator) SourceTestTemplatePath() string {
	return path.Join("consumer", fmt.Sprintf("%s_consumer_test.go.tpl", c.Type))
}

func (c *ConsumerGenerator) OutTestTemplatePath() string {
	return path.Join(c.Project.Path, c.AppContextPath(), fmt.Sprintf("%s_consumer_test.go", c.Name))
}
