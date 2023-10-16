package context

import "net/http"

//var Context = NewContext()

type MsContext struct {
	Request  *http.Request
	W        http.ResponseWriter
	routers  map[string]func(ctx *MsContext)
	pathArgs map[string]map[string]string
}

//func NewContext() *MsContext {
//ctx := &MsContext{}
//ctx.routers = make(map[string]func(ctx2 *MsContext))
//}
