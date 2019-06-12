package controller

import (
	"fmt"
	"strconv"

	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("news", NewNews)

	mggo.AppendRight("News.Read", mggo.RRightGuest)
	mggo.AppendRight("News.List", mggo.RRightGuest)
	mggo.AppendRight("News.Update", mggo.RRightEditor)
	mggo.AppendRight("News.Delete", mggo.RRightEditor)

	mggo.AppendViewRight("News.Update", mggo.RRightEditor)
	mggo.InitCallback(func() {
		mggo.CreateTable([]interface{}{(*News)(nil)})
	})
}
func NewNews() *News {
	return &News{}
}

type News struct {
	ID   int
	Name string
}

func (c *News) Read(ctx *mggo.BaseContext) News {
	mggo.SQL().Select(c)
	return *c
}

func (c *News) List(ctx *mggo.BaseContext) (newss []News) {
	mggo.SQL().Model(c).Select(&newss)
	return
}
func (c News) Update(ctx *mggo.BaseContext) int {
	if c.ID == 0 {
		mggo.SQL().Insert(&c)
	} else {
		mggo.SQL().Update(&c)
	}
	return c.ID
}
func (c News) Delete(ctx *mggo.BaseContext) {
	if c.ID != 0 {
		mggo.SQL().Delete(&c)
	}
}
func (c News) ReadByName(ctx *mggo.BaseContext) News {
	if c.Name != "" {
		mggo.SQL().Model(&c).Where("name = ?", c.Name).Select()
	}
	return c
}

func (v News) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	data.View = "news/news.html"
	data.Data["Title"] = "News"
	data.Data["Newss"] = v.List(ctx)
}
func (v News) ReadView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	if len(ctx.Path) > 2 {
		if i, err := strconv.Atoi(ctx.Path[2]); err == nil {
			data.View = "news/read.html"
			c := News{ID: i}
			r := c.Read(ctx)
			if r.ID > 0 {
				data.Data["Title"] = r.Name
				data.Data["News"] = r
				return
			}
		}
	}
	panic(mggo.ErrorViewNotFound{})
}
func (v News) UpdateView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	data.View = "news/update.html"
	if len(ctx.Path) > 2 {
		if i, err := strconv.Atoi(ctx.Path[2]); err == nil {
			data.View = "news/update.html"
			c := News{ID: i}
			r := c.Read(ctx)
			if r.ID == 0 {
				panic(mggo.ErrorViewNotFound{})
			}
			data.Data["Title"] = r.Name
			data.Data["News"] = r
		} else {
			panic(mggo.ErrorViewNotFound{})
		}
	} else {
		data.Data["Title"] = "Ceate News"
		data.Data["News"] = News{}
	}
}
func (n News) View(ctx *mggo.BaseContext, data *mggo.ViewData) {
	fmt.Println(ctx.Path)
	if len(ctx.Path) > 1 {
		n.Name = ctx.Path[1]
		fmt.Println(n.Name)
		news := n.ReadByName(ctx)
		if news.ID != 0 {
			data.View = "news/read.html"
			data.Data["Title"] = news.Name
			data.Data["News"] = news
			return
		}
	}
	panic(mggo.ErrorViewNotFound{})
}
