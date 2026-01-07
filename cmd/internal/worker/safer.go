package worker

import (
	"sync"
)

type SafeURLCollection struct {
	mux  sync.Mutex
	urls map[string]struct{}
}

func (s *SafeURLCollection) Len() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return len(s.urls)
}

func (s *SafeURLCollection) AddIfNotExists(url string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.urls == nil {
		s.urls = make(map[string]struct{})
	}

	if _, exists := s.urls[url]; exists {
		return false
	}

	s.urls[url] = struct{}{}
	return true
}
