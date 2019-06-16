package main

import (
	"net/http"

	"github.com/webgev/mggo"
	_ "github.com/webgev/mggo-example/controller"
)

type hooks struct {
	mggo.RouterHooks
}

func (h hooks) Before(r *mggo.Router, w http.ResponseWriter, req *http.Request) {

}

func (h hooks) After(r *mggo.Router, w http.ResponseWriter, req *http.Request) {
}

func main() {
	temp := mggo.ViewData{
		DirView:  "./view/",
		Template: "_template.html",
	}

	rout := mggo.Router{
		ViewData:    temp,
		Menu:        getMenu(),
		RouterHooks: hooks{},
	}

	mggo.Run(rout, "./config.ini")
}

func getMenu() mggo.Menu {
	menu := mggo.Menu{}
	menu.Append("catalog", "Catalog", "/catalog")
	menu.Append("user", "User", "/user")
	menu.Append("news", "News", "/news")
	return menu
}
