package main

import (
	"errors"
	"time"
)

type ApiContact struct {
	UUID string `json:"uuid"`

	TitleBefore string `json:"title_before,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	TitleAfter  string `json:"title_after,omitempty"`

	Emails    []*ApiEmail   `json:"emails,omitempty"`
	Addresses []*ApiAddress `json:"addresses,omitempty"`
	Phones    []*ApiPhone   `json:"phones,omitempty"`

	Birthday     *time.Time `json:"birthday,omitempty"`
	Note         string     `json:"note,omitempty"`
	Organization string     `json:"organization,omitempty"`
}

func (e *ApiContact) Validate() error {
	var toValidate []Validation
	for _, e := range e.Emails {
		toValidate = append(toValidate, e)
	}

	for _, e := range e.Addresses {
		toValidate = append(toValidate, e)
	}

	for _, e := range e.Phones {
		toValidate = append(toValidate, e)
	}

	for _, o := range toValidate {
		if err := o.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ApiEmail struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

func (e *ApiEmail) Validate() error {
	if len(e.Type) == 0 {
		return errors.New("invalid email type")
	}

	if len(e.Email) == 0 {
		return errors.New("invalid email")
	}

	return nil
}

type ApiAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

func (e *ApiAddress) Validate() error {
	if len(e.Type) == 0 {
		return errors.New("invalid address type")
	}

	if len(e.Address) == 0 {
		return errors.New("invalid address")
	}

	return nil
}

type ApiPhone struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phone_number"`
}

func (e *ApiPhone) Validate() error {
	if len(e.Type) == 0 {
		return errors.New("invalid phone type")
	}

	if len(e.PhoneNumber) == 0 {
		return errors.New("invalid phone")
	}

	return nil
}

type Validation interface {
	Validate() error
}
