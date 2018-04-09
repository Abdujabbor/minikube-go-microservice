package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func welcome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("You are welcome"))
}