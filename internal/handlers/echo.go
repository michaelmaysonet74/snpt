package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Echo struct {
	l *log.Logger
}

func NewEcho(l *log.Logger) *Echo {
	return &Echo{l}
}

func (e *Echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.l.Println("POST /echo")
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			"Error: "+err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.Write([]byte(string(body)))
}
