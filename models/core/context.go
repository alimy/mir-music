package core

import (
	"github.com/alimy/mir-music/models/cache"
	"github.com/alimy/mir-music/pkg/context"
	"net/http"
)

// Context core context for gin handler
type Context struct {
	cache.Cache

	Repo Repository
}

// Retrieve write response content to c
func (ctx *Context) Retrieve(c context.Context, cacheKey string, action func() (interface{}, error)) {
	if err := ctx.Cache.EntryTo(cacheKey, c); err != nil {
		data, err := action()
		if err == nil {
			c.JSON(http.StatusOK, data)
			ctx.CacheFrom(cacheKey, data)
		}
		if err != nil {
			ctx.ErrInternalServer(c, err.Error())
		}
	}
}

func (ctx *Context) ErrInternalServer(c context.Context, msg ...string) {
	if len(msg) > 1 {
		c.String(http.StatusInternalServerError, msg[0])
	} else {
		c.String(http.StatusInternalServerError, "internal server error")
	}
}

func (ctx *Context) ErrNotFound(c context.Context, msg ...string) {
	if len(msg) > 1 {
		c.String(http.StatusNotFound, msg[0])
	} else {
		c.String(http.StatusNotFound, "not found")
	}
}
