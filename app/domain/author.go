package domain

import (
	"errors"
	"time"
)

type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// constructor
func NewAuthor(id int, name string, userId int, updatedAt, createdAt time.Time) (*Author, error) {
	author := &Author{
		Id:        id,
		Name:      name,
		UserId:    userId,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return author, nil
}

// setter
func (author *Author) Set(id int, name string, userId int, updatedAt, createdAt time.Time) error {
	if name == "" {
		return errors.New("name is required")
	}
	author.Id = id
	author.Name = name
	author.UserId = userId
	author.UpdatedAt = updatedAt
	author.CreatedAt = createdAt

	return nil
}
