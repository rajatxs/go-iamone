package rpc

import "net/http"

type Common struct{}

func (c *Common) Ping(r *http.Request, args *any, reply *string) error {
	*reply = "Pong!"
	return nil
}
