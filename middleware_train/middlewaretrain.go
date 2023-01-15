package middlewaretrain

import "fmt"

type Context struct {
	Key   string
	Value string
}

type HandleFunc func(ctx *Context)

type Middleware func(next HandleFunc) HandleFunc

type OrginalContruct struct {
	ctx  *Context
	mdls []Middleware
}

func (o *OrginalContruct) Run(c *Context) {
	fmt.Println("the real logic run in Run Func")
}

func (o *OrginalContruct) ServeReady() {
	root := o.Run
	for i := len(o.mdls) - 1; i >= 0; i-- {
		root = o.mdls[i](root)
	}
	root(o.ctx)
}
