package serverhandler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Router struct
type Router struct {
	*httprouter.Router
	URLAllowMethods map[string][]string
}

// Get routes Get request
func (r *Router) Get(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "GET")
	r.GET(path, DefaultHandler(handler))
}

// Delete routes Delete request
func (r *Router) Delete(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "DELETE")
	r.DELETE(path, DefaultHandler(handler))
}

// Patch routes Patch request
func (r *Router) Patch(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "PATCH")
	r.PATCH(path, DefaultHandler(handler))
}

// Post routes Post request
func (r *Router) Post(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "POST")
	r.POST(path, DefaultHandler(handler))
}

// Put routes Put request
func (r *Router) Put(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "PUT")
	r.PUT(path, DefaultHandler(handler))
}

// Options routes Options request
func (r *Router) Options(path string, handler func(s *RequestState)) {
	r.URLAllowMethods[path] = append(r.URLAllowMethods[path], "OPTIONS")
	r.OPTIONS(path, func(w http.ResponseWriter, r *http.Request,
		ps httprouter.Params) {

		s := &RequestState{
			W:  w,
			R:  r,
			Ps: ps,
		}
		handler(s)
	})
}

//NewRouter builder for router struct
func NewRouter() *Router {
	return &Router{httprouter.New(), make(map[string][]string)}
}
