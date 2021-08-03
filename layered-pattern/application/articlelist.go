package application

import (
	"context"
	"log"

	"github.com/art-es/architecture-patterns/layered-pattern/domain/article"
	"github.com/art-es/architecture-patterns/layered-pattern/domain/cache"
)

type ArticleListUsecase struct {
	Paginator *article.Paginator
	Cache     *cache.Cache
}

func (uc *ArticleListUsecase) Do(ctx context.Context, page int) (*article.PaginatedList, error) {
	// Retrieving a list from cache.
	list, err := uc.Cache.GetPaginatedArticlesByPage(ctx, page)
	if err == nil {
		return list, nil
	}

	// Is an internal error.
	if !cache.IsDoesNotExistError(err) {
		return nil, err
	}

	// Contacting the database.
	list, err = uc.Paginator.Paginate(ctx, page)
	if err != nil {
		return nil, err
	}

	// Save list to cache
	if err = uc.Cache.SetPaginatedArticlesByPage(ctx, list, page); err != nil {
		log.Printf("[ERROR] application.ArticleListUsecase: Cache.SetPaginatedArticlesByPage, unexpected error: %v\n", err)
	}
	return list, nil
}
