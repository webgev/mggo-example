package controller 

import (
    "github.com/webgev/mggo"
)

type Home struct {
    View HomeView
}
type HomeView struct {}

func (u HomeView) Index(data *mggo.ViewData, path []string) {
    data.View = "home/home.html"
    data.Data["Title"] = "Home"
    user := User{}
    data.Data["Users"] = user.List()
}