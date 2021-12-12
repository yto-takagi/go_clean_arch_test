package domain

import (
	"go_clean_arch_test/app/domain"
)

type ArticleInputChannel struct {
	article domain.Article
	err     error
}

// constructor
func NewArticleInputChannel(article domain.Article, err error) (*ArticleInputChannel, error) {
	articleInputChannel := &ArticleInputChannel{
		article: article,
		err:     err,
	}

	return articleInputChannel, nil
}

// setter
func (articleInputChannel *ArticleInputChannel) Set(article domain.Article, err error) error {
	articleInputChannel.article = article
	articleInputChannel.err = err

	return nil
}

// getter
func (articleInputChannel *ArticleInputChannel) GetArticle() domain.Article {
	return articleInputChannel.article
}

func (articleInputChannel *ArticleInputChannel) GetErr() error {
	return articleInputChannel.err
}
