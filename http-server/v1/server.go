package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := PlayerServer{
		store: store,
	}

	router :=
		p.handler.Handle("/league", p.leagueHandler())
	p.handler.Handle("/players/", p.playersHandler())

	return &p
}

func (s *PlayerServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(rw, r)
}

func (s *PlayerServer) playersHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case http.MethodPost:
			s.processWin(rw, player)
		case http.MethodGet:
			s.showScore(rw, player)
		}
	}
}

func (s *PlayerServer) leagueHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}
}

func (s *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	s.store.RecordWin(player)
	rw.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)
	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}
