package web

import "testing"

func TestServera(t *testing.T) {
	h := NewHTTPServer()
	h.Get("/", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello world"))
	})
	h.Get("/uer", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello user"))
	})
	h.Start(":8081")
}
