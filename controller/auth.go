package controller

import (
	"github.com/webgev/mggo"
)

func init() {
	mggo.AppendRight("Auth.Authenticate", mggo.RRightGuest)
	mggo.AppendRight("Auth.Exit", mggo.RRightGuest)
}

// Auth - Функционал аутентификации
type Auth struct {
	Login    string
	Password string
}

// Authenticate - Аутентификацировать
func (a Auth) Authenticate() bool {
	return mggo.SAP{}.Authenticate(a.Login, a.Password)
}

// IsAuth - Проверит аутентифицирован ли пользователь
func (a Auth) IsAuth() bool {
	return mggo.SAP{}.IsAuth()
}

// Exit - Метод выхода
func (a Auth) Exit() {
	mggo.SAP{}.Exit()
}

// Index - главная страница
func (a Auth) IndexView(data *mggo.ViewData, path []string) {
	data.View = "auth/auth.html"
	data.Data["Title"] = "Auth"
}
