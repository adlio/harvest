package harvest

import (
	"fmt"
	"strconv"
	"time"
)

type ContactsResponse struct {
	PagedResponse
	Contacts []*Contact `json:"contacts"`
}

type Contact struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneOffice string     `json:"phone_office"`
	PhoneMobile string     `json:"phone_mobile"`
	Fax         string     `json:"fax"`
	Client      ClientStub `json:"client"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (a *API) GetContact(contactID int64, args Arguments) (contact *Contact, err error) {
	contact = &Contact{}
	path := fmt.Sprintf("/contacts/%v", contactID)
	err = a.Get(path, args, contact)
	return contact, err
}

func (a *API) GetClientContacts(clientID int64, args Arguments) (contacts []*Contact, err error) {
	args["client_id"] = strconv.FormatInt(clientID, 10)
	return a.GetContacts(args)
}

func (a *API) GetContacts(args Arguments) (contacts []*Contact, err error) {
	contacts = make([]*Contact, 0)
	contactsResponse := ContactsResponse{}
	path := fmt.Sprintf("/contacts")
	err = a.GetPaginated(path, args, &contactsResponse, func() {
		for _, c := range contactsResponse.Contacts {
			contacts = append(contacts, c)
		}
		contactsResponse.Contacts = make([]*Contact, 0)
	})
	return contacts, err
}
