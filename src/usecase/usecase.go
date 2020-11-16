package usecase

import (
	"encoding/json"
	"errors"
)

var (
	userError        = errors.New("User iteraction error")
	domainError      = errors.New("Domain exception catched")
	applicationError = errors.New("Something bad happened")
)

func concatErrors(errors []error) string {
	var messages []string

	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	message, _ := json.Marshal(messages)
	return string(message)
}

type UserError struct {
	errors []error
}

func (m *UserError) Error() string {
	return concatErrors(m.errors)
}

func NewUserError(previous error) *UserError {
	return &UserError{
		[]error{
			userError,
			previous,
		},
	}
}

type DomainError struct {
	errors []error
}

func (m *DomainError) Error() string {
	return concatErrors(m.errors)
}

func NewDomainError(previous error) *DomainError {
	return &DomainError{
		[]error{
			domainError,
			previous,
		},
	}
}

type ApplicationError struct {
	errors []error
}

func (m *ApplicationError) Error() string {
	return concatErrors(m.errors)
}

func NewApplicationError(previous error) *ApplicationError {
	return &ApplicationError{
		[]error{
			applicationError,
			previous,
		},
	}
}
