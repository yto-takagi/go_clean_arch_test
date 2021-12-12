package domain

import (
	"errors"
	"time"
)

// struct
type Article struct {
	// gorm.Model
	id        int       `json:"id"`
	title     string    `json:"title"`
	content   string    `json:"content"`
	updatedAt time.Time `json:"updated_at"`
	createdAt time.Time `json:"created_at"`
	Author    `json:"author"`
	ExpPool   `json:"exp_pool"`
}

// constructor
func NewArticle(id int, title, content string, updatedAt, createdAt time.Time, author Author, expPool ExpPool) (*Article, error) {

	if title == "" {
		return nil, errors.New("title is required")
	}

	article := &Article{
		id:        id,
		title:     title,
		content:   content,
		updatedAt: updatedAt,
		createdAt: createdAt,
		Author:    author,
		ExpPool:   expPool,
	}

	return article, nil
}

// setter
func (article *Article) Set(id int, title, content string, updatedAt, createdAt time.Time, author Author, expPool ExpPool) error {

	if title == "" {
		return errors.New("title is required")
	}
	article.id = id
	article.title = title
	article.content = content
	article.updatedAt = updatedAt
	article.createdAt = createdAt
	article.Author = author
	article.ExpPool = expPool

	return nil
}

// getter
func (article *Article) GetId() int {
	return article.id
}

func (article *Article) GetTitle() string {
	return article.title
}

func (article *Article) GetContent() string {
	return article.content
}

func (article *Article) GetUpdatedAt() time.Time {
	return article.updatedAt
}

func (article *Article) GetCreatedAt() time.Time {
	return article.createdAt
}

func (article *Article) GetAuthor() Author {
	return article.Author
}

func (article *Article) GetExpPool() ExpPool {
	return article.ExpPool
}
