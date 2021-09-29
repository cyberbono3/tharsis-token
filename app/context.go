package app

type Context struct {
	Client *Client
}

// WithClient puts client on Context.Client field
func (ctx Context) WithClient(client *Client) Context {
	ctx.Client = client
	return ctx
}

