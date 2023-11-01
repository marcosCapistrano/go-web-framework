package templates

import (
	"errors"
	"html/template"
	"io"
	"platform/config"
	"platform/logging"
)

type TemplateExecutor struct {
	templates *template.Template
}

func NewTemplateExecutor(config config.Configuration, logger logging.Logger) (*TemplateExecutor, error) {
	path, ok := config.GetString("templates:path")
	if !ok {
		return nil, errors.New("cannot load template config")
	}

	templates, err := template.ParseGlob(path)
	if err != nil {
		return nil, err
	}

	executor := &TemplateExecutor{
		templates: templates,
	}

	return executor, nil
}

func (executor *TemplateExecutor) ExecTemplate(writer io.Writer, name string, data interface{}) (err error) {
	err = executor.templates.ExecuteTemplate(writer, name, data)

	return
}
