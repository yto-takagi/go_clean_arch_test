package domain

import (
	"errors"
	"time"
)

// struct
type Article struct {
	// gorm.Model
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Author    `json:"author"`
}

// constructor
func NewArticle(id int, title, content string, updatedAt, createdAt time.Time, author Author) (*Article, error) {

	if title == "" {
		return nil, errors.New("title is required")
	}

	article := &Article{
		Id:        id,
		Title:     title,
		Content:   content,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
		Author:    author,
	}

	return article, nil
}

// setter
func (article *Article) Set(id int, title, content string, updatedAt, createdAt time.Time, author Author) error {

	if title == "" {
		return errors.New("title is required")
	}
	article.Id = id
	article.Title = title
	article.Content = content
	article.UpdatedAt = updatedAt
	article.CreatedAt = createdAt
	article.Author = author

	return nil
}
