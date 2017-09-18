package harvest

import (
	"fmt"
	"time"
)

type Contact struct {
	ID          int64     `json:"id"`
	ClientID    int64     `json:"client_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneOffice string    `json:"phone_office"`
	PhoneMobile string    `json:"phone_mobile"`
	Fax         string    `json:"fax"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *API) GetContact(contactID int64, args Arguments) (contact *Contact, err error) {
	contact = &Contact{}
	path := fmt.Sprintf("/contacts/%v", contactID)
	err = a.Get(path, args, contact)
	return contact, err
}

func (a *API) GetClientContacts(clientID int64, args Arguments) (contacts []*Contact, err error) {
	path := fmt.Sprintf("/clients/%v/contacts", clientID)
	err = a.Get(path, args, &contacts)
	return contacts, err
}

func (a *API) GetContacts(args Arguments) (contacts []*Contact, err error) {
	path := fmt.Sprintf("/contacts")
	err = a.Get(path, args, &contacts)
	return contacts, err
}
