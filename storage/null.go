package storage

import (
	"net/url"
)

type NullStorage struct {
	events chan *ChangeSet
}

func init() {
	register["null"] = &NullStorage{}
}

func (s *NullStorage) Init(u *url.URL, c chan *ChangeSet) error {
	s.events = c
	return nil
}

func (s *NullStorage) Listen() {
	<-s.events
}

func (s *NullStorage) Close() error {
	return nil
}
