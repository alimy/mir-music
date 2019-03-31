// +build portal

package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/music-ui/dist"
	"net/http"
)

type PortalAssets struct {
	index            mir.Get  `mir:"/"`
	getMainAssets    mir.Get  `mir:"/index.html#Index"`
	getStaticAssets  mir.Get  `mir:"/static/*filepath"`
	headStaticAssets mir.Head `mir:"/static/*filepath"`

	staticHandler http.Handler
}

// GetMainAssets GET handler of "/"
func (p *PortalAssets) Index(c Context) {
	c.Status(http.StatusOK)
	c.Writer.Write(dist.MustAsset("index.html"))
}

// GetStaticAssets GET handler of "/static/*filepath"
func (p *PortalAssets) GetStaticAssets(c Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

// HeadStaticAssets HEAD handler of "/static/*filepath"
func (p *PortalAssets) HeadStaticAssets(c Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

// mirPortal return a portal mir entry
func MirPortal() interface{} {
	assetFile := dist.AssetFile()
	return &PortalAssets{
		staticHandler: http.StripPrefix("/static", http.FileServer(assetFile)),
	}
}
