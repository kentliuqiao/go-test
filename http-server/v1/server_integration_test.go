package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinAndRetrievingThem(t *testing.T) {
	store := NewInMemoryStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newPlayerScoreRequest(player))

	assertStatus(t, resp.Code, http.StatusOK)

	assertResponseBody(t, resp.Body.String(), "3")
}
