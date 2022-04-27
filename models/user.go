package models

type Users struct {
	UserId   int    `json:"userId" orm:"auto"`
	Email    string `json:"email" orm:"size(128)"`
	Password string `json:"password" orm:"size(64)"`
	UserName string `json:"user_name" orm:"size(32)"`
}
