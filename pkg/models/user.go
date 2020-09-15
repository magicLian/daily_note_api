package models

import (
	"errors"
	"time"
)

var (
	UserNotFound = errors.New("user not found")
)

type User struct {
	Id              string     `json:"id" gorm:"type:varchar(50);primary_key"`
	Username        string     `json:"username" gorm:"type:varchar(50)"`
	Password        string     `json:"password"`
	Age             int        `json:"age" gorm:"type:varchar(5)"`
	LastLoginTime   *time.Time `json:"lastLoginTime"`
	TokenUid        string     `json:"tokenUid"`
	Token           string     `json:"token"`
	TokenCreateTime *time.Time `json:"tokenCreateTime"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
}

type TokenUser struct {
	Id            string     `json:"id" gorm:"type:varchar(50);primary_key"`
	Username      string     `json:"username" gorm:"type:varchar(50)"`
	Password      string     `json:"password"`
	Age           int        `json:"age" gorm:"type:varchar(5)"`
	LastLoginTime *time.Time `json:"lastLoginTime"`
}
