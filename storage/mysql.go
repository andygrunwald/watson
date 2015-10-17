package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
)

type MySQLStorage struct {
	db     *sql.DB
	events chan *ChangeSet
}

func init() {
	register["mysql"] = &MySQLStorage{}
}

func (s *MySQLStorage) Init(u *url.URL, c chan *ChangeSet) error {
	s.events = c

	// Build MySQL connection
	var err error
	// Strip "mysql://" from DSN
	s.db, err = sql.Open(u.Scheme, u.String()[8:])
	if err != nil {
		return err
	}

	// Open doesn't open a connection. Validate DSN data
	err = s.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (s *MySQLStorage) Listen() {
	for c := range s.events {
		// TODO Consume events and store them
		log.Println("MySQLStorage: New Event", c)
	}
}

func (s *MySQLStorage) Close() error {
	s.db.Close()
	return nil
}
