package identity

import (
	"io"
	"net/url"
	"sync"
)

type Storage interface {
	Init(u *url.URL, c chan *Identity) error
	Listen()
	io.Closer
}

type Identity struct {
	Name     string
	Email    string
	Username string
}

const (
	DefaultStorage = "null"
)

var (
	register = map[string]Storage{}
)

func GetStorage(storage string, wg *sync.WaitGroup) (chan *Identity, Storage, error) {
	var s Storage
	identityChan := make(chan *Identity, 1)

	u, err := url.Parse(storage)
	if err != nil {
		return identityChan, s, err
	}

	var ok bool
	if s, ok = register[u.Scheme]; !ok {
		// TODO Add a null storage here
		s = register[DefaultStorage]
	}

	err = s.Init(u, identityChan)
	if err != nil {
		return identityChan, s, err
	}

	wg.Add(1)
	go func() {
		s.Listen()
		wg.Done()
	}()

	return identityChan, s, nil
}
