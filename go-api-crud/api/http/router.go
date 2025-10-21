package http

import "net/http"

type Method string

type Handler func(ctx *Context)

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

type Route struct {
	endpoint string
	method   Method
	handler  Handler
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

func (r *Router) Register() {
	routeMap := make(map[string][]Route)

	for _, route := range r.routes {
		routeMap[route.endpoint] = append(routeMap[route.endpoint], route)
	}

	for endpoint, routes := range routeMap {
		routesCopy := routes

		http.HandleFunc(endpoint, func(w http.ResponseWriter, req *http.Request) {
			ctx := &Context{w, req}

			for _, route := range routesCopy {
				if req.Method == string(route.method) {
					route.handler(ctx)
					return
				}
			}
			//TODO: change to handle error
			ctx.ResponseJson(http.StatusMethodNotAllowed, &map[string]interface{}{"message": "Método não permitido"})
		})
	}
}

func (r *Router) add(endpoint string, method Method, handler Handler) *Router {
	r.routes = append(r.routes, Route{
		endpoint: endpoint,
		method:   method,
		handler:  handler,
	})

	return r
}

func (r *Router) Get(endpoint string, handler Handler) *Router {
	return r.add(endpoint, GET, handler)
}

func (r *Router) Post(endpoint string, handler Handler) *Router {
	return r.add(endpoint, POST, handler)
}

func (r *Router) Put(endpoint string, handler Handler) *Router {
	return r.add(endpoint, PUT, handler)
}

func (r *Router) Delete(endpoint string, handler Handler) *Router {
	return r.add(endpoint, DELETE, handler)
}

func (r *Router) Patch(endpoint string, handler Handler) *Router {
	return r.add(endpoint, PATCH, handler)
}
