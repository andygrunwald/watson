package identity

import (
	"log"
	"net/url"
)

type StortinghatStorage struct {
	events chan *Identity
	u *url.URL
}

func init() {
	register["sortinghat"] = &StortinghatStorage{}
}

func (s *StortinghatStorage) Init(u *url.URL, c chan *Identity) error {
	s.events = c
	s.u = u

	return nil
}

func (s *StortinghatStorage) Listen() {
	for c := range s.events {
		// TODO Consume events and store them
		log.Println("StortinghatStorage: New Event", c)
	}
}

func (s *StortinghatStorage) Close() error {
	return nil
}
