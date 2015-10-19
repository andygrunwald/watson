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

type ApprovalInfo gerrit.ApprovalInfo

func (p *GitPersonInfo) Identify() *Identity {
	if len(p.Name) == 0 && len(p.Email) == 0 {
		return nil
	}

	i := &Identity{
		Name:  p.Name,
		Email: p.Email,
	}

	return i
}

func (p AccountInfo) Identify() *Identity {
	if len(p.Name) == 0 && len(p.Email) == 0 && len(p.Username) == 0 {
		return nil
	}

	i := &Identity{
		Name:     p.Name,
		Email:    p.Email,
		Username: p.Username,
	}

	return i
}

func (p ApprovalInfo) Identify() *Identity {
	if len(p.Name) == 0 && len(p.Email) == 0 && len(p.Username) == 0 {
		return nil
	}

	i := &Identity{
		Name:     p.Name,
		Email:    p.Email,
		Username: p.Username,
	}

	return i
}

func (p *EmailInfo) Identify() *Identity {
	if len(p.Email) == 0 {
		return nil
	}

	i := &Identity{
		Email: p.Email,
	}

	return i
}
