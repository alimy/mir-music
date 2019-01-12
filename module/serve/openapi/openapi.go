package openapi

import "github.com/gin-gonic/gin"

// Context alias type of *gin.Context
type Context = *gin.Context

// MirEntries get all entries that used to register to Mir
func MirEntries() []interface{} {
	entries := []interface{}{
		&profile{},
		&media{},
	}
	if portal := mirPortal(); portal != nil {
		entries = append(entries, portal)
	}
	return entries
}
