package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server error:  %s path %s error %s", r.Method, r.URL, err)
	writeJSON(w, http.StatusInternalServerError, "The server encountered an internal error")

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error:  %s path %s error %s", r.Method, r.URL, err)
	writeJSON(w, http.StatusInternalServerError, err.Error())

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error %s path %s error %s", r.Method, r.URL, err)
	writeJSON(w, http.StatusInternalServerError, "not found error")

}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSON(w, http.StatusConflict, err.Error())
}
