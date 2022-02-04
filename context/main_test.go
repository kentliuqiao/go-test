package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyServe struct {
	response  string
	cancelled bool
}

func (s *SpyServe) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var res string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.cancelled = true
				return
			default:
				time.Sleep(10 * time.Millisecond)
				res += string(c)
			}
		}
		data <- res
	}()

	select {
	case <-ctx.Done():
		s.cancelled = true
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func (s *SpyServe) Cancel() {
	s.cancelled = true
}

func (s *SpyServe) assertCancelled(t *testing.T) {
	t.Helper()
	if !s.cancelled {
		t.Error("store should be cancelled")
	}
}

func (s *SpyServe) assertNotCancelled(t *testing.T) {
	t.Helper()
	if s.cancelled {
		t.Error("store should not be cancelled")
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write(data []byte) (int, error) {
	s.written = true
	return 0, errors.New("to be implemented")
}

func (s *SpyResponseWriter) WriteHeader(code int) {
	s.written = true
}

func TestServe(t *testing.T) {
	data := "Hello world"
	t.Run("returns data from store", func(t *testing.T) {
		spy := SpyServe{response: data}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		Serve(&spy).ServeHTTP(resp, req)
		if resp.Body.String() != data {
			t.Errorf("Serve got %s but want %s", resp.Body.String(), data)
		}
		spy.assertNotCancelled(t)
	})

	t.Run("cancel request after timeout", func(t *testing.T) {
		spy := SpyServe{response: data}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(ctx)

		resp := httptest.NewRecorder()

		Serve(&spy).ServeHTTP(resp, req)

		spy.assertCancelled(t)
	})

	t.Run("tell store to stop work if request is cancelled", func(t *testing.T) {
		spyServe := SpyServe{response: data}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(ctx)

		spyRW := SpyResponseWriter{}

		Serve(&spyServe).ServeHTTP(&spyRW, req)

		if spyRW.written {
			t.Error("response should not be written")
		}
	})
}
