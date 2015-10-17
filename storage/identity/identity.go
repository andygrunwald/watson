package identity

import (
	"sync"
)

type Storage interface {
	Init(c chan *Identity)
	Listen()
}

type Identity struct {
	Name     string
	Email    string
	Username string
}

var (
	register = map[string]Storage{}
)

func GetStorage(storage string, wg *sync.WaitGroup) chan *Identity {
	identityChan := make(chan *Identity, 1)

	var s Storage
	var ok bool
	if s, ok = register[storage]; !ok {
		// TODO Add a null storage here
		s = register["sortinghat"]
	}
	s.Init(identityChan)
	wg.Add(1)
	go func() {
		s.Listen()
		wg.Done()
	}()

	return identityChan
}
