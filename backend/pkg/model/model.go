package model

type User struct {
	ID           int    `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	Gender       string `json:"gender,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Age          int    `json:"age,omitempty"`
	Username     string `json:"username,omitempty"`
	PasswordHash string `json:"-"`
	CreationTime string `json:"creation_time"`
	LoginTime    string `json:"login_Time"`
	LogoutTime   string `json:"logout_time"`
	IsAdmin      string `json:"is_admin"`
	Token        string `json:"-"`
	History 	 string `json:"history"`
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
	ID int `json:"id"`
	//ParentID     int    `json:"parent_id"`
	UserID       int    `json:"user_id"`
	PostID       int    `json:"post_id"`
	Body         string `json:"body"`
	CreationTime string `json:"creation_time"`
	//LikeAmount    int    `json:"likeAmount"`
	//DislikeAmount int    `json:"dislikeAmount"`
}
