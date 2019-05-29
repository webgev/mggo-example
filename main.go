package main 

import (
    "strings" 
    "github.com/webgev/mggo-example/controller"
    "github.com/webgev/mggo"
    "github.com/go-ini/ini"
)

func main() {
    temp := mggo.ViewData {
        DirView: "./view/",
        Template: "_template.html",
        Data: map[string]interface{}{},
    }
   
    rout := mggo.Router{
        GetController: getController,
        ViewData: temp,
        Menu: getMenu(),
    }
    cfg, err := ini.Load("./config.ini")
    if err != nil {
        os.Exit(1)
    }
    mggo.Run(rout, cfg)
}

func getMenu() mggo.Menu{
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
