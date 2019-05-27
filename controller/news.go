package controller

import (
    "github.com/webgev/mggo"
    "strconv"
)
func init() {
    mggo.AppendRight("News.Read", mggo.RRightGuest)
    mggo.AppendRight("News.List", mggo.RRightGuest)
    mggo.AppendRight("News.Update", mggo.RRightEditor)
    mggo.AppendRight("News.Delete", mggo.RRightEditor)

    mggo.AppendViewRight("News.Update", mggo.RRightEditor)
    mggo.InitCallback(func () {
        mggo.CreateTable( []interface{}{ (*Message)(nil) } )
    })
}

type News struct {
    ID int
    Name string
    View NewsView `sql:"-" structtomap:"-"`
}

func(c News) Read() News {
    mggo.SQL().Select(&c)
    return c
}

func (c *News) List() (newss []News) {
    mggo.SQL().Model(c).Select(&newss)
    return 
}
func (c News) Update() int {
    if c.ID == 0 {
        mggo.SQL().Insert(&c)
    } else {
        mggo.SQL().Update(&c)
    }
    return c.ID
}
func (c News) Delete() {
    if c.ID != 0 {
        mggo.SQL().Delete(&c)
    }
}

type NewsView struct{}

func(v NewsView) Index(data *mggo.ViewData, path []string){
    data.View = "news/news.html"
    data.Data["Title"] = "News"
    c := News{}
    data.Data["Newss"] = c.List()
}
func(v NewsView) Read(data *mggo.ViewData, path []string){
    if len(path) > 2 {
        if i, err := strconv.Atoi(path[2]); err == nil {
            data.View = "news/read.html"
            c := News{ID: i}
            r := c.Read()
            if r.ID > 0 {
                data.Data["Title"] = r.Name
                data.Data["News"] = r
                return
            }
        } 
    }
    panic(mggo.ErrorViewNotFound{})
}
func(v NewsView) Update(data *mggo.ViewData, path []string){
    data.View = "news/Update.html"
    if len(path) > 2 {
        if i, err := strconv.Atoi(path[2]); err == nil {
            data.View = "news/Update.html"
            c := News{ID: i,}
            r := c.Read()
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
