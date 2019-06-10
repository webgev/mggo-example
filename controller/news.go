package controller

import (
	"strconv"

	"github.com/webgev/mggo"
)

func init() {
	mggo.AppendRight("News.Read", mggo.RRightGuest)
	mggo.AppendRight("News.List", mggo.RRightGuest)
	mggo.AppendRight("News.Update", mggo.RRightEditor)
	mggo.AppendRight("News.Delete", mggo.RRightEditor)

	mggo.AppendViewRight("News.Update", mggo.RRightEditor)
	mggo.InitCallback(func() {
		mggo.CreateTable([]interface{}{(*News)(nil)})
	})
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

func (v News) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "news/news.html"
	data.Data["Title"] = "News"
	data.Data["Newss"] = v.List(ctx)
}
func (v News) ReadView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	if len(path) > 2 {
		if i, err := strconv.Atoi(path[2]); err == nil {
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
func (v News) UpdateView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "news/update.html"
	if len(path) > 2 {
		if i, err := strconv.Atoi(path[2]); err == nil {
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
