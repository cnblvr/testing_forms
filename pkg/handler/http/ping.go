package handlerHttp

import (
	"io"
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
