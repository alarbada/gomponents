package actions

import (
	"fmt"
	"strings"

	g "github.com/alarbada/gomponents"
	"github.com/gin-gonic/gin"
)

type Router struct {
	groups []string
	engine gin.IRouter
}

func NewRouter() Router {
	router := gin.Default()
	return Router{nil, router}
}

func (r *Router) Group(path string) Router {
	group := r.engine.Group(path)

	pathGroups := make([]string, len(r.groups))
	copy(pathGroups, r.groups)
	pathGroups = append(pathGroups, path)

	return Router{pathGroups, group}
}

func (r *Router) Engine() *gin.Engine {
	if e, ok :=  r.engine.(*gin.Engine); ok {
		return e
	} else {
		panic("cannot get static engine from non-engine router")
	}
}

type Action struct {
	Method, Path string
	router       *Router
}

func (r *Router) action(method, path string) *Action {
	return &Action{
		Method: method,
		Path:   path,
		router: r,
	}
}

func (r *Router) POST(path string) *Action   { return r.action("POST", path) }
func (r *Router) GET(path string) *Action    { return r.action("GET", path) }
func (r *Router) PUT(path string) *Action    { return r.action("PUT", path) }
func (r *Router) DELETE(path string) *Action { return r.action("DELETE", path) }
func (r *Router) PATCH(path string) *Action  { return r.action("PATCH", path) }

func (a *Action) Handle(action func(c *gin.Context) g.Node) *Action {
	wrapped := func(c *gin.Context) {
		if result := action(c); result != nil {
			result.Render(c.Writer)
		}
	}

	r := a.router.engine
	method := a.Method
	url := a.Path

	switch method {
	case "GET":
		r.GET(url, wrapped)
	case "POST":
		r.POST(url, wrapped)
	case "PUT":
		r.PUT(url, wrapped)
	case "DELETE":
		r.DELETE(url, wrapped)
	case "PATCH":
		r.PATCH(url, wrapped)
	default:
		panic(fmt.Sprintf("invalid method %s", method))
	}

	return a
}

func (a *Action) Hx() g.Node {
	groupPath := strings.Join(a.router.groups, "/")
	path := groupPath + a.Path

	return g.Attr("hx-"+strings.ToLower(a.Method), path)
}
