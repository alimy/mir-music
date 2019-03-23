// +build portal

package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/music-ui/dist"
	"net/http"
)

type portalAssets struct {
	mainHandler      http.Handler
	staticHandler    http.Handler
	getMainAssets    mir.Get  `mir:"/"`
	headMainAssets   mir.Head `mir:"/"`
	getStaticAssets  mir.Get  `mir:"/static/*filepath"`
	headStaticAssets mir.Head `mir:"/static/*filepath"`
}

// GetMainAssets GET handler of "/"
func (p *portalAssets) GetMainAssets(c Context) {
	p.mainHandler.ServeHTTP(c.Writer, c.Request)
}

// HeadMainAssets HEAD handler of "/"
func (p *portalAssets) HeadMainAssets(c Context) {
	p.mainHandler.ServeHTTP(c.Writer, c.Request)
}

// GetStaticAssets GET handler of "/static/*filepath"
func (p *portalAssets) GetStaticAssets(c Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

// HeadStaticAssets HEAD handler of "/static/*filepath"
func (p *portalAssets) HeadStaticAssets(c Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

// mirPortal return a portal mir entry
func mirPortal() interface{} {
	assetFile := dist.AssetFile()
	return &portalAssets{
		mainHandler:   http.StripPrefix("/", http.FileServer(assetFile)),
		staticHandler: http.StripPrefix("/static", http.FileServer(assetFile)),
	}
}
