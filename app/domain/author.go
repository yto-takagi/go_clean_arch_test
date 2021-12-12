package domain

import (
	"errors"
	"time"
)

type Author struct {
	id        int       `json:"id"`
	name      string    `json:"name"`
	userId    int       `json:"user_id"`
	updatedAt time.Time `json:"updated_at"`
	createdAt time.Time `json:"created_at"`
}

// constructor
func NewAuthor(id int, name string, userId int, updatedAt, createdAt time.Time) (*Author, error) {
	author := &Author{
		id:        id,
		name:      name,
		userId:    userId,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}

	return author, nil
}

// setter
func (author *Author) Set(id int, name string, userId int, updatedAt, createdAt time.Time) error {
	if name == "" {
		return errors.New("name is required")
	}
	author.id = id
	author.name = name
	author.userId = userId
	author.updatedAt = updatedAt
	author.createdAt = createdAt

	return nil
}

// getter
func (author *Author) GetId() int {
	return author.id
}

func (author *Author) GetName() string {
	return author.name
}

func (author *Author) GetUserId() int {
	return author.userId
}

func (author *Author) GetUpdatedAt() time.Time {
	return author.updatedAt
}

func (author *Author) GetCreatedAt() time.Time {
	return author.createdAt
}
