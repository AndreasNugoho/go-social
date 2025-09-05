package main

import (
	"log"
	"net/http"

	"github.com/AndreasNugoho/go-social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	log.Printf("Fetching feed with limit=%d, offset=%d, sort=%s", fq.Limit, fq.Offset, fq.Sort)
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(2), fq)
	log.Printf("Fetched %d posts", len(feed))
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
