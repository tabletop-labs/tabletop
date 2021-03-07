package project

import (
	"fmt"
	"os"
	"path"

	"github.com/fremantle-industries/tabletop/pkg/appcontext"
	"github.com/fremantle-industries/tabletop/pkg/generators"
)

const (
	projectContext       = "ticks"
	projectComponentName = "tick"
	projectProducerType  = "hello_tick"
	projectProcessorType = "hello_tick"
	projectConsumerType  = "hello_tick"
)

type ProjectGenerator struct {
	Module string
	Dir    string
}

func (g *ProjectGenerator) Steps() ([]generators.Step, error) {
	var err error
	var resolvedDir string = g.Dir

	if resolvedDir == "." {
		resolvedDir, err = os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory: %w", err)
		}
	}
	project := generators.Project{
		Path: path.Join(resolvedDir, g.projectName()),
	}
	appCtxGenerator := generators.ApplicationContext{
		App:     g.projectName(),
		Context: projectContext,
	}

	acc := projectStepAccumulator{}
	acc.AppendSteps(
		&generators.EnsureDirStep{Path: project.Path},
		&generators.EnsureDirStep{Path: path.Join(project.Path, "internal", "apps", g.projectName())},
	)
	acc.AppendStepsFromGenerators(
		&projectRootGenerator{
			Project:     project,
			Module:      g.Module,
			ProjectName: g.projectName(),
		},
		&projectBinGenerator{
			Project: project,
		},
		&projectCmdGenerator{
			Project:     project,
			Module:      g.Module,
			ProjectName: g.projectName(),
		},
		&projectPkgGenerator{
			Project: project,
		},
		&projectDockerGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServiceGrafanaGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServiceLokiGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServicePrometheusGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServiceRedpandaBrokerGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServiceRedpandaConsoleGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&projectDockerServiceReverseProxyGenerator{
			Project:     project,
			ProjectName: g.projectName(),
		},
		&appcontext.ProducerGenerator{
			Project:            project,
			ApplicationContext: appCtxGenerator,
			Name:               projectComponentName,
			Type:               projectProducerType,
		},
		&appcontext.ProcessorGenerator{
			Project:            project,
			ApplicationContext: appCtxGenerator,
			Name:               projectComponentName,
			Type:               projectProcessorType,
		},
		&appcontext.ConsumerGenerator{
			Project:            project,
			ApplicationContext: appCtxGenerator,
			Name:               projectComponentName,
			Type:               projectConsumerType,
		},
	)

	return acc.Steps(), nil
}

func (g *ProjectGenerator) projectName() string {
	return path.Base(g.Module)
}

// projectRootGenerator creates the directory and files for the project /
type projectRootGenerator struct {
	Project     generators.Project
	Module      string
	ProjectName string
}

func (g *projectRootGenerator) Steps() ([]generators.Step, error) {
	tplData := struct {
		Module      string
		ProjectName string
	}{g.Module, g.ProjectName}

	steps := []generators.Step{
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "README.md.tpl"),
			OutPath:    path.Join(g.Project.Path, "README.md"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "Makefile.tpl"),
			OutPath:    path.Join(g.Project.Path, "Makefile"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "main.go.tpl"),
			OutPath:    path.Join(g.Project.Path, "main.go"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "go.mod.tpl"),
			OutPath:    path.Join(g.Project.Path, "go.mod"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "go.sum.tpl"),
			OutPath:    path.Join(g.Project.Path, "go.sum"),
			Data:       tplData,
		},
	}

	return steps, nil
}

// projectBinGenerator creates the directory and files for the project /bin
type projectBinGenerator struct {
	Project generators.Project
}

func (g *projectBinGenerator) Steps() ([]generators.Step, error) {
	binPath := path.Join(g.Project.Path, "bin")

	steps := []generators.Step{
		&generators.EnsureDirStep{Path: binPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "bin", ".gitkeep.tpl"),
			OutPath:    path.Join(binPath, ".gitkeep"),
			Data:       struct{}{},
		},
	}

	return steps, nil
}

// projectCmdGenerator creates the directory and files for the project /cmd
type projectCmdGenerator struct {
	Project     generators.Project
	Module      string
	ProjectName string
}

func (g *projectCmdGenerator) Steps() ([]generators.Step, error) {
	cmdPath := path.Join(g.Project.Path, "cmd")
	tplData := struct {
		Module      string
		ProjectName string
	}{g.Module, g.ProjectName}

	steps := []generators.Step{
		&generators.EnsureDirStep{Path: cmdPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "cmd", "root.go.tpl"),
			OutPath:    path.Join(cmdPath, "root.go"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "cmd", "commands.go.tpl"),
			OutPath:    path.Join(cmdPath, "commands.go"),
			Data:       tplData,
		},
	}

	return steps, nil
}

// projectPkgGenerator creates the directory and files for the project /cmd
type projectPkgGenerator struct {
	Project generators.Project
}

func (g *projectPkgGenerator) Steps() ([]generators.Step, error) {
	pkgPath := path.Join(g.Project.Path, "pkg")

	steps := []generators.Step{
		&generators.EnsureDirStep{Path: pkgPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "pkg", "version.go.tpl"),
			OutPath:    path.Join(pkgPath, "version.go"),
			Data:       struct{}{},
		},
	}

	return steps, nil
}

// projectDockerGenerator creates the directory and files for the project /docker
type projectDockerGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerGenerator) Steps() ([]generators.Step, error) {
	dockerPath := path.Join(g.Project.Path, "docker")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: dockerPath},
		&generators.EnsureDirStep{Path: path.Join(dockerPath, "services")},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "docker", "Dockerfile.tpl"),
			OutPath:    path.Join(dockerPath, "Dockerfile"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join("project", "docker-compose.yml.tpl"),
			OutPath:    path.Join(g.Project.Path, "docker-compose.yml"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServiceGrafanaGenerator creates the directory and files for the project /docker/service/grafana
type projectDockerServiceGrafanaGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServiceGrafanaGenerator) Steps() ([]generators.Step, error) {
	grafanaPath := path.Join(g.Project.Path, "docker", "services", "grafana")
	sourceGrafanaPath := path.Join("project", "docker", "services", "grafana")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: grafanaPath},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "plugins")},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "provisioning")},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "provisioning", "dashboards")},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "provisioning", "datasources")},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "provisioning", "notifiers")},
		&generators.EnsureDirStep{Path: path.Join(grafanaPath, "provisioning", "plugins")},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "Dockerfile.tpl"),
			OutPath:    path.Join(grafanaPath, "Dockerfile"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "grafana.ini.tpl"),
			OutPath:    path.Join(grafanaPath, "grafana.ini"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "plugins", ".gitkeep.tpl"),
			OutPath:    path.Join(grafanaPath, "plugins", ".gitkeep"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "provisioning", "dashboards", "dashboard.yml.tpl"),
			OutPath:    path.Join(grafanaPath, "provisioning", "dashboards", "dashboard.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "provisioning", "dashboards", "home.json.tpl"),
			OutPath:    path.Join(grafanaPath, "provisioning", "dashboards", "home.json"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "provisioning", "datasources", "datasource.yml.tpl"),
			OutPath:    path.Join(grafanaPath, "provisioning", "datasources", "datasource.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "provisioning", "notifiers", ".gitkeep.tpl"),
			OutPath:    path.Join(grafanaPath, "provisioning", "notifiers", ".gitkeep"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceGrafanaPath, "provisioning", "plugins", ".gitkeep.tpl"),
			OutPath:    path.Join(grafanaPath, "provisioning", "plugins", ".gitkeep"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServiceLokiGenerator creates the directory and files for the project /docker/service/loki
type projectDockerServiceLokiGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServiceLokiGenerator) Steps() ([]generators.Step, error) {
	lokiPath := path.Join(g.Project.Path, "docker", "services", "loki")
	sourceLokiPath := path.Join("project", "docker", "services", "loki")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: lokiPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceLokiPath, "Dockerfile.tpl"),
			OutPath:    path.Join(lokiPath, "Dockerfile"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServicePrometheusGenerator creates the directory and files for the project /docker/service/prometheus
type projectDockerServicePrometheusGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServicePrometheusGenerator) Steps() ([]generators.Step, error) {
	prometheusPath := path.Join(g.Project.Path, "docker", "services", "prometheus")
	sourcePrometheusPath := path.Join("project", "docker", "services", "prometheus")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: prometheusPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "Dockerfile.tpl"),
			OutPath:    path.Join(prometheusPath, "Dockerfile"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "alerts.yml.tpl"),
			OutPath:    path.Join(prometheusPath, "alerts.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "prometheus.development.yml.tpl"),
			OutPath:    path.Join(prometheusPath, "prometheus.development.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "prometheus.local.yml.tpl"),
			OutPath:    path.Join(prometheusPath, "prometheus.local.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "prometheus.production.yml.tpl"),
			OutPath:    path.Join(prometheusPath, "prometheus.production.yml"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourcePrometheusPath, "prometheus.staging.yml.tpl"),
			OutPath:    path.Join(prometheusPath, "prometheus.staging.yml"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServiceRedpandaBrokerGenerator creates the directory and files for the project /docker/service/redpanda_broker
type projectDockerServiceRedpandaBrokerGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServiceRedpandaBrokerGenerator) Steps() ([]generators.Step, error) {
	redpandaBrokerPath := path.Join(g.Project.Path, "docker", "services", "redpanda_broker")
	sourceRedpandaBrokerPath := path.Join("project", "docker", "services", "redpanda_broker")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: redpandaBrokerPath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceRedpandaBrokerPath, "Dockerfile.tpl"),
			OutPath:    path.Join(redpandaBrokerPath, "Dockerfile"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServiceRedpandaConsoleGenerator creates the directory and files for the project /docker/service/redpanda_console
type projectDockerServiceRedpandaConsoleGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServiceRedpandaConsoleGenerator) Steps() ([]generators.Step, error) {
	redpandaConsolePath := path.Join(g.Project.Path, "docker", "services", "redpanda_console")
	sourceRedpandaConsolePath := path.Join("project", "docker", "services", "redpanda_console")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: redpandaConsolePath},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceRedpandaConsolePath, "Dockerfile.tpl"),
			OutPath:    path.Join(redpandaConsolePath, "Dockerfile"),
			Data:       tplData,
		},
	}, nil
}

// projectDockerServiceReverseProxyGenerator creates the directory and files for the project /docker/service/reverse_proxy
type projectDockerServiceReverseProxyGenerator struct {
	Project     generators.Project
	ProjectName string
}

func (g *projectDockerServiceReverseProxyGenerator) Steps() ([]generators.Step, error) {
	reverseProxyPath := path.Join(g.Project.Path, "docker", "services", "reverse_proxy")
	sourceReverseProxyPath := path.Join("project", "docker", "services", "reverse_proxy")
	tplData := struct {
		ProjectName string
	}{g.ProjectName}

	return []generators.Step{
		&generators.EnsureDirStep{Path: reverseProxyPath},
		&generators.EnsureDirStep{Path: path.Join(reverseProxyPath, "templates")},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceReverseProxyPath, "Dockerfile.tpl"),
			OutPath:    path.Join(reverseProxyPath, "Dockerfile"),
			Data:       tplData,
		},
		&generators.EmbeddedTemplateStep{
			SourcePath: path.Join(sourceReverseProxyPath, "templates", "nginx.conf.tpl"),
			OutPath:    path.Join(reverseProxyPath, "templates", "nginx.conf"),
			Data:       tplData,
		},
	}, nil
}

type projectStepAccumulator struct {
	Project    generators.Project
	AppContext generators.ApplicationContext
	steps      []generators.Step
}

func (a *projectStepAccumulator) AppendSteps(steps ...generators.Step) {
	a.steps = append(a.steps, steps...)
}

func (a *projectStepAccumulator) AppendStepsFromGenerators(projectGenerators ...generators.Generator) error {
	for _, g := range projectGenerators {
		s, err := g.Steps()
		if err != nil {
			return fmt.Errorf("failed to get steps: %w", err)
		}
		a.steps = append(a.steps, s...)
	}

	return nil
}

func (a *projectStepAccumulator) Steps() []generators.Step {
	return a.steps
}
