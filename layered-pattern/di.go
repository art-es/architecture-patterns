package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/art-es/architecture-patterns/layered-pattern/application"
	"github.com/art-es/architecture-patterns/layered-pattern/domain/article"
	"github.com/art-es/architecture-patterns/layered-pattern/domain/cache"
	"github.com/art-es/architecture-patterns/layered-pattern/infrastructure/postgres"
	"github.com/art-es/architecture-patterns/layered-pattern/infrastructure/redis"
	"github.com/art-es/architecture-patterns/layered-pattern/persistence"

	_ "github.com/lib/pq"
	"go.uber.org/dig"
)

func newDI() (*dig.Container, error) {
	c := dig.New()

	pp := []interface{}{
		newRedisClient,
		newCache,
		newDBPool,
		newArticleRepository,
		newArticlePaginator,
		newArticleListUsecase,
	}
	for _, p := range pp {
		if err := c.Provide(p); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func newDBPool() (*sql.DB, error) {
	port, err := getIntEnvOrDefault("DB_PORT", 5432)
	if err != nil {
		return nil, err
	}

	return postgres.Connect(postgres.Config{
		Host:         getEnvOrDefault("DB_HOST", "127.0.0.1"),
		Port:         port,
		User:         getEnvOrDefault("DB_USER", "postgres"),
		Password:     getEnvOrDefault("DB_PASSWORD", "postgres"),
		Database:     getEnvOrDefault("DB_NAME", "postgres"),
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})
}

func newRedisClient() (*redis.Client, error) {
	port, err := getIntEnvOrDefault("REDIS_PORT", 6379)
	if err != nil {
		return nil, err
	}

	db, err := getIntEnvOrDefault("REDIS_DB", 0)
	if err != nil {
		return nil, err
	}

	return redis.Connect(redis.Config{
		Addr:     fmt.Sprintf("%s:%d", getEnvOrDefault("REDIS_HOST", "127.0.0.1"), port),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: db,
	})
}

func newCache(redisClient *redis.Client) *cache.Cache {
	return &cache.Cache{redisClient}
}

func newArticleRepository(db *sql.DB) *persistence.ArticleRepository {
	return &persistence.ArticleRepository{
		DB:             db,
		BaseArticleURL: getEnvOrDefault("SITE_URL", "http://example.com") + "/articles/",
	}
}

func newArticlePaginator(articleRepository *persistence.ArticleRepository) (*article.Paginator, error) {
	perPage, err := getIntEnvOrDefault("ARTICLES_PERPAGE", 20)
	if err != nil {
		return nil, err
	}

	return &article.Paginator{
		ArticleRepository: articleRepository,
		PerPage:           perPage,
	}, nil
}

func newArticleListUsecase(paginator *article.Paginator, cacheObj *cache.Cache) *application.ArticleListUsecase {
	return &application.ArticleListUsecase{
		Paginator: paginator,
		Cache:     cacheObj,
	}
}

func getEnvOrDefault(varname, defaultvalue string) string {
	if value := os.Getenv(varname); value != "" {
		return value
	}
	return defaultvalue
}

func getIntEnvOrDefault(varname string, defaultvalue int) (int, error) {
	strvalue := os.Getenv(varname)
	if strvalue == "" {
		return defaultvalue, nil
	}

	intvalue, err := strconv.Atoi(strvalue)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer, %v", varname, err)
	}
	return intvalue, nil
}
