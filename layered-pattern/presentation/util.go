package presentation

import (
	"errors"
	"log"
	"net/http"

	"github.com/art-es/architecture-patterns/layered-pattern/util/json"
)

var ErrDummy = errors.New("sorry! something went wrong, please try again later")

func respondDummyError(w http.ResponseWriter, where *http.Request, err error) {
	log.Printf("[ERROR] %s %q unexpected error: %v\n", where.Method, where.URL.RawPath, err)
	respondError(w, http.StatusInternalServerError, ErrDummy)
}

type errorResponse struct {
	Error error `json:"error"`
}

func respondError(w http.ResponseWriter, status int, err error) {
	respond(w, status, &errorResponse{err})
}

func respond(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("[ERROR] presentation respond: json.NewEncoder.Encode unexpected error: %v", err)
	}
}
