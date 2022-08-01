package model

import (
	"fmt"

	"github.com/google/uuid"
	database "github.com/sam-app/hackernews/packages/db/postgress"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

func (post *Post) Save() (string, error) {
	id := uuid.New().String()
	post.ID = id
	result := database.Db.Create(&post)
	if result.Error != nil {
		return "", result.Error
	}
	return post.ID, nil
}

func (post *Post) GetAllPosts() ([]Post, error) {
	var posts []Post
	result := database.Db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (post *Post) GetPostById(id string) (
	Post, error) {
	var p Post
	p.ID = id
	result := database.Db.First(&p)
	if result.Error != nil {
		return p, result.Error
	}
	return p, nil
}

func GetUserPosts(userId string) ([]*Post, error) {
	var posts []*Post
	result := database.Db.Where("user_id = ?", userId).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

//delete post
func (p *Post) Delete() (Post, error) {
	var post Post
	database.Db.First(&post, p.ID)
	if post.ID == "" {
		return post, fmt.Errorf("post not found")
	}
	result := database.Db.Delete(&post)

	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}
