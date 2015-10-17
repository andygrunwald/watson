package identity

import (
	"github.com/andygrunwald/go-gerrit"
)

type Identify interface {
	Identify() Identity
}

type GitPersonInfo gerrit.GitPersonInfo

type AccountInfo gerrit.AccountInfo

type EmailInfo gerrit.EmailInfo

func (p *GitPersonInfo) Identify() *Identity {
	i := &Identity{
		Name:  p.Name,
		Email: p.Email,
	}

	return i
}

func (p AccountInfo) Identify() *Identity {
	i := &Identity{
		Name:     p.Name,
		Email:    p.Email,
		Username: p.Username,
	}

	return i
}
func (p *EmailInfo) Identify() *Identity {
	i := &Identity{
		Email: p.Email,
	}

	return i
}
