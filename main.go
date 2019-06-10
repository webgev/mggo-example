package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-ini/ini"
	"github.com/webgev/mggo"
	"github.com/webgev/mggo-example/controller"
)

type hooks struct {
	mggo.RouterHooks
}

func (h hooks) Before(r *mggo.Router, w http.ResponseWriter, req *http.Request) {
	/*command := "node ./cli/dist/server.js"
	parts := strings.Fields(command)
	data, _ := exec.Command(parts[0], parts[1:]...).Output()
	output := string(data)
	r.ViewData.Data["AAA"] = output*/
}

func (h hooks) After(r *mggo.Router, w http.ResponseWriter, req *http.Request) {
}

func main() {
	temp := mggo.ViewData{
		DirView:  "./view/",
		Template: "_template.html",
		Data:     map[string]interface{}{},
	}

	rout := mggo.Router{
		GetController: getController,
		ViewData:      temp,
		Menu:          getMenu(),
		RouterHooks:   hooks{},
	}
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		os.Exit(1)
	}
	mggo.Run(rout, cfg)
}

func getMenu() mggo.Menu {
	menu := mggo.Menu{}
	menu.Append("catalog", "Catalog", "/catalog")
	menu.Append("user", "User", "/user")
	return menu
}

func getController(controllerName string) interface{} {
	switch strings.ToLower(controllerName) {
	case "user":
		return &controller.User{}
	case "home":
		return &controller.Home{}
	case "auth":
		return &controller.Auth{}
	case "reg":
		return &controller.Reg{}
	case "catalog":
		return &controller.Catalog{}
	case "message":
		return &controller.Message{}
	case "news":
		return &controller.News{}
	}

	return nil
}
