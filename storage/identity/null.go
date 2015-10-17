package identity

import (
	"net/url"
)

type Null struct {
	events chan *Identity
}

func init() {
	register["null"] = &Null{}
}

func (s *Null) Init(u *url.URL, c chan *Identity) error {
	s.events = c
	return nil
}

func (s *Null) Listen() {
	<- s.events
}

func (s *Null) Close() error {
	return nil
}
