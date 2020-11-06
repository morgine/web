package io

import (
	"net/http"
)

type Router interface {
	Use(handlers ...Handler) Router
	Handle(method, path string, handler Handler)
	HandleFunc(method, path string, handlerFunc HandlerFunc)
}

type Engine struct {
	groups  []*Group
	filters []Handler
	routes  []*Route
	matcher Matcher
}

type Matcher interface {
	Add(route *Route)
	Match(method, path string) *Route
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (e *Engine) Group(name, comment, path string) Router {
	group := &Group{
		ID:      len(e.groups) + 1,
		Name:    name,
		Path:    path,
		Comment: comment,
	}
	e.groups = append(e.groups, group)
	return &groupEngine{group: group, engine: e}
}

func (e *Engine) Use(handlers ...Handler) Router {
	filterEngine{}
	e.filters = append(e.filters, handlers...)
	return e
}

func (e *Engine) Handle(method, path string, handler Handler) {
	panic("implement me")
}

func (e *Engine) HandleFunc(method, path string, handlerFunc HandlerFunc) {
	panic("implement me")
}

type groupEngine struct {
	group  *Group
	engine *Engine
}

func (g *groupEngine) Use(handlers ...Handler) Router {
	panic("implement me")
}

func (g *groupEngine) Handle(method, path string, handler Handler) {
	panic("implement me")
}

func (g *groupEngine) HandleFunc(method, path string, handlerFunc HandlerFunc) {
	panic("implement me")
}

type filterEngine struct {
	group   *Group
	filters []Handler
}

func (f *filterEngine) Use(handlers ...Handler) Router {

}

func (f *filterEngine) Handle(method, path string, h Handler) {
	panic("implement me")
}

func (f *filterEngine) HandleFunc(method, path string, h HandlerFunc) {
	panic("implement me")
}

type Group struct {
	ID      int
	Name    string
	Comment string
	Path    string
}

type Route struct {
	ID       int
	Method   string
	Path     string
	Handlers []Handler
}

type Handler interface {
	Handle(ctx *Context)
}

type HandlerFunc func(ctx *Context)

func (h HandlerFunc) Handle(ctx *Context) {
	h(ctx)
}

type Context struct {
}

// 去重
func distinct(handlers []Handler) []Handler {
	exists := make(map[Handler]struct{}, len(handlers))
	var hs []Handler
	for _, handler := range handlers {
		if _, ok := exists[handler]; !ok {
			hs = append(hs, handler)
			exists[handler] = struct{}{}
		}
	}
	return hs
}
