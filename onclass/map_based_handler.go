package main

import (
	"net/http"
)

type Routable interface {
	Route(mathod string, pattern string, handleFunc func(ctx *Context))
}

type Handler interface {
	//http.Handler
	ServeHTTP(c *Context)
	Routable
}

type HandlerBaseOnMap struct {
	handlers map[string]func(ctx *Context)
}

// Route 注册路由
func (h *HandlerBaseOnMap) Route(mathod string, pattern string, handleFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
	key := h.Key(mathod, pattern)
	h.handlers[key] = handleFunc

}

func (h *HandlerBaseOnMap) ServeHTTP(c *Context) {
	key := h.Key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}

}

func (h *HandlerBaseOnMap) Key(mathod string, pattern string) string {
	return mathod + "#" + pattern
}

var _ Handler = &HandlerBaseOnMap{}

func NewHandlersBaseOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
