package article

import "context"

type Paginator struct {
	ArticleRepository ArticleRepository
	PerPage           int
}

type PaginatedList struct {
	Items []*Article `json:"articles"`
	Count int
}

func (p *Paginator) Paginate(ctx context.Context, page int) (*PaginatedList, error) {
	total, err := p.ArticleRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	offset, limit := (page-1)*p.PerPage, p.PerPage

	count := total - offset
	if count > limit {
		count = limit
	}
	if count < 0 {
		count = 0
	}

	items, err := p.ArticleRepository.GetByOffsetLimit(ctx, offset, limit, count)
	if err != nil {
		return nil, err
	}
	return &PaginatedList{Count: total, Items: items}, nil
}
