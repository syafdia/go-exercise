package article

import "context"

type ArticleRepo interface {
	Create(ctx context.Context, input CreateArticleInput) (Article, error)
	FindOneByID(ctx context.Context, id int64) (Article, error)
	Destroy(ctx context.Context, id int64) error
}
