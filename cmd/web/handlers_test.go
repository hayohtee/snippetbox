package main

import (
	"bytes"
	"github.com/hayohtee/snippetbox/internal/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()
	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rs.StatusCode)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(rs.Body)

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
