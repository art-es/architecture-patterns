package persistence

import (
	"context"
	"database/sql"
	"log"

	"github.com/art-es/architecture-patterns/layered-pattern/domain/article"
)

type ArticleRepository struct {
	DB             *sql.DB
	BaseArticleURL string
}

func (r *ArticleRepository) Count(ctx context.Context) (count int, err error) {
	const query = "SELECT COUNT(*) FROM articles"
	if err = r.DB.QueryRowContext(ctx, query).Scan(&count); err != nil {
		log.Printf("[ERROR] ArticleRepository.Count: unexpected error: %v\n", err)
	}
	return
}

func (r *ArticleRepository) GetByOffsetLimit(ctx context.Context, offset, limit, count int) ([]*article.Article, error) {
	aa := make([]*article.Article, 0, count)
	if count == 0 {
		return aa, nil
	}

	const query = "SELECT title, desc, slug FROM articles OFFSET $1 LIMIT $2"
	rows, err := r.DB.QueryContext(ctx, query, offset, limit)
	if err == sql.ErrNoRows {
		return aa, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dto articleDTO
		if err = rows.Scan(&dto.Title, &dto.Desc, &dto.Slug); err != nil {
			return nil, err
		}

		aa = append(aa, dto.ToArticle(r.BaseArticleURL))
	}
	return aa, nil
}

type articleDTO struct {
	Title string
	Desc  string
	Slug  string
}

func (dto articleDTO) ToArticle(baseURL string) *article.Article {
	return &article.Article{
		Title: dto.Title,
		Desc:  dto.Desc,
		URL:   article.GetArticleURL(baseURL, dto.Slug),
	}
}
