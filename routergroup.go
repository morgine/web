package web

type Router interface {
	Use(handlers ...Handler) Router
	Handle(method, path string, handler Handler)
	HandleFunc(method, path string, handlerFunc HandlerFunc)
}

type Grouper interface {
	Group(name, comment, path string) Grouper
	Router
}

type Engine struct {
	groups  map[int]*Group
	routes  map[int]*Route
	filters map[int]*Filter
}

func (e Engine) Group(name, comment, path string) {
	panic("implement me")
}

func (e Engine) Use(handlers ...Handler) Router {
	panic("implement me")
}

func (e Engine) Handle(method, path string, handler Handler) {
	panic("implement me")
}

func (e Engine) HandleFunc(method, path string, handlerFunc HandlerFunc) {
	panic("implement me")
}

type Group struct {
	ID      int
	Name    string
	Comment string
	Path    string
	Subs    []*Group
	Routes  []*Route
}

type Route struct {
	ID      int
	Method  string
	Path    string
	Filters []Filter
	Handler Handler
}

type Filter struct {
	ID      int
	Handler Handler
}
