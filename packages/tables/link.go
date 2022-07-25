package tables

import (
	"log"

	"github.com/google/uuid"
	database "github.com/sam-app/hackernews/packages/db/postgress"
)

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	UserId  string `json:"user_id"`
}

func (link *Link) Save() string {
	id := uuid.New().String()
	result := database.Db.Create(&Link{ID: id, Title: link.Title, Address: link.Address})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return link.ID
}

func GetAllLinks() ([]Link, error) {
	var links []Link
	result := database.Db.Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}
