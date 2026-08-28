package handlers

import (
	"net/http"

	"github.com/esmaeel67/golang-tdd.git/db"
)

// Handler contains the handler and all its dependencies.
type Handler struct {
	bs *db.BookService
	us *db.UserService
}

// NewHandler initializes a new handler, given dependencies
func NewHandler(bs *db.BookService, us *db.UserService) *Handler {
	return &Handler{
		bs: bs,
		us: us,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP status & a hardcoded message
	resp := &Response{
		Message: "Welcome to the BookSwap service!",
		Books:   h.bs.List(),
	}
	writeResponse(w, http.StatusOK, resp)
}
