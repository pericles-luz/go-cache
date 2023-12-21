package cache_test

import (
	"testing"

	"github.com/pericles-luz/go-cache/pkg/cache"
	"github.com/stretchr/testify/require"
)

func TestMemoryCache(t *testing.T) {
	cache := cache.NewMemory()
	err := cache.Ping()
	require.NoError(t, err)
	payload := "value"
	err = cache.Set("key", payload, 1)
	require.NoError(t, err)
	val, err := cache.Get("key")
	require.NoError(t, err)
	require.Equal(t, payload, val)
}
