package article

import (
	"context"
)

type Article struct {
	Title string
	Desc  string
	URL   string
}

func GetArticleURL(baseURL, slug string) string {
	if baseURL[len(baseURL)-1] == '/' {
		return baseURL + slug
	}
	return baseURL + "/" + slug
}

type ArticleRepository interface {
	Count(context.Context) (int, error)
	GetByOffsetLimit(ctx context.Context, offset, limit, count int) ([]*Article, error)
}
