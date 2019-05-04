package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/mir-music/models/core"
	"github.com/unisx/logus"
	"net/http"
)

// Profile simple profile
type Profile struct {
	group      mir.Group `mir:"v1"`
	getAppInfo mir.Get   `mir:"/appinfo/"`

	*core.Context
}

// GetAppInfo GET handler of "/appinfo/"
func (p *Profile) GetAppInfo(c Context) {
	//profiles := p.Repo.GetProfiles()
	//c.JSON(http.StatusOK, profiles)
	c.String(http.StatusOK, `{"info": {
"profiles": "Mir music info",
"services": "provide music info search"
}}`)
	logus.Debug("get application information")
}
