package cmd

import (
	"github.com/alimy/mir-music/module/serve/info"
	"github.com/gin-gonic/gin"

	mirE "github.com/alimy/mir/module/gin"
)

// newGin return a initialed gin.Engine instance
func newGin() *gin.Engine {
	// Instance a default gin engine
	e := gin.Default()

	// Register API
	entries := info.MirEntries()
	mirE.Register(e, entries...)

	return e
}
