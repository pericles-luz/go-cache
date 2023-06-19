package cache_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-cache/pkg/cache"
	"github.com/stretchr/testify/require"
)

func TestConfigMustLoad(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	config := cache.NewConfig()
	err := config.Load("redis")
	require.NoError(t, err)
	require.NotZero(t, len(config.GetConfig()["DE_Host"].(string)))
	require.NotZero(t, config.GetConfig()["NU_Port"].(int))
	require.NotZero(t, len(config.GetConfig()["PW_Senha"].(string)))
	require.Zero(t, config.GetConfig()["NU_Banco"].(int))
}
