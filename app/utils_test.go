package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitEthClient(t *testing.T) {
	testCases := []struct {
		name        string
		endpointStr string
		expErr      bool
	}{
		{"success", rpcEndpoint, false},
		{"invalid endpoint", "127.0.0.1:4546", true},
		{"empty endpoint", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := initEthClient(tc.endpointStr)
			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, res)
			} else {
				require.NoError(t, err)
				require.NotNil(t, res)
			}
		})
	}
}

func TestAddressFromMnemonic(t *testing.T){
	mnemonic := "sound practice disease erupt basket pumpkin truck file gorilla behave find exchange napkin boy congress address city net prosper crop chair marine chase seven"

	privKey, err := privKeyFromMnemonic(mnemonic)
	require.NoError(t, err)

	addrHex, err := addressFromPrivKey(privKey)
	require.NoError(t, err)

	require.Equal(t, "0x3943412CBEEEd4b68d73382b136F36b0CB82F481", addrHex)
}