package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentsRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents", nil)
	router.ServeHTTP(w, req)

	var expResult = "{\"documents\":[{\"ID\":\"61097a12f064fd0007ba1c18\",\"Name\":\"this is my first name\"},{\"ID\":\"61097a19f064fd0007ba1c1a\",\"Name\":\"this is my second name\"},{\"ID\":\"61097a22f064fd0007ba1c1c\",\"Name\":\"and my last name\"}]}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expResult, w.Body.String())
}