package main

import (
	"fmt"
	"net/http"
	"platform/config"
	"platform/logging"
	"platform/templates"
)

type HTMLHandler struct {
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
}

func main() {
	/* Load config */
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}
	/* --------- */

	logger := logging.NewDefaultLogger(cfg)

	templateExecutor, err := templates.NewTemplateExecutor(cfg, logger)
	if err != nil {
		logger.Panicf("cannot create template executor: %s\n", err.Error())
	}

	htmlHandler := &HTMLHandler{
		templateExecutor: templateExecutor,
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", htmlHandler)

	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Printf("error listening and serving: %s\n", err.Error())
	}
}
