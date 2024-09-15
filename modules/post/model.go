package post

import "time"

type PostRequest struct {
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

type PostResponse struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	UserId     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type PostResponseWithUsername struct {
	PostResponse
	Username string `json:"username"`
}

type PostUpdateRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

var (
	PostCollection []PostResponse
)
