package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQLStorage struct {
	db     *sql.DB
	events chan *ChangeSet
}

func init() {
	register["mysql"] = &MySQLStorage{}
}

func (s *MySQLStorage) Init(c chan *ChangeSet) {
	s.events = c
}

func (s *MySQLStorage) Listen() {
	for c := range s.events {
		log.Println("MySQLStorage: New Event", c)
	}
}
