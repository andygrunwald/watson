package storage

import (
	"github.com/andygrunwald/go-gerrit"
	"sync"
)

type Storage interface {
	Init(c chan *ChangeSet)
	Listen()
}

type ChangeSet struct {
	Change   *gerrit.ChangeInfo
	Comments *map[string]gerrit.CommentInfo
	Files    *map[string]gerrit.FileInfo
}

type Project gerrit.ProjectInfo

var (
	register = map[string]Storage{}
)

func GetStorage(storage string, wg *sync.WaitGroup) chan *ChangeSet {
	changeChan := make(chan *ChangeSet, 1)

	var s Storage
	var ok bool
	if s, ok = register[storage]; !ok {
		// TODO Add a null storage here
		s = register["mysql"]
	}

	s.Init(changeChan)
	wg.Add(1)
	go func() {
		s.Listen()
		wg.Done()
	}()

	return changeChan
}
