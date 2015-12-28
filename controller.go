package main

import "github.com/netopssh/megacli-agent/cache"

// Controller
type Controller struct {
	Cache *cache.Cache
}

// New
func New(c *cache.Cache) *Controller {
	ctrl := &Controller{}
	ctrl.Cache = c

	return ctrl
}
