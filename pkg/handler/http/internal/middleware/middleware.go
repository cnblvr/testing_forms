package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
)

type logSettings uint8

const (
	LogWithHeaders logSettings = 1 << iota
	LogWithBodies
)

func BuildMiddlewareLogger(settings logSettings) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method, path := r.Method, r.URL.Path
			if r.URL.RawQuery != "" {
				path += "?" + r.URL.RawQuery
			}
			var reqHeaders string
			if settings&LogWithHeaders != 0 {
				bts, _ := json.Marshal(r.Header)
				reqHeaders = string(bts)
			}
			var reqBody []byte
			if settings&LogWithBodies != 0 {
				reqBody, _ = io.ReadAll(r.Body)
				r.Body.Close()
				r.Body = io.NopCloser(bytes.NewReader(reqBody))
			}

			respBuf := newResponseBuffer(settings, w)
			next.ServeHTTP(respBuf, r)

			var message strings.Builder
			message.WriteString(fmt.Sprintf("%s\t%s\tstatus_code=%d", method, path, respBuf.StatusCode()))
			if settings&LogWithHeaders != 0 {
				message.WriteString(fmt.Sprintf("\nRequest Headers: %s", reqHeaders))
			}
			if settings&LogWithBodies != 0 {
				message.WriteString(fmt.Sprintf("\nRequest Body: %s", reqBody))
			}
			if settings&LogWithHeaders != 0 {
				message.WriteString(fmt.Sprintf("\nResponse Headers: %s", respBuf.GetHeaderJSON()))
			}
			if settings&LogWithBodies != 0 {
				message.WriteString(fmt.Sprintf("\nResponse Body: %s", respBuf.Data()))
			}
			log.Debug().Msg(message.String())
		})
	}
}

type responseBuffer struct {
	settings   logSettings
	original   http.ResponseWriter
	statusCode int
	data       *bytes.Buffer
}

func newResponseBuffer(settings logSettings, original http.ResponseWriter) *responseBuffer {
	w := &responseBuffer{
		settings:   settings,
		original:   original,
		statusCode: http.StatusOK,
	}
	if settings&LogWithBodies != 0 {
		w.data = bytes.NewBuffer(nil)
	}
	return w
}

func (w *responseBuffer) Header() http.Header {
	return w.original.Header()
}

func (w *responseBuffer) Write(data []byte) (int, error) {
	if w.settings&LogWithBodies != 0 {
		w.data.Write(data)
	}
	return w.original.Write(data)
}

func (w *responseBuffer) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.original.WriteHeader(statusCode)
}

func (w responseBuffer) StatusCode() int {
	return w.statusCode
}

func (w responseBuffer) Data() []byte {
	if w.data == nil {
		return []byte{}
	}
	return w.data.Bytes()
}

func (w responseBuffer) GetHeaderJSON() string {
	if w.settings&LogWithHeaders == 0 {
		return ""
	}
	bts, _ := json.Marshal(w.Header())
	return string(bts)
}
