package web

import (
	"net/http"
	"strings"
)

type Router interface {
	Use(handlers ...Handler) Router
	Version(number string) Router
	Handle(method, path string, handler Handler)
	HandleFunc(method, path string, handlerFunc HandlerFunc)
}

type Grouper interface {
	Group(name, comment, path string) Router
}

type group struct {
	Name    string
	Comment string
	Path    string
	Routes  []*route
}

type route struct {
	Method     string
	Path       string
	Version    string
	Deprecated bool
	Handlers   []Handler
}

type Group struct {
	ID       int
	Name     string
	Comment  string
	Path     string
	Handlers []*HandlerInfo
	Routes   []*Route
}

func (gi *Group) getIDs(handlers []Handler) []int {
	var ids []int
	for _, newH := range handlers {
		var id int
		func() {
			for _, oldH := range gi.Handlers {
				if newH == oldH.Handler {
					id = oldH.ID
					return
				}
			}
			id = len(gi.Handlers) + 1
			gi.Handlers = append(gi.Handlers, &HandlerInfo{
				ID:      id,
				Handler: newH,
			})
		}()
		ids = append(ids, id)
	}
	return ids
}

type HandlerInfo struct {
	ID      int
	Handler Handler
}

type Route struct {
	ID         int
	Method     string
	Path       string
	Version    string
	Deprecated bool
	Handlers   []int
	Handler    Handler
}

type Info struct {
	Version string
	Groups  []*Group
}

type Engine struct {
	groups []*group
	root   *routerEngine
}

func (e *Engine) HTTPHandler() http.Handler {

}

type matcher struct {
	methods   []*route
	anonymous []*route
	root      *route
	all       []*route
}

func (m *matcher) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r := m.match(request.Method, request.URL.Path)
	if r != nil {
		ctx := &Context{}
	} else {
		http.NotFound(writer, request)
	}
}

func (m *matcher) add(route *route) {
	if route.Method != "" {
		m.methods = append(m.methods, route)
	} else {
		if route.Path != "/" {
			m.anonymous = append(m.anonymous, route)
		} else {
			m.root = route
		}
	}
	m.all = append(m.all, route)
}

func (m *matcher) match(method, path string) *route {
	for _, route := range m.methods {
		if route.Method == method && matchPath(route.Path, path) {
			return route
		}
	}
	for _, a := range m.anonymous {
		if matchPath(a.Path, path) {
			return a
		}
	}
	return m.root
}

func matchPath(route, path string) bool {
	if route == path {
		return true
	}
	if !strings.HasSuffix(route, "/") {
		return false
	}

	// 下列特殊情况将不被匹配:
	//  route: "/directory/"
	//  path: "/directory"
	path = strings.TrimSuffix(path, "/")
	return strings.HasPrefix(path, route)
}

func (e *Engine) Info() *Info {
	var infos = &Info{}
	groups := append(e.groups, e.root.gp)
	for gid, g := range groups {
		gi := &Group{
			ID:       gid,
			Name:     g.Name,
			Comment:  g.Comment,
			Path:     g.Path,
			Handlers: nil,
			Routes:   nil,
		}
		for rid, r := range g.Routes {
			handlers, handler := func() ([]Handler, Handler) {
				if l := len(r.Handlers); l > 1 {
					return r.Handlers[:l-1], r.Handlers[l-1]
				}
				return nil, r.Handlers[0]
			}()
			gi.Routes = append(gi.Routes, &Route{
				ID:       rid,
				Method:   r.Method,
				Path:     r.Path,
				Handlers: gi.getIDs(handlers),
				Handler:  handler,
			})
		}
		infos.Groups = append(infos.Groups, gi)
	}
	return infos
}

func (e *Engine) Version(number string) Router {
	return e.root.Version(number)
}

func (e *Engine) Group(name, comment, path string) Router {
	gp := &group{
		Name:    name,
		Comment: comment,
		Path:    path,
	}
	e.groups = append(e.groups, gp)
	return &routerEngine{gp: gp, handlers: append([]Handler{}, e.root.handlers...), version: e.root.version}
}

func NewEngine(version string) *Engine {
	gp := &group{
		Name:    "DEFAULT",
		Comment: "default apis",
	}
	return &Engine{
		root: &routerEngine{
			gp:       gp,
			handlers: nil,
			version:  version,
		},
	}
}

// 设置之后所有注册的路由都将加上这些处理器
func (e *Engine) Use(handlers ...Handler) Router {
	return e.root.Use(handlers...)
}

func (e *Engine) Handle(method, path string, handler Handler) {
	e.root.Handle(method, path, handler)
}

func (e *Engine) HandleFunc(method, path string, handlerFunc HandlerFunc) {
	e.root.HandleFunc(method, path, handlerFunc)
}

type routerEngine struct {
	gp       *group
	handlers []Handler
	version  string
}

func (g *routerEngine) copy() *routerEngine {
	return &routerEngine{
		gp:       g.gp,
		version:  g.version,
		handlers: append([]Handler{}, g.handlers...),
	}
}

func (g *routerEngine) Version(number string) Router {
	newRE := g.copy()
	newRE.version = number
	return newRE
}

func (g *routerEngine) Use(handlers ...Handler) Router {
	newGE := g.copy()
	newGE.handlers = append(newGE.handlers, handlers...)
	return newGE
}

func (g *routerEngine) Handle(method, path string, handler Handler) {
	g.gp.Routes = append(g.gp.Routes, &route{
		Method:   method,
		Path:     path,
		Version:  g.version,
		Handlers: append(g.handlers, handler),
	})
}

func (g *routerEngine) HandleFunc(method, path string, handlerFunc HandlerFunc) {
	g.Handle(method, path, handlerFunc)
}
