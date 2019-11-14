package service

import (
	"net/http"

	"github.com/unrolled/render"
)



func InfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "table", struct {
			Name  string
			ID string
		}{Name: req.Form["name"][0], ID: req.Form["id"][0]})
	}
}
