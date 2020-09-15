package models

import "errors"

var (
	AuthNativeParseFailed      = errors.New("auth native parse failed")
	AuthNativeUsernameNotFound = errors.New("auth native username not found")
	AuthNativePasswordNotFound = errors.New("auth native password not found")
)

type AuthNative struct {
	Username string `json:"username"`
	Password string `json:"password"`
}