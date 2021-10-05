package cmd

import (
	"testing"
	"context"
	"strings"

	"github.com/stretchr/testify/require"
	"github.com/tharsis/token/app"
)

func Test_runDeployCmd(t *testing.T) {
	var (
		err error
		client *app.Client
	)


	tt := []struct{
		name string
		pretest func()
		expErr bool
	}{
		{
		"success",
		func() {
			client, err = app.NewClient(app.Mnemonic)
			// how to check for an error 
		},
		false,
	    },
		{
		"invalid mnemonic",
		func() {
			invalidMnemonic := strings.Replace(app.Mnemonic, "sight", "seen", -1)
			client, err = app.NewClient(invalidMnemonic)
			// how to check for an error 
		},
		true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			tc.pretest()
			cmd := deployCmd
			
			clientCtx := app.Context{}.WithClient(client)
			ctx := context.WithValue(context.Background(), app.ClientContextKey, &clientCtx)
			if tc.expErr {
				require.NoError(t, err)
				require.NoError(t, cmd.ExecuteContext(ctx))
			} else {
				require.Error(t, err)
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}

}