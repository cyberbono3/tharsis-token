package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tharsis/token/app"
)

// TODO fix invalid mnemonic test case
// Error: DeployToken err: " --- at github.com/tharsis/ethermint/app/ante/eth.go:217 (EthNonceVerificationDecorator.AnteHandle) ---\nCaused by: invalid nonce; got 11, expected 10: invalid sequence"
func Test_runDeployCmd(t *testing.T) {
	var (
		err    error
		client *app.Client
	)

	tt := []struct {
		name    string
		pretest func()
		expErr  bool
	}{
		{
			"success",
			func() {
				client, err = app.NewClient(app.Mnemonic)
				// how to check for an error
			},
			false,
		},
		/*
			{
			"invalid mnemonic",
			func() {
				invalidMnemonic := strings.Replace(app.Mnemonic, "sight", "seen", -1)
				client, err = app.NewClient(invalidMnemonic)
				// how to check for an error
			},
			true,
			},
		*/

	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cmd := deployContractCommand()
			tc.pretest()

			clientCtx := app.Context{}.WithClient(client)
			ctx := context.WithValue(context.Background(), app.ClientContextKey, &clientCtx)
			if tc.expErr {
				require.Error(t, cmd.ExecuteContext(ctx))
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}
