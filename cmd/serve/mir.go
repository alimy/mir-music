package serve

import (
	"github.com/alimy/mir-music/models"
	"github.com/alimy/mir-music/pkg/openapi"
)

// mirEntries get all entries that used to register to Mir
// Notice: this func must call after models.InitWith(...)
func mirEntries() []interface{} {
	ctx := models.NewContext()

	entries := []interface{}{
		&openapi.Profile{Context: ctx},
		&openapi.Media{Context: ctx},
	}
	if portal := openapi.MirPortal(); portal != nil {
		entries = append(entries, portal)
	}
	return entries
}
