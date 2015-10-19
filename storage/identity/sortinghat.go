package identity

import (
	"log"
	"net/url"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

type StortinghatStorage struct {
	events chan *Identity
	mode int
	u *url.URL
	client *http.Client
	source string
}

const (
	DefaultSource = "gerrit"
	SortinghatAddEndpoint = "/v1.0/identities"
	ModeHTTP = iota
	ModeHTTPS
)

func init() {
	// HTTP
	register["sortinghat"] = &StortinghatStorage{
		mode: ModeHTTP,
	}
	// HTTPS
	register["sortinghats"] = &StortinghatStorage{
		mode: ModeHTTPS,
	}
}

func (s *StortinghatStorage) Init(u *url.URL, c chan *Identity) error {
	s.events = c
	s.u = s.buildApiURL(u)
	// TODO Make this configurable
	s.source = DefaultSource
	s.client = &http.Client{}

	return nil
}

func (s *StortinghatStorage) Listen() {
	for c := range s.events {
		err := s.add(c)
		if err != nil {
			log.Printf("Stortinghat: %s\n", err)
		}
	}
}

func (s *StortinghatStorage) Close() error {
	return nil
}

func (s *StortinghatStorage) buildApiURL(u *url.URL) *url.URL {
	// Set correct scheme
	switch s.mode {
	case ModeHTTP:
		u.Scheme = "http"
	case ModeHTTPS:
		u.Scheme = "https"
	}

	// Set correct path
	if strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path[0:len(u.Path)-1]
	}
	u.Path += SortinghatAddEndpoint
	return u
}

func (s *StortinghatStorage) add(i *Identity) error {
	v := url.Values{}
	v.Add("name", i.Name)
	v.Add("username", i.Username)
	v.Add("email", i.Email)
	v.Add("source", s.source)

	req, err := http.NewRequest("POST", s.u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	// Identity already exists
	case http.StatusConflict:
		return nil
	// Identity successful created
	case http.StatusCreated:
		return nil
	}

	// All other codes are unexpected
	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%+v", v)
		return fmt.Errorf("Sortinghat API responds with status code %d and error %s", resp.StatusCode, body)
	}

	return nil
}