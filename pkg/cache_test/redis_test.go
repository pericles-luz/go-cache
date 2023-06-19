package cache_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/cache"
	"github.com/stretchr/testify/require"
)

func TestRedisConnection(t *testing.T) {
	t.Skip("Test only if necessary")
	cache := cache.NewRedis("redis")
	err := cache.Ping()
	require.NoError(t, err)
}

func TestRedisSetGetDel(t *testing.T) {
	t.Skip("Test only if necessary")
	cache := cache.NewRedis("redis")
	err := cache.Set("key", "value", 0)
	require.NoError(t, err)
	val, err := cache.Get("key")
	require.NoError(t, err)
	require.Equal(t, "value", val)
	err = cache.Del("key")
	require.NoError(t, err)
}
