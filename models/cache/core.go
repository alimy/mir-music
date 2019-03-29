package cache

import (
	"context"
	"fmt"
	"github.com/alimy/mir-music/models/model"
	"github.com/mediocregopher/radix/v3"
	"github.com/unisx/logus"
	"runtime"

	appCtx "github.com/alimy/mir-music/pkg/context"
)

const (
	defaultCacheEntrySize   = 10
	defaultCacheEntryExpire = 60

	actExpire = iota
	actJsonGet
)

var (
	rds radix.Client

	cache           chan *jsonEntry
	errNoExistEntry = fmt.Errorf("no exist entry")
)

// Cache indicate content cache operator
type Cache interface {
	// EntryTo write cached entry to gin.Context
	EntryTo(key string, c appCtx.Context) error

	// CacheFrom cache entry
	CacheFrom(key string, entry interface{})

	// Expire make cache expire
	Expire(key string)
}

type jsonEntry struct {
	act int
	key string
	val string
}

// Start start redis handler
func Start(c context.Context, conf *model.Config) {
	var err error
	rds, err = newRedisClient(conf)
	if err != nil {
		logus.F("start redis routine", err)
	}

	// start cache routine
	cache = make(chan *jsonEntry, defaultCacheEntrySize)
	go cacheEntry(c, cache)
}

// Done clean redis resource
func Done() {
	if rds != nil {
		rds.Close()
	}
}

func cacheEntry(ctx context.Context, cache <-chan *jsonEntry) {
	for {
		select {
		case entry := <-cache:
			handleCacheAct(entry)
		case <-ctx.Done():
			logus.Info("cache service cancel")
			return
		}
	}
}

func handleCacheAct(entry *jsonEntry) {
	switch entry.act {
	case actJsonGet:
		err := rds.Do(radix.Cmd(nil, "JSON.SET", entry.key, ".", entry.val))
		if err == nil {
			err = rds.Do(radix.FlatCmd(nil, "EXPIRE", entry.key, defaultCacheEntryExpire))
		}
		if err != nil {
			logus.Error("cache entry", logus.ErrorField(err))
		}
	case actExpire:
		if err := rds.Do(radix.FlatCmd(nil, "EXPIRE", entry.key, 0)); err != nil {
			logus.Error("cache entry", logus.ErrorField(err))
		}
	}
}

func newRedisClient(conf *model.Config) (radix.Client, error) {
	// parallel defines a multiplicand used for determining the number of goroutines
	// for running benchmarks. this value will be multiplied by GOMAXPROCS inside RunParallel.
	// since these benchmarks are mostly I/O bound and applications tend to have more
	// active goroutines accessing Redis than cores, especially with higher core numbers,
	// we set this to GOMAXPROCS so that we get GOMAXPROCS^2 connections.
	parallel := runtime.GOMAXPROCS(0)

	// multiply parallel with GOMAXPROCS to get the actual number of goroutines and thus
	// connections needed for the benchmarks.
	poolSize := parallel * runtime.GOMAXPROCS(0)

	return radix.NewPool(conf.Redis.Network, conf.Redis.Addr, poolSize, radix.PoolPipelineWindow(0, 0))
}
