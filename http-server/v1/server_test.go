package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayer(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("return Pepper's score", func(t *testing.T) {
		req := newPlayerScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := resp.Body.String()
		want := "20"

		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})
	t.Run("return Floyd's score", func(t *testing.T) {
		req := newPlayerScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		want := "10"
		got := resp.Body.String()

		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})
	t.Run("return 404 on missing player", func(t *testing.T) {
		req := newPlayerScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		want := "0"
		got := resp.Body.String()

		assertStatus(t, resp.Code, http.StatusNotFound)
		assertResponseBody(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(&store)

	t.Run("return accepted on POST", func(t *testing.T) {
		player := "Bob"
		req := newPostWinRequest(player)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("got %s want %s", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("returns 200 on /league", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func newPlayerScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect http status, got %d want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
