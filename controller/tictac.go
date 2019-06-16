package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("tictac", NewTictac)

	mggo.AppendRight("Tictac.Connection", mggo.RRightGuest)
	mggo.AppendRight("Tictac.Step", mggo.RRightGuest)
}

func NewTictac() *Tictac {
	return &Tictac{}
}

type Tictac struct {
	ID   string
	User int

	Field int
}

var waitTictac = map[string]int{}
var connectionTictac = map[string]int{}

func (t Tictac) Step(ctx *mggo.BaseContext) {
	if id, ok := connectionTictac[t.ID]; ok {
		s := &mggo.SocketData{
			EventName: "Tictac.Step",
			Msg: mggo.MapStringAny{
				"id":    t.ID,
				"Field": t.Field,
			},
		}
		mggo.SendSocketUser(s, id)
	}
}
func (t Tictac) Connection(ctx *mggo.BaseContext) {
	if t.ID == "" {
		return
	}
	waitTictac[t.ID] = ctx.CurrentUser.ID
	for k, id := range waitTictac {
		if k != t.ID {
			s := mggo.SocketData{}
			s.EventName = "Tictac.Connect"
			s.Msg = mggo.MapStringAny{
				"id":   k,
				"step": 1,
			}
			mggo.SendSocketUser(&s, ctx.CurrentUser.ID)
			s.Msg = mggo.MapStringAny{
				"id":   t.ID,
				"step": 0,
			}
			mggo.SendSocketUser(&s, id)
			delete(waitTictac, t.ID)
			delete(waitTictac, k)
			connectionTictac[t.ID] = ctx.CurrentUser.ID
			connectionTictac[k] = id
			return
		}
	}
}

func (c Tictac) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	data.View = "tictac/tictac.html"
	data.Data["Title"] = "Tictac"
}
