package cache

import (
	"github.com/alimy/mir-music/pkg/context"
	"github.com/alimy/mir-music/pkg/json"
	"github.com/mediocregopher/radix/v3"
	"github.com/unisx/logus"

	"net/http"
)

// RedisCache Cache interface implement by redis
type RedisCache struct {
	// TODO
}

// EntryTo write cached entry to gin.Context
func (r *RedisCache) EntryTo(key string, c context.Context) error {
	content, err := r.entryFrom(key)
	if err == nil {
		c.Writer.WriteHeader(http.StatusOK)
		header := c.Writer.Header()
		header.Set("Content-Type", "application/json; charset=utf-8")
		_, err = c.Writer.Write(content)
	}
	return err
}

// CacheFrom cache entry
func (r *RedisCache) CacheFrom(key string, entry interface{}) {
	jsonVal, err := json.Marshal(entry)
	if err != nil {
		logus.E("marshal entry", err)
		return
	}
	cache <- &jsonEntry{act: actJsonGet, key: key, val: string(jsonVal)}
}

func (r *RedisCache) Expire(key string) {
	cache <- &jsonEntry{act: actExpire, key: key}
}

func (r *RedisCache) entryFrom(key string) ([]byte, error) {
	var entry []byte
	err := rds.Do(radix.Cmd(&entry, "JSON.GET", key, "."))
	if err == nil && len(entry) == 0 {
		err = errNoExistEntry
	}
	return entry, err
}
