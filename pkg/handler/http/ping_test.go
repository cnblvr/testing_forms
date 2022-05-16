package handlerHttp

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Ping(t *testing.T) {
	srv := httptest.NewServer(New())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/ping")
	if err != nil {
		t.Error(err)
		return
	}
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(bts, []byte("pong")) {
		t.Errorf("response is not pong")
		return
	}
}
