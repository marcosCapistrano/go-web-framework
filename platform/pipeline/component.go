package pipeline

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mwc *ComponentContext) Error(err error) {
	mwc.error = err
}

func (mwc *ComponentContext) GetError() error {
	return mwc.error
}

/*
Init performs any one-off setup needed.
ProcessRequest processes http requests
*/
type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}
