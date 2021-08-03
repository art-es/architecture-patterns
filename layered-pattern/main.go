package main

import (
	"log"
	"net/http"

	"github.com/art-es/architecture-patterns/layered-pattern/application"
	"github.com/art-es/architecture-patterns/layered-pattern/presentation"

	_ "github.com/lib/pq"
	"go.uber.org/dig"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
}

func run() error {
	server := &http.Server{Addr: ":8080"}

	{
		c, err := newDI()
		if err != nil {
			return err
		}
		if server.Handler, err = newHTTPHandler(c); err != nil {
			return err
		}
	}

	return server.ListenAndServe()
}

func newHTTPHandler(c *dig.Container) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	err := c.Invoke(func(uc *application.ArticleListUsecase) {
		mux.HandleFunc("/", presentation.ArticleListHandler(uc))
	})
	if err != nil {
		return nil, err
	}

	return mux, nil
}
