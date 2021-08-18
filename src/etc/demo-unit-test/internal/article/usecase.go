package article

import (
	"context"

	"github.com/syafdia/demo-unit-test/internal/user"
)

type WriteArticleUseCase interface {
	Execute(ctx context.Context, input CreateArticleInput) (ArticleWithAuthor, error)
}

type writeArticleUseCase struct {
	articleRepo ArticleRepo
	userRepo    user.UserRepo
}

func NewWriteArticleUseCase(
	articleRepo ArticleRepo,
	userRepo user.UserRepo,
) WriteArticleUseCase {
	return &writeArticleUseCase{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}

func (w *writeArticleUseCase) Execute(
	ctx context.Context,
	input CreateArticleInput,
) (ArticleWithAuthor, error) {
	article, err := w.articleRepo.Create(ctx, input)
	if err != nil {
		return ArticleWithAuthor{}, err
	}

	author, err := w.userRepo.FindOneByID(ctx, article.UserID)
	if err != nil {
		return ArticleWithAuthor{}, err
	}

	return ArticleWithAuthor{
		Article:       article,
		UserFirstName: author.FirstName,
		UserLastName:  author.LastName,
	}, nil
}
