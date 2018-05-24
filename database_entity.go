package main

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type DbModel struct {
	gorm.Model
}

type IdentifiableDbModel struct {
	UUID string
}

func (e *IdentifiableDbModel) BeforeSave() error {
	if len(e.UUID) == 0 {
		e.UUID = uuid.New().String()
	}

	return nil
}

type DbContact struct {
	DbModel
	IdentifiableDbModel

	UserId string

	TitleBefore string
	FirstName   string
	MiddleName  string
	LastName    string
	TitleAfter  string

	Emails    []*DbEmail   `gorm:"foreignkey:ContactID"`
	Addresses []*DbAddress `gorm:"foreignkey:ContactID"`
	Phones    []*DbPhone   `gorm:"foreignkey:ContactID"`

	Birthday     *time.Time
	Note         string
	Organization string
}

func (DbContact) TableName() string {
	return "contact"
}

func (e *DbContact) UpdateWith(cnt2 *DbContact) {
	e.TitleBefore = cnt2.TitleBefore
	e.FirstName = cnt2.FirstName
	e.MiddleName = cnt2.MiddleName
	e.LastName = cnt2.LastName
	e.TitleAfter = cnt2.TitleAfter

	e.Emails = cnt2.Emails
	e.Addresses = cnt2.Addresses
	e.Phones = cnt2.Phones

	e.Birthday = cnt2.Birthday
	e.Note = cnt2.Note
	e.Organization = cnt2.Organization
}

type DbEmail struct {
	DbModel

	ContactID uint
	Type      string
	Email     string
}

func (DbEmail) TableName() string {
	return "email"
}

type DbAddress struct {
	DbModel

	ContactID uint
	Type      string
	Address   string
}

func (DbAddress) TableName() string {
	return "address"
}

type DbPhone struct {
	DbModel

	ContactID   uint
	Type        string
	PhoneNumber string
}

func (DbPhone) TableName() string {
	return "phone"
}
