package response

import (
	"time"
)

// struct
type Article struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Author    Author    `json:"author"`
	ExpPool   ExpPool   `json:"exp_pool"`
}

// construcor
func NewArticle(id int, title, content string, updatedAt, createdAt time.Time, author Author, expPool ExpPool) *Article {
	Article := new(Article)
	Article.Id = id
	Article.Title = title
	Article.Content = content
	Article.UpdatedAt = updatedAt
	Article.CreatedAt = createdAt
	Article.Author = author
	Article.ExpPool = expPool
	return Article
}
