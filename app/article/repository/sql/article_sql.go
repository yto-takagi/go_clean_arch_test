package sql

import (
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 全件取得
func GetAll(db *gorm.DB, article []domain.Article, userId int) []domain.Article {
	// db.Order("created_at asc").Find(&articles)
	// db.Debug().Table("article").Find(&article)
	db.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.user_id = ?", userId).
		Scan(&article)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetAll"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)

	return article
}

// Idに紐付いたデータ取得
func GetById(db *gorm.DB, article domain.Article, id int) domain.Article {
	db.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("articles.id = ?", id).
		Scan(&article)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)

	return article

}

// Id、ユーザーIdに紐付いたデータ取得
func GetByIdAndUserId(db *gorm.DB, article domain.Article, id int, userId int) domain.Article {
	db.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("articles.id = ? AND authors.user_id = ?", id, userId).
		Scan(&article)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)

	return article

}

// authorId、ユーザーIdに紐付いたデータ取得
func GetByAuthorIdAndUserId(db *gorm.DB, articles []domain.Article, id int, userId int) []domain.Article {
	db.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.id = ? AND authors.user_id = ?", id, userId).
		Scan(&articles)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByAuthorIdAndUserId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articles)

	return articles

}

// titleとcontentを曖昧検索
func SearchContent(db *gorm.DB, articles []domain.Article, searchContent string, userId int) []domain.Article {
	db.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.user_id = ? AND (articles.title LIKE ? OR articles.content LIKE ?)",
			userId,
			"%"+searchContent+"%",
			"%"+searchContent+"%",
		).
		Scan(&articles)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "SearchContent"),
		zap.String("param searchContent", searchContent),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articles)

	return articles

}

// 新規登録
func Input(db *gorm.DB, articleForm *form.ArticleForm) {
	db.
		Debug().
		Table("articles").
		Create(&articleForm)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleForm)

}

// 更新
func Update(db *gorm.DB, articleForm *form.ArticleForm) {
	// TODO 更新SQL
	db.
		Debug().
		Model(&articleForm).
		Table("articles").
		Omit("createdAt").
		Where("id = ?", articleForm.Id).
		Updates(map[string]interface{}{
			"title":      articleForm.Title,
			"content":    articleForm.Content,
			"updated_at": articleForm.UpdatedAt,
			"author_id":  articleForm.AuthorId})

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleForm)

}

// 削除
func Delete(db *gorm.DB, articleForm *form.ArticleForm) {
	db.
		Debug().
		Table("articles").
		Where("id = ?", articleForm.Id).
		Delete(&articleForm)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}

// 削除(authorId指定)
func DeleteByAuthorId(db *gorm.DB, articleForm *form.ArticleForm) {
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■articleForm.AuthorId■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(articleForm.AuthorId)
	db.
		Debug().
		Table("articles").
		Where("author_id = ?", articleForm.AuthorId).
		Delete(&articleForm)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "DeleteByAuthorId"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
