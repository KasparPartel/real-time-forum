package model

import "time"

type User struct {
	ID           int       `json:"id,omitempty"`
	Email        string    `json:"email,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	PasswordHash string    `json:"password,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	LoginTime    time.Time `json:"login_Time"`
	IsAdmin      string    `json:"is_admin"`
	//Age          int       `json:"age,omitempty"`
}

type Post struct {
	ID           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	UserID       int    `json:"user_id,omitempty"`
	Filename     string `json:"filename,omitempty"`
	CreationTime string `json:"creation_time"`
	UpdatedTime  string `json:"updated_time"`
}

type Comment struct {
	ID int `json:"id,omitempty"`
	//Post          Post   `json:"post"`
	Body          string `json:"body,omitempty"`
	Author        User   `json:"author"`
	LikeAmount    int    `json:"likeAmount,omitempty"`
	DislikeAmount int    `json:"dislikeAmount,omitempty"`
}
