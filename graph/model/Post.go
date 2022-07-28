package model

import (
	"fmt"

	database "github.com/sam-app/hackernews/packages/db/postgress"
)

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

func (post *Post) Save() (int64, error) {
	// id := rand.Int()
	result := database.Db.Create(&Post{Title: post.Title, Desc: post.Desc, Content: post.Content})
	if result.Error != nil {
		return 0, result.Error
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
	result := database.Db.First(&p, id)
	fmt.Println("result: ", result)
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
