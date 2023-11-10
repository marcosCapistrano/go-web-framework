package basic

import (
	"platform/config"
	"platform/logging"
	"platform/pipeline"
	"platform/templates"
	"strings"
)

type RouterComponent struct {
	config.Configuration
	logging.Logger
	templates.TemplateExecutor
}

func (c *RouterComponent) Init() {}
func (c *RouterComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	method := ctx.Request.Method

	if strings.HasPrefix(ctx.URL.Path, "/api") {
		//Todo handle
	} else {
		if method == "GET" {
			c.TemplateExecutor.ExecTemplate("")
		}
	}

	next(ctx)
}
