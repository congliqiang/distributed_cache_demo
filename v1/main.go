package main

import (
	"github.com/congliqiang/distributed_cache_demo/v1/cache"
	"github.com/congliqiang/distributed_cache_demo/v1/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
