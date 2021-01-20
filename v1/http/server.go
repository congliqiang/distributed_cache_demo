package http

import (
	"github.com/congliqiang/distributed_cache_demo/v1/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}
