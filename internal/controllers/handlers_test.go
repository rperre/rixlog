package controllers

import (
	"fmt"
	"io"
	_ "log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
	"log"
	"os"
	"path/filepath"
)

func TestHomeControllerIndexHander(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	v := Home()
	server := httptest.NewServer(http.HandlerFunc(v.Index))

	defer server.Close()
	httpResp, httpErr := http.Get(server.URL)
	assert.Assert(t, httpErr == nil)

	defer httpResp.Body.Close()
	assert.Assert(t, httpResp.StatusCode == http.StatusOK)

	body, IOerr := io.ReadAll(httpResp.Body)
	assert.Assert(t, IOerr == nil)
	fmt.Printf("\n%v\n", string(body))
	assert.Assert(t, string(body) == "{\"message\":\"no such table: articles\"}")
}

// TODO: Test all handlers
