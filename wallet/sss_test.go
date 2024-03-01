package wallet

import (
	"bytes"
	"testing"

	"github.com/k0yote/backend-wallet/util"
	"github.com/stretchr/testify/require"
)

func TestShamirShareSecret(t *testing.T) {

	privateKey, err := util.RandomPrivateKey()
	require.NoError(t, err)

	shared, err := GenerateSharedSecretBytes(privateKey)
	require.NoError(t, err)
	require.NotEmpty(t, shared)
	require.Equal(t, 3, len(shared))
}

func TestShamirRecoverSecret(t *testing.T) {
	privateKey, err := util.RandomPrivateKey()
	require.NoError(t, err)

	shared, err := GenerateSharedSecretBytes(privateKey)
	require.NoError(t, err)

	recovery, err := CombineThresholdShares(shared)
	require.NoError(t, err)

	require.True(t, bytes.Equal(privateKey, recovery))
}
