package http1

type HandleFunc func(*Request) *Response

type Router struct {
	routes map[string]map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandleFunc),
	}
}

func (r *Router) Handle(method string, path string, f HandleFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandleFunc)
	}
	r.routes[method][path] = f
}

func (r *Router) FindHandler(method string, path string) (HandleFunc, bool) {
	if m, ok := r.routes[method]; ok {
		if h, ok := m[path]; ok {
			return h, true
		} else {
			return PathNotFound, true
		}
	} else {
		return PathMethodNotAllowed, true
	}
}
