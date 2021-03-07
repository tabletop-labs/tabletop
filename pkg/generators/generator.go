package generators

import (
	"fmt"
	"path"
)

// Execute the steps of a generator in order
func Execute(g Generator) error {
	steps, err := g.Steps()
	if err != nil {
		return fmt.Errorf("Execute: failed to get steps: %w", err)
	}

	for _, s := range steps {
		err := s.Execute()
		if err != nil {
			return fmt.Errorf("Execute: failed to execute step: %w", err)
		}
	}

	return nil
}

type Generator interface {
	Steps() ([]Step, error)
}

type Step interface {
	Execute() error
}

// The project the generator belongs to
type Project struct {
	Path string
}

// A bounded context within an application
type ApplicationContext struct {
	App     string
	Context string
}

func (c *ApplicationContext) AppContextPath() string {
	return path.Join("internal", "apps", c.App, c.Context)
}
