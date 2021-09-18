package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"

	"example/web-service-gin/src/builders"
	"github.com/stretchr/testify/assert"
)


func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestUserRoute(t *testing.T) {
   router := setupRouter()
   user := builders.GenerateUser()
   w := httptest.NewRecorder()
   req, _ := http.NewRequest("GET", "/user/" + user.FirstName, nil)
   router.ServeHTTP(w, req)

   assert.Equal(t, http.StatusOK, w.Code)
   assert.Contains(t, w.Body.String(), user.FirstName)
}


func TestAdminRoute(t *testing.T) {
	postBody, _ := json.Marshal(builders.GenerateUser())
	b := bytes.NewBuffer(postBody)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin", b)
	req.Header.Add("authorization", "Basic Zm9vOmJhcg==")
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)
 
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
 }

