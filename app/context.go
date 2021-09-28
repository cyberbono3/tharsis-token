package app

type Context struct {
	Client *Client
}

func (ctx Context) WithClient(client *Client) Context {
	ctx.Client = client
	return ctx
}

