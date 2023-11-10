package main

import (
	"platform/config"
	"platform/http"
	"platform/logging"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/templates"
)

/* type HTMLHandler struct {
	templateExecutor *templates.TemplateExecutor
}

func (handler *HTMLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path: %s\n", r.URL.Path)

	if r.Method == "GET" {
		if r.URL.Path == "/" {
			handler.templateExecutor.ExecTemplate(w, "home.html", nil)
		} else if r.URL.Path == "/about" {
			handler.templateExecutor.ExecTemplate(w, "about.html", nil)
		}
	}
} */

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	logger := logging.NewDefaultLogger(cfg)
	executor, err := templates.NewTemplateExecutor(cfg, logger)
	if err != nil {
		panic(err)
	}

	pipeline := createPipeline(cfg, logger, executor)

	http.Serve(pipeline, cfg, logger).Wait()
}

func createPipeline(config config.Configuration, logger logging.Logger, executor *templates.TemplateExecutor) pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.LoggingComponent{Logger: logger},
		&basic.ErrorComponent{Logger: logger},
		&basic.StaticFileComponent{Config: config},
		&basic.RouterComponent{Logger: logger},
		/* 		&placeholder.SimpleMessageComponent{Config: config, TemplateExecutor: executor}, */
	)
}
