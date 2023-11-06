package placeholder

import (
	"errors"
	"io"
	"platform/config"
	"platform/pipeline"
)

type SimpleMessageComponent struct {
	Config config.Configuration
}

func (c *SimpleMessageComponent) Init() {}
func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {
	msg, ok := c.Config.GetString("main:message")
	if ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("cannot find config setting"))
	}
	next(ctx)
}
