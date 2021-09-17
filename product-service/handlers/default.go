package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Default struct {
	l *log.Logger
}

func (h *Default) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Default requests")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Uh oh. No endpoint found. %s", b)
}

func NewDefault(l *log.Logger) *Default {
	return &Default{l}
}
