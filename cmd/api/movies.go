package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.rabin.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.bedRequestResponse(w, r, http.StatusBadRequest, err.Error())

		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) getMovieById(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "Test title",
		Year:      2022,
		Runtime:   12,
		Genres:    []string{"Action", "Horror", "Science"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
