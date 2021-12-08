package domain

import (
	"errors"
	"time"
)

// struct
type ArticleForm struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId  int       `json:"author_id"`
}

// constructor
func NewArticleForm(id int, title, content string, updatedAt, createdAt time.Time, authorId int) (*ArticleForm, error) {
	articleForm := &ArticleForm{
		Id:        id,
		Title:     title,
		Content:   content,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
		AuthorId:  authorId,
	}

	return articleForm, nil
}

// setter
func (articleForm *ArticleForm) Set(id int, title, content string, updatedAt, createdAt time.Time, authorId int) error {
	if title == "" {
		return errors.New("title is required")
	}
	articleForm.Id = id
	articleForm.Title = title
	articleForm.Content = content
	articleForm.UpdatedAt = updatedAt
	articleForm.CreatedAt = createdAt
	articleForm.AuthorId = authorId

	return nil
}
