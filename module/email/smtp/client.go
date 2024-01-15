package smtp

import "context"

type NewEmail struct {
}
type Client struct {
}

func (c *Client) Send(context.Context, NewEmail) error {
	return nil
}
