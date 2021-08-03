package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/art-es/architecture-patterns/layered-pattern/domain/article"
	"github.com/art-es/architecture-patterns/layered-pattern/util/json"
)

type Client interface {
	Get(ctx context.Context, key string) ([]byte, error)
	SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type Cache struct {
	Client
}

func (c *Cache) GetPaginatedArticlesByPage(ctx context.Context, page int) (*article.PaginatedList, error) {
	data, err := c.Get(ctx, articlesByPage(page))
	if err != nil {
		return nil, err
	}

	var list article.PaginatedList
	if err = json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (c *Cache) SetPaginatedArticlesByPage(ctx context.Context, list *article.PaginatedList, page int) error {
	return c.SetJSON(ctx, articlesByPage(page), list, 2*time.Hour)
}

func articlesByPage(page int) string { return fmt.Sprintf("paginated-articles:page-%d", page) }
