package schemas

import (
	"errors"
	"strings"
	"time"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (body AddUserBody) isValid() bool {
	if strings.TrimSpace(body.Name) == "" {
		return false
	}
	if strings.TrimSpace(body.Email) == "" {
		return false
	}
	if body.CreatedAt.After(time.Now()) {
		return false
	}

	return true
}

func (body AddUserBody) Serialize() (AddUserBody, error) {
	if !body.isValid() {
		return AddUserBody{}, errors.New("invalid request")
	}

	if body.CreatedAt.IsZero() {
		body.CreatedAt = time.Now()
	}
	body.Name = cases.Title(language.English).String(strings.ToLower(body.Name))
	body.Email = strings.ToLower(body.Email)
	return body, nil
}
