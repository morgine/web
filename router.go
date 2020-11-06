package web

//
//import (
//	"net/http"
//	"path/filepath"
//	"strings"
//)
//
//type Morgine interface {
//	http.Handler
//	Group(group *Group) Router
//}
//
//type Engine interface {
//	Group(name, comment, path string) Router
//	Container() Container
//	Matcher() Matcher
//}
//
//type Router interface {
//	Use(handlers ...Handler) Router
//	Handle(method, path string, h Handler)
//	HandleFunc(method, path string, h HandlerFunc)
//}
//
//type Container interface {
//	Groups() []*Group
//	Routes(groupID int) []*Route
//	Filters() []Handler
//}
//
//type Matcher interface {
//	Add(route *Route)
//	Routes() []*Route
//	Match(method, path string) *Route
//}
//
//type Group struct {
//	ID      int
//	Name    string
//	Comment string
//	Path    string
//}
//
//type Route struct {
//	Method   string
//	Path     string
//	Handlers []Handler
//}
//
//func NewEngine() Engine {
//	return &engine{}
//}
//
//type engine struct {
//	groups []*group
//}
//
//func (e *engine) Container() Container {
//	c := &container{routes: map[int][]*Route{}}
//	existsFilters := map[Handler]struct{}{}
//	for _, g := range e.groups {
//		c.groups = append(c.groups, &Group{
//			ID:      g.ID,
//			Name:    g.Name,
//			Comment: g.Comment,
//			Path:    g.Path,
//		})
//		var routes []*Route
//		for _, r := range g.Routes {
//			routes = append(routes, &Route{
//				Method:   r.Method,
//				Path:     filepath.FromSlash(filepath.Join(g.Path, r.Path)),
//				Handlers: r.Handlers[:],
//			})
//			if ln := len(r.Handlers); ln > 1 {
//				for i := 0; i < ln-1; i++ {
//					filter := r.Handlers[i]
//					if _, ok := existsFilters[filter]; !ok {
//						existsFilters[filter] = struct{}{}
//						c.filters = append(c.filters, filter)
//					}
//				}
//			}
//		}
//		c.routes[g.ID] = routes
//	}
//	return c
//}
//
//func (e *engine) Matcher() Matcher {
//	m := NewMatcher()
//	for _, group := range e.groups {
//		for _, route := range group.Routes {
//			m.Add(route)
//		}
//	}
//	return m
//}
//
//func (e *engine) Group(name, comment, path string) Router {
//	g := &group{
//		Group: Group{
//			ID:      len(e.groups) + 1,
//			Name:    name,
//			Comment: comment,
//			Path:    path,
//		},
//		Routes: nil,
//	}
//	e.groups = append(e.groups, g)
//	return g
//}
//
//type group struct {
//	Group
//	Routes []*Route
//}
//
//func (r *group) Use(handlers ...Handler) Router {
//	return &conditionalHolder{
//		g:       r,
//		filters: handlers,
//	}
//}
//
//func (r *group) HandleFunc(method, path string, h HandlerFunc) {
//	r.Handle(method, path, h)
//}
//
//func (r *group) Handle(method, path string, h Handler) {
//	r.Routes = append(r.Routes, &Route{
//		Method:   method,
//		Path:     path,
//		Handlers: []Handler{h},
//	})
//}
//
//type container struct {
//	groups  []*Group
//	routes  map[int][]*Route
//	filters []Handler
//}
//
//func (c *container) Routes(groupID int) []*Route {
//	return c.routes[groupID]
//}
//
//func (c *container) Groups() []*Group {
//	return c.groups
//}
//
//func (c *container) Filters() []Handler {
//	return c.filters
//}
//
//type matcher struct {
//	methods   []*Route
//	anonymous []*Route
//	root      *Route
//	all []*Route
//}
//
//func NewMatcher() Matcher {
//	return &matcher{}
//}
//
//func (m *matcher) Routes() []*Route {
//	return m.all
//}
//
//func (m *matcher) Add(route *Route) {
//	if route.Method != "" {
//		m.methods = append(m.methods, route)
//	} else {
//		if route.Path != "/" {
//			m.anonymous = append(m.anonymous, route)
//		} else {
//			m.root = route
//		}
//	}
//	m.all = append(m.all, route)
//}
//
//func (m *matcher) Match(method, path string) *Route {
//	for _, route := range m.methods {
//		if route.Method == method && matchPath(route.Path, path) {
//			return route
//		}
//	}
//	for _, a := range m.anonymous {
//		if matchPath(a.Path, path) {
//			return a
//		}
//	}
//	return m.root
//}
//
//func matchPath(route, path string) bool {
//	if route == path {
//		return true
//	}
//	if !strings.HasSuffix(route, "/") {
//		return false
//	}
//
//	// 下列特殊情况将不被匹配:
//	//  route: "/directory/"
//	//  path: "/directory"
//	path = strings.TrimSuffix(path, "/")
//	return strings.HasPrefix(path, route)
//}
//
//type conditionalHolder struct {
//	g       *group
//	filters []Handler
//}
//
//func (r *conditionalHolder) Use(handlers ...Handler) Router {
//	rn := r.copy()
//	for _, h := range handlers {
//		rn.filters = append(rn.filters, h)
//	}
//	return rn
//}
//
//func (r *conditionalHolder) HandleFunc(method, path string, h HandlerFunc) {
//	r.Handle(method, path, h)
//}
//
//func (r *conditionalHolder) Handle(method, path string, h Handler) {
//	r.g.Routes = append(r.g.Routes, &Route{
//		Method:   method,
//		Path:     path,
//		Handlers: append(r.filters[:], h),
//	})
//}
//
//func (r *conditionalHolder) copy() *conditionalHolder {
//	return &conditionalHolder{
//		g:       r.g,
//		filters: r.filters[:],
//	}
//}
