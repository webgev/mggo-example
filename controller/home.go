package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("home", NewHome)
}

func NewHome() *Home {
	return &Home{}
}

type Home struct {
}

func (u Home) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "home/home.html"
	data.Data["Title"] = "Home"
	user := User{}
	data.Data["Users"] = user.List(ctx)
}
