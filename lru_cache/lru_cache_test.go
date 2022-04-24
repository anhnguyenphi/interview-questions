package lru_cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCacheImpl_SetAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	val, err := cache.Get("a")
	assert.NotNil(t, err)

	cache.Set("a", []byte("1"))
	val, err = cache.Get("a")
	assert.Nil(t, err)
	assert.Equal(t, val, []byte("1"))

	cache.Set("b", []byte("2"))
	val, err = cache.Get("b")
	assert.Nil(t, err)
	assert.Equal(t, val, []byte("2"))

	cache.Set("a", []byte("3"))
	val, err = cache.Get("a")
	assert.Nil(t, err)
	assert.Equal(t, val, []byte("3"))

	cache.Set("c", []byte("4"))
	val, err = cache.Get("c")
	assert.Nil(t, err)
	assert.Equal(t, val, []byte("4"))

	val, err = cache.Get("b")
	assert.NotNil(t, err)
}
