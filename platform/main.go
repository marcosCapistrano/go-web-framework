package main

import (
	"platform/config"
	"platform/http"
	"platform/logging"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/placeholder"
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
	pipeline := createPipeline(cfg, logger)

	http.Serve(pipeline, cfg, logger).Wait()
}

func createPipeline(config config.Configuration, logger logging.Logger) pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.LoggingComponent{Logger: logger},
		&basic.ErrorComponent{Logger: logger},
		&basic.StaticFileComponent{Config: config},
		&placeholder.SimpleMessageComponent{Config: config},
	)
}
