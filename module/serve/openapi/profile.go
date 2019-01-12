package openapi

import (
	"github.com/alimy/mir"
	"github.com/unisx/logus"
	"net/http"
)

type profile struct {
	group      mir.Group `mir:"v1"`
	getAppInfo mir.Get   `mir:"/appinfo/"`
}

// GetAppInfo GET handler of "/appinfo/"
func (p *profile) GetAppInfo(c Context) {
	// TODO
	logus.Debug("get application information")
	c.String(http.StatusOK, "get application information")
}
