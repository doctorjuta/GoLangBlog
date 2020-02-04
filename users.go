package main

import (
	"context"
	"net/http"

	"github.com/go-session/session"
	"github.com/jinzhu/gorm"
)

type BlogUser struct{}

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);unique_index"`
	Email        string `gorm:"type:varchar(100);unique_index"`
	Active       bool
	PasswordHash string
}

type UserAuthIdent struct {
	gorm.Model
	UserID    int    `gorm:"type:text"`
	AuthIdent string `gorm:"type:varchar(100);unique_index"`
}

func (it *BlogUser) IsUserLogin(w http.ResponseWriter, r *http.Request) bool {
	store, err := session.Start(context.Background(), w, r)
	if err != nil {
		BG.handleErrors(err.Error())
		return false
	}
	authident, ok := store.Get("authident")
	if ok {
		var user_authident UserAuthIdent
		DB.Connect()
		defer DB.Close()
		DB.cnct.Where("authident = ?", authident).First(&user_authident)
		if user_authident.AuthIdent != "" {
			return true
		} else {
			return false
		}
	}
	return false
}
