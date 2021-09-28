package app

import (
	"github.com/spf13/cobra"
)

// ClientContextKey defines the context key used to retrieve a client.Context from
// a command's Context.
const ClientContextKey = ContextKey("client.context")

// ContextKey defines a type alias for a stdlib Context key.
type ContextKey string

// GetClientContextFromCmd returns a Context from a command or an empty Context
// if it has not been set.
func getClientContextFromCmd(cmd *cobra.Command) Context {
	if v := cmd.Context().Value(ClientContextKey); v != nil {
		clientCtxPtr := v.(*Context)
		return *clientCtxPtr
	}

	return Context{}
}

func GetClientContext(cmd *cobra.Command) (Context, error) {
	ctx := getClientContextFromCmd(cmd)
	if ctx.Client == nil {
		client, err := NewClient(mnemonic)
		if err != nil {
			return ctx, err
		}

		ctx = ctx.WithClient(client)
	}

	return ctx, nil
}
