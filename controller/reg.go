package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.AppendRight("Reg.Request", mggo.RRightGuest)
	mggo.AppendRight("Reg.Registration", mggo.RRightGuest)
}

type Reg struct {
	Token string
	Email string
	Code  int
	User
}

func (r Reg) Request() string {
	valid := mggo.Validation{
		Type:  mggo.ValidationEmailType,
		Email: r.Email,
	}
	return valid.Create()
}

func (r Reg) Registration() int {
	valid := mggo.Validation{
		Type:  mggo.ValidationEmailType,
		Token: r.Token,
		Code:  r.Code,
	}

	if !valid.Verification() {
		return 0
	}
	return r.User.Update()
}

func (c Reg) IndexView(data *mggo.ViewData, path []string) {
	data.View = "reg/reg.html"
	data.Data["Title"] = "Reg"
}
