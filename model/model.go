package model

import (
	"net/http"
)

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname"`
	Email    string `form:"email" json:"email"`
}

func Success(m string, d interface{}) *R {
	return &R{
		Code:    http.StatusOK,
		Message: m,
		Data:    d,
	}
}

func Fail(m string, d interface{}) *R {
	return &R{
		Code:    http.StatusInternalServerError,
		Message: m,
		Data:    d,
	}
}
