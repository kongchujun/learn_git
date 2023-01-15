package middlewaretrain

import (
	"fmt"
	"testing"
)

func TestMiddleTrain(t *testing.T) {
	o := &OrginalContruct{
		ctx: &Context{
			Key:   "name",
			Value: "kong",
		},
	}
	o.mdls = []Middleware{
		func(next HandleFunc) HandleFunc {
			return func(ctx *Context) {
				fmt.Println("第一个")
				next(ctx)
			}
		},
		func(next HandleFunc) HandleFunc {
			return func(ctx *Context) {
				fmt.Println("第二个")
				next(ctx)
			}
		},
		func(next HandleFunc) HandleFunc {
			return func(ctx *Context) {
				fmt.Println("第三个")
				next(ctx)
			}
		},
	}
	o.ServeReady()
}
