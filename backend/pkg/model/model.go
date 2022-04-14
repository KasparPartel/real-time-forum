package model

import "time"

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Age       int    `json:"age,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	//Posts     []Post `json:"posts"`
	//Likes        []Like    `json:"likes"`
	//Comments     []Comment `json:"comments"`
	//Friends      []User    `json:"friends"`
	CreationTime time.Time `json:"creationTime"`
}

type Post struct {
	ID            int       `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Body          string    `json:"body,omitempty"`
	AuthorId      int       `json:"author_id,omitempty"`
	Filename      string    `json:"filename,omitempty"`
	LikeAmount    int       `json:"likeAmount,omitempty"`
	DislikeAmount int       `json:"dislikeAmount,omitempty"`
	CommentAmount int       `json:"commentAmount,omitempty"`
	CommentID     string    `json:"comments"`
	CreationTime  time.Time `json:"creationTime"`
}

type Comment struct {
	ID int `json:"id,omitempty"`
	//Post          Post   `json:"post"`
	Body          string `json:"body,omitempty"`
	Author        User   `json:"author"`
	LikeAmount    int    `json:"likeAmount,omitempty"`
	DislikeAmount int    `json:"dislikeAmount,omitempty"`
}
