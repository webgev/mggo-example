package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("reg", NewReg)

	mggo.AppendRight("Reg.Request", mggo.RRightGuest)
	mggo.AppendRight("Reg.Registration", mggo.RRightGuest)
}

func NewReg() *Reg {
	return &Reg{}
}

type Reg struct {
	Token string
	Email string
	Code  int
	User
}

func (r Reg) Request(ctx *mggo.BaseContext) string {
	valid := mggo.Validation{
		Type:  mggo.ValidationEmailType,
		Email: r.Email,
	}
	return valid.Create()
}

func (r Reg) Registration(ctx *mggo.BaseContext) int {
	valid := mggo.Validation{
		Type:  mggo.ValidationEmailType,
		Token: r.Token,
		Code:  r.Code,
	}

	if !valid.Verification() {
		return 0
	}
	return r.User.Update(ctx)
}

func (c Reg) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "reg/reg.html"
	data.Data["Title"] = "Reg"
}
