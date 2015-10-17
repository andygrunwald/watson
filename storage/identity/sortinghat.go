package identity

import (
	"log"
)

type StortinghatStorage struct {
	events chan *Identity
}

func init() {
	register["sortinghat"] = &StortinghatStorage{}
}

func (s *StortinghatStorage) Init(c chan *Identity) {
	s.events = c
}

func (s *StortinghatStorage) Listen() {
	for c := range s.events {
		log.Println("StortinghatStorage: New Event", c)
	}
}
