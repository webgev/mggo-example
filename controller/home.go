package controller

import (
	"github.com/webgev/mggo"
)

type Home struct {
}

func (u Home) IndexView(data *mggo.ViewData, path []string) {
	data.View = "home/home.html"
	data.Data["Title"] = "Home"
	user := User{}
	data.Data["Users"] = user.List()
}
