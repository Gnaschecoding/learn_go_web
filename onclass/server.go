package main

import "net/http"

type Server interface {
	//Route(pattern string, handleFunc func(ctx *Context))
	Routable
	Start(address string) error
}

// sdkHttpServer 基于 http实现
type sdkHttpServer struct {
	Name    string
	Handler Handler
	Root    Filter
}

// Route 注册路由
func (s *sdkHttpServer) Route(mathod string, pattern string, handleFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
	s.Handler.Route(mathod, pattern, handleFunc)

}

// Start
func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		s.Root(c)
	})
	//http.Handle("/", s.Handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlersBaseOnMap()
	var root Filter = func(c *Context) {
		handler.ServeHTTP(c)
	}
	for i := len(builders); i > 0; i-- {
		b := builders[i-1]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		Handler: handler,
		Root:    root,
	}
}
