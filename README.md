# Web Framework Application

Controller - api and view methods

## Get Started

Enter Go's path (format varies based on OS):

	cd $GOPATH

Install Mggo:
	go get -u github.com/webgev/mggo

Install Depends:
	go get -u github.com/mitchellh/mapstructure
	go get -u github.com/go-pg/pg
	go get -u github.com/gorilla/websocket
	go get -u github.com/go-ini/ini

Open http://localhost:9000 in your browser and you should see "It works!"
```shell
go get -u github.com/webgev/mggo

```

## Example

```go
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
	}
	
	return nil
}
```

## Depends

- https://github.com/mitchellh/mapstructure
- https://github.com/go-pg/pg
- https://github.com/gorilla/websocket
- https://github.com/go-ini/ini