package domain

import (
	"go_clean_arch_test/app/domain"
)

type AuthorGetChannel struct {
	author []domain.Author
	err    error
}

// constructor
func NewAuthorGetChannel(author []domain.Author, err error) (*AuthorGetChannel, error) {
	authorGetChannel := &AuthorGetChannel{
		author: author,
		err:    err,
	}

	return authorGetChannel, nil
}

// setter
func (authorGetChannel *AuthorGetChannel) Set(author []domain.Author, err error) error {
	authorGetChannel.author = author
	authorGetChannel.err = err

	return nil
}

// getter
func (authorGetChannel *AuthorGetChannel) GetAuthor() []domain.Author {
	return authorGetChannel.author
}

func (authorGetChannel *AuthorGetChannel) GetErr() error {
	return authorGetChannel.err
}
