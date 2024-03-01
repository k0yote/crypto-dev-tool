package wallet

import (
	"testing"

	"github.com/k0yote/backend-wallet/config"
	"github.com/stretchr/testify/require"
)

func TestEntropy(t *testing.T) {

	config, err := config.LoadConfig()
	require.NoError(t, err)

	b, err := entropy(config, 32)
	require.NoError(t, err)
	require.NotEmpty(t, b)
}
