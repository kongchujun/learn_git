package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errKeyNotFound = errors.New("cache: 找不到key")
)

type BuildInMapCache struct {
	data      map[string]*item
	mutex     sync.RWMutex
	close     chan struct{}
	onEvicted func(key string, val any)
}

type BuildInMapCacheOption func(cache *BuildInMapCache)

func BuildInMapCacheWithOnEvictedFunc(fn func(key string, val any)) BuildInMapCacheOption {
	return func(cache *BuildInMapCache) {
		cache.onEvicted = fn
	}
}

func NewBuildInMapCache(interval time.Duration, opts ...BuildInMapCacheOption) *BuildInMapCache {
	res := &BuildInMapCache{
		data:  make(map[string]*item, 100),
		close: make(chan struct{}),
	}
	for _, opt := range opts {
		opt(res)
	}
	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case t := <-ticker.C: //返回当前时间对象
				res.mutex.Lock()
				i := 0
				for key, val := range res.data {
					if i > 10000 {
						break
					}
					if !val.deadline.IsZero() && val.deadline.Before(t) {
						res.delete(key)
					}
				}
				res.mutex.Unlock()
			case <-res.close:
				return
			}
		}
	}()
	return res
}

func (b *BuildInMapCache) Get(ctx context.Context, key string) (any, error) {
	b.mutex.RLock()
	res, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return nil, fmt.Errorf("%w:%s", errKeyNotFound, key)
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		defer b.mutex.Unlock()
		res, ok = b.data[key]
		if !ok {
			return nil, fmt.Errorf("%w:%s", errKeyNotFound, key)
		}
		if res.deadlineBefore(now) {
			b.delete(key)
			return nil, fmt.Errorf("%w:%s", errKeyNotFound, key)
		}
		return res, nil
	}
	return res, nil
}

func (b *BuildInMapCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.data[key] = &item{
		val:      val,
		deadline: time.Now().Add(expiration),
	}
	return nil
}

func (b *BuildInMapCache) delete(key string) {
	itm, ok := b.data[key]
	if !ok {
		return
	}
	delete(b.data, key)
	b.onEvicted(key, itm.val)
}

func (b *BuildInMapCache) Delete(ctx context.Context, key string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.delete(key)
	return nil
}

func (b *BuildInMapCache) Close() error {
	select {
	case b.close <- struct{}{}:
	default:
		return errors.New("重复退出")
	}
	return nil
}

type item struct {
	val      any
	deadline time.Time
}

func (i *item) deadlineBefore(t time.Time) bool {
	return !i.deadline.IsZero() && i.deadline.Before(t)
}
