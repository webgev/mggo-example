package controller 

import (
    "github.com/webgev/mggo"
)
func init() {
    mggo.AppendRight("User.Update", mggo.RRightManager)
    mggo.AppendRight("User.Read", mggo.RRightGuest)
    mggo.AppendRight("User.List", mggo.RRightGuest)

    /*
    mggo.EventSubscribe("SAP.Auth", func (params interface{}) {
        fmt.Println(params)
    })*/
}

type User struct {
    mggo.User `mapstructure:",squash"`
    Password string `sql:"-" structtomap:"-"`
    View UserView `sql:"-" structtomap:"-"`
    mggo.ListFilter `sql:"-" structtomap:"-" mapstructure:",squash"`
}

func (u User) Read() mggo.User {
    return u.User.Read()
}

func (u *User) List() (users []User) {
    query := mggo.SQL().Model(&users)
    for key, value := range u.Filter {
        switch key {
        case "Name":
            query.Where("name = ?", value)
        case "Login":
            query.Where("login = ?", value)
        }
    }
    u.ListFilter.Paging(query).Select()
    return 
}

func (u User) Update() int {
    res := u.User.Update()
    if res != 0 && u.Password != "" {
        u.User.SetPassword(res, u.Password)
    }
    return res
}

type UserView struct {}

func (c UserView) Index(data *mggo.ViewData, path []string) {
    data.View = "user/user.html"
    data.Data["Title"] = "User"
    user := User{}
    data.Data["Users"] = user.List()
}