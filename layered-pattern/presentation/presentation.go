package presentation

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/art-es/architecture-patterns/layered-pattern/application"
)

func ArticleListHandler(uc *application.ArticleListUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		if pageStr == "" {
			respondError(w, http.StatusBadRequest, errors.New("page: cannot be blank"))
			return
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, errors.New("page: must be string"))
			return
		}
		list, err := uc.Do(r.Context(), page)
		if err != nil {
			respondDummyError(w, r, err)
			return
		}
		respond(w, http.StatusOK, list)
	}
}
