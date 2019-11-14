
package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func UnknownHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "5xx Not found")
	}
}
