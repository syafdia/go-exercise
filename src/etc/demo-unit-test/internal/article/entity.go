package article

type Article struct {
	ID     int64
	Title  string
	Body   string
	UserID int64
}

type ArticleWithAuthor struct {
	Article
	UserFirstName string
	UserLastName  string
}

type CreateArticleInput struct {
	Title  string
	Body   string
	UserID int64
}
