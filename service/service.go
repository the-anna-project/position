// Package service provides a service to manage position peers within the
// connection space.
package service

import servicespec "github.com/the-anna-project/spec/service"

// New creates a new position service.
func New() servicespec.PeerService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		closer:   make(chan struct{}, 1),
		metadata: map[string]string{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	closer   chan struct{}
	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "position",
		"type": "service",
	}
}

func (s *service) Create(peer string) (string, error) {
	s.Service().Log().Line("func", "Create")

	// TODO

	var position string
	return position, nil
}

func (s *service) Delete(peer string) (string, error) {
	s.Service().Log().Line("func", "Delete")

	// TODO

	var position string
	return position, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Search(peer string) (string, error) {
	s.Service().Log().Line("func", "Search")

	// TODO

	var position string
	return position, nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
