package model

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	CreationTime string `json:"creation_time"`
	LoginTime    string `json:"login_Time"`
	LogoutTime   string `json:"logout_time"`
	IsAdmin      string `json:"is_admin"`
	Token        string `json:"-"`
	History 	 string `json:"history"`
}

type Post struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	CategoryID   int    `json:"category_id"`
	UserID       int    `json:"user_id"`
	Filename     string `json:"filename"`
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
}

type Like struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type Dislike struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
