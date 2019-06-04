package controller

import (
	"strconv"

	"github.com/webgev/mggo"
)

func init() {
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

type Message struct {
	ID       int         `mapstructure:"id"`
	UserID   int         `structtomap:"-"`
	ToUserID int         `mapstructure:"user_id"`
	Message  string      `mapstructure:"message"`
	View     MessageView `sql:"-" structtomap:"-"`
}

func (c Message) Read() Message {
	mggo.SQL().Select(&c)
	return c
}

func (c *Message) Send() {
	if c.ToUserID == 0 || c.Message == "" {
		panic(mggo.ErrorInternalServer{"not user or message"})
	}
	if c.UserID == 0 {
		user := mggo.SAP{}.SessionUserID()
		if user == 0 {
			panic(mggo.ErrorAuthenticate{})
		}
		c.UserID = user
	}

	c.Update()
	mggo.EventPublish("message.send", mggo.EventTypeClient, []int{c.ToUserID}, c.Message)
}

func (c Message) List() (messages []Message) {
	m := new(Message)
	if m.UserID == 0 {
		m.UserID = mggo.SAP{}.SessionUserID()
	}
	query := mggo.SQL().Model(m)
	query.Where("user_id = ?", m.UserID)
	if m.ToUserID != 0 {
		query.Where("to_user_id = ?", m.ToUserID)
	}
	query.Select(&messages)
	return
}

func (c Message) Update() int {
	if c.ID == 0 {
		mggo.SQL().Insert(&c)
	} else {
		mggo.SQL().Update(&c)
	}
	return c.ID
}
func (c Message) Delete() {
	if c.ID != 0 {
		mggo.SQL().Delete(&c)
	}
}

type MessageView struct{}

func (v MessageView) Index(data *mggo.ViewData, path []string) {
	data.View = "message/message.html"
	data.Data["Title"] = "Message"
	c := Message{}
	data.Data["Messages"] = c.List()
}

func (v MessageView) Dialog(data *mggo.ViewData, path []string) {
	if len(path) > 2 {
		if i, err := strconv.Atoi(path[2]); err == nil {
			c := Message{ToUserID: i}
			data.Data["ToUserID"] = i
			data.Data["Messages"] = c.List()
		}
	}
	data.View = "message/dialog.html"
	data.Data["Title"] = "Dialog"
}
