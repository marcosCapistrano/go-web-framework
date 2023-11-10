package placeholder

import (
	"platform/config"
	"platform/pipeline"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Config           config.Configuration
	TemplateExecutor *templates.TemplateExecutor
}

func (c *SimpleMessageComponent) Init() {}
func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	err := c.TemplateExecutor.ExecTemplate(ctx.ResponseWriter, "index.html", nil)
	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
