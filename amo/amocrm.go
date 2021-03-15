package amo

import (
	"fmt"
	"github.com/qdimka/go-amo/api"
	"github.com/qdimka/go-amo/models/companies"
	"github.com/qdimka/go-amo/models/contacts"
	"github.com/qdimka/go-amo/models/leads"
	"github.com/qdimka/go-amo/models/users"
)

type IAmoClient interface {
	GetUser(userId string) (*users.User, error)
	GetLead(leadId string, query string) (*leads.Lead, error)
	UpdateLead(lead *leads.Lead) error
	GetCompany(companyId string, query string) (*companies.Company, error)
	UpdateCompany(company *companies.Company) error
	GetContact(contactId string, query string) (*contacts.Contact, error)
	UpdateContact(contact *contacts.Contact) error
}

type Client struct {
	Api *api.Client
}

//goland:noinspection GoUnusedExportedFunction
func NewAmoClient(options *api.ClientOptions) (*Client, error) {
	apiClient, err := api.NewClient(options)
	if err != nil {
		return nil, err
	}

	return &Client{
		Api: apiClient,
	}, nil
}

func (client *Client) updateEntity(url string, id int, body interface{}) error {
	err := client.Api.Patch(fmt.Sprintf("%s/%d", url, id), body, nil)
	return err
}

func (client *Client) GetUser(userId string) (*users.User, error) {
	user := new(users.User)
	err := client.Api.Get(fmt.Sprintf("/api/v4/users/%s", userId), user)
	return user, err
}

func (client *Client) GetLead(leadId string, query string) (*leads.Lead, error) {
	deal := new(leads.Lead)
	resource := fmt.Sprintf("/api/v4/leads/%s", leadId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *Client) UpdateLead(lead *leads.Lead) error {
	return client.updateEntity("/api/v4/leads", lead.ID, lead)
}

func (client *Client) GetCompany(companyId string, query string) (*companies.Company, error) {
	deal := new(companies.Company)
	resource := fmt.Sprintf("/api/v4/companies/%s", companyId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *Client) UpdateCompany(company *companies.Company) error {
	return client.updateEntity("/api/v4/companies", company.ID, company)
}

func (client *Client) GetContact(contactId string, query string) (*contacts.Contact, error) {
	deal := new(contacts.Contact)
	resource := fmt.Sprintf("/api/v4/contacts/%s", contactId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *Client) UpdateContact(contact *contacts.Contact) error {
	return client.updateEntity("/api/v4/contacts", contact.ID, contact)
}
