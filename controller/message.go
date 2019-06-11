package controller

import (
	"strconv"

	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("message", NewMessage)

	mggo.AppendRight("Message.Read", mggo.RRightUser)
	mggo.AppendRight("Message.List", mggo.RRightUser)
	mggo.AppendRight("Message.Send", mggo.RRightUser)
	mggo.AppendRight("Message.Delete", mggo.RRightUser)
	// all view
	mggo.AppendViewRight("Message", mggo.RRightUser)

	mggo.InitCallback(func() {
		mggo.CreateTable([]interface{}{(*Message)(nil)})
	})

}

func NewMessage() *Message {
	return &Message{}
}

type Message struct {
	ID       int    `mapstructure:"id"`
	UserID   int    `structtomap:"-"`
	ToUserID int    `mapstructure:"user_id"`
	Message  string `mapstructure:"message"`
}

func (c Message) Read(ctx *mggo.BaseContext) Message {
	mggo.SQL().Select(&c)
	return c
}

func (c *Message) Send(ctx *mggo.BaseContext) {
	if c.ToUserID == 0 || c.Message == "" {
		panic(mggo.ErrorInternalServer{"not user or message"})
	}
	if c.UserID == 0 {
		user := mggo.SAP{}.SessionUserID(ctx)
		if user == 0 {
			panic(mggo.ErrorAuthenticate{})
		}
		c.UserID = user
	}

	c.Update(ctx)
	mggo.EventPublish("message.send", mggo.EventTypeClient, []int{c.ToUserID}, c.Message)
}

func (c Message) List(ctx *mggo.BaseContext) (messages []Message) {
	m := new(Message)
	if m.UserID == 0 {
		m.UserID = mggo.SAP{}.SessionUserID(ctx)
	}
	query := mggo.SQL().Model(m)
	query.Where("user_id = ?", m.UserID)
	if m.ToUserID != 0 {
		query.Where("to_user_id = ?", m.ToUserID)
	}
	query.Select(&messages)
	return
}

func (c Message) Update(ctx *mggo.BaseContext) int {
	if c.ID == 0 {
		mggo.SQL().Insert(&c)
	} else {
		mggo.SQL().Update(&c)
	}
	return c.ID
}
func (c Message) Delete(ctx *mggo.BaseContext) {
	if c.ID != 0 {
		mggo.SQL().Delete(&c)
	}
}

type MessageView struct{}

func (v Message) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	data.View = "message/message.html"
	data.Data["Title"] = "Message"
	c := Message{}
	data.Data["Messages"] = c.List(ctx)
}

func (v Message) DialogView(ctx *mggo.BaseContext, data *mggo.ViewData) {
	if len(ctx.Path) > 2 {
		if i, err := strconv.Atoi(ctx.Path[2]); err == nil {
			c := Message{ToUserID: i}
			data.Data["ToUserID"] = i
			data.Data["Messages"] = c.List(ctx)
		}
	}
	data.View = "message/dialog.html"
	data.Data["Title"] = "Dialog"
}
