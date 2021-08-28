package domain

// struct
type Article struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtをフィールドに持つ構造体
	// gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	// Author  Author `json:"author"`
	// Author  Author `json:"author"`
	// UpdatedAt time.Time `json:"updated_at"`
	// CreatedAt time.Time `json:"created_at"`
}

// // usecase interface
// type ArticleUsecase interface {
// 	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
// 	GetByID(ctx context.Context, id int64) (Article, error)
// 	Update(ctx context.Context, ar *Article) error
// 	GetByTitle(ctx context.Context, title string) (Article, error)
// 	Insert(context.Context, *Article) error
// 	Delete(ctx context.Context, id int64) error
// }

// // repository interface
// type ArticleRepository interface {
// 	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
// 	GetByID(ctx context.Context, id int64) (Article, error)
// 	GetByTitle(ctx context.Context, title string) (Article, error)
// 	Update(ctx context.Context, ar *Article) error
// 	Insert(ctx context.Context, a *Article) error
// 	Delete(ctx context.Context, id int64) error
// }
