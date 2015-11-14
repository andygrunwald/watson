package storage

import (
	"github.com/andygrunwald/go-gerrit"
	"io"
	"net/url"
	"sync"
)

type Storage interface {
	Init(u *url.URL, c chan *ChangeSet) error
	Listen()
	io.Closer
}

type ChangeSet struct {
	Change   *gerrit.ChangeInfo
	Comments *map[string]gerrit.CommentInfo
	Files    *map[string]gerrit.FileInfo
}

type Project gerrit.ProjectInfo

const (
	DefaultStorage = "null"
)

var (
	register = map[string]Storage{}
)

func GetStorage(storage string, wg *sync.WaitGroup) (chan *ChangeSet, Storage, error) {
	var s Storage
	changeChan := make(chan *ChangeSet, 1)

	u, err := url.Parse(storage)
	if err != nil {
		return changeChan, s, err
	}

	var ok bool
	if s, ok = register[u.Scheme]; !ok {
		// TODO Add a null storage here
		s = register[DefaultStorage]
	}

	err = s.Init(u, changeChan)
	if err != nil {
		return changeChan, s, err
	}

	wg.Add(1)
	go func() {
		s.Listen()
		wg.Done()
	}()

	return changeChan, s, nil
}
