package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.AppendRight("User.Update", mggo.RRightManager)
	mggo.AppendRight("User.Read", mggo.RRightGuest)
	mggo.AppendRight("User.List", mggo.RRightGuest)

	mggo.InitCallback(func() {
		models := []interface{}{(*mggo.User)(nil), (*mggo.UserPassword)(nil)}
		mggo.CreateTable(models)
		mggo.Cache.AddMethod("User.List", mggo.CacheTypeMethodParams, 60*60*24)
		mggo.Cache.AddMethod("User.Read", mggo.CacheTypeUser, 60*60*24)
	})
	/*
	    mggo.EventSubscribe("SAP.Auth", func (params interface{}) {
	        fmt.Println(params)
		})*/
	// cache 1 day
}

type User struct {
	mggo.User       `mapstructure:",squash"`
	Password        string `sql:"-" structtomap:"-"`
	mggo.ListFilter `sql:"-" structtomap:"-" mapstructure:",squash"`
}

func (u User) Read(ctx *mggo.BaseContext) mggo.User {
	return u.User.Read(ctx)
}

func (u *User) List(ctx *mggo.BaseContext) (users []User) {
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

func (u User) Update(ctx *mggo.BaseContext) int {
	res := u.User.Update(ctx)
	if res != 0 && u.Password != "" {
		u.User.SetPassword(res, u.Password)
	}
	u.delCache()

	return res
}

func (u User) delCache() {
	mggo.Cache.ClearCacheMethod("User.List")
	mggo.Cache.ClearCacheMethodByUserID("User.Read", u.ID)
}

func (c User) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "user/user.html"
	data.Data["Title"] = "User"
	user := User{}
	data.Data["Users"] = mggo.Invoke(ctx, &user, "List")
}
