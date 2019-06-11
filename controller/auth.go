package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.RegisterController("auth", NewAuth)

	mggo.AppendRight("Auth.Authenticate", mggo.RRightGuest)
	mggo.AppendRight("Auth.Exit", mggo.RRightGuest)
}

func NewAuth() *Auth {
	return &Auth{}
}

// Auth - Функционал аутентификации
type Auth struct {
	Login    string
	Password string
}

// Authenticate - Аутентификацировать
func (a Auth) Authenticate(ctx *mggo.BaseContext) bool {
	return mggo.SAP{}.Authenticate(ctx, a.Login, a.Password)
}

// IsAuth - Проверит аутентифицирован ли пользователь
func (a Auth) IsAuth(ctx *mggo.BaseContext) bool {
	return mggo.SAP{}.IsAuth(ctx)
}

// Exit - Метод выхода
func (a Auth) Exit(ctx *mggo.BaseContext) {
	mggo.SAP{}.Exit(ctx)
}

// Index - главная страница
func (a Auth) IndexView(ctx *mggo.BaseContext, data *mggo.ViewData, path []string) {
	data.View = "auth/auth.html"
	data.Data["Title"] = "Auth"
}
