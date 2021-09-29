package cmd

import (
	"testing"
	"context"

	"github.com/tharsis/token/app"
	"github.com/stretchr/testify/require"
)

func Test_runDeployCmd(t *testing.T) {
	cmd := deployCmd

	client, err := app.NewClient(app.Mnemonic)
	require.NoError(t, err)
	clientCtx := app.Context{}.WithClient(client)
	ctx := context.WithValue(context.Background(), app.ClientContextKey, &clientCtx)
	cmd.SetArgs([]string{
		"keyname1", keyfile,
	err := cmd.ExecuteContext(ctx)


}