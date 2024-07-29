package helpers

import (
	"time"
)

type sendData struct {
	Data map[string]interface{}
}

type Posts struct {
	ID          int
	Title       string
	Content     string
	Author      string
	Categories  []string
	CreatedAt   time.Time
	ImgUrl      string
	BlogUrl     string
	CategoryUrl []string
	AvatarUrl   string
	Likes       int
	Dislikes    int
}
type Categories struct {
	ID          int
	Categories  []string
	CategoryUrl string
}

var SendData = sendData{
	Data: make(map[string]interface{}),
}

type User struct {
	ID         int
	Username   string
	Email      string
	ProfileUrl string
	AvatarUrl  string
	Content    string
}

type Comments struct {
	ID        int
	PostID    int
	Author    string
	Content   string
	CreatedAt time.Time
	AvatarUrl string
	Likes     int
	Dislikes  int
	EntryUrl  string
}
