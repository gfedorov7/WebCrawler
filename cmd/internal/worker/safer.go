package worker

import (
	"sync"
)

type SafeURLCollection struct {
	mux  sync.RWMutex
	urls map[string]struct{}
}

func (s *SafeURLCollection) Len() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.urls)
}

func (s *SafeURLCollection) Add(url string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.urls == nil {
		s.urls = make(map[string]struct{})
	}
	if _, exists := s.urls[url]; !exists {
		s.urls[url] = struct{}{}
	}
}

func (s *SafeURLCollection) Exists(url string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if s.urls == nil {
		return false
	}
	_, ok := s.urls[url]
	return ok
}
