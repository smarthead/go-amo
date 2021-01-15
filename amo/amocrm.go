package amo

import (
	"fmt"
	"github.com/qdimka/go-amo/api"
	"github.com/qdimka/go-amo/models/companies"
	"github.com/qdimka/go-amo/models/contacts"
	"github.com/qdimka/go-amo/models/leads"
	"github.com/qdimka/go-amo/models/users"
)

type AmoClient struct {
	Api *api.Client
}

func NewAmoClient(options *api.ClientOptions) (*AmoClient, error) {
	apiClient, err := api.NewClient(options)
	if err != nil {
		return nil, err
	}

	return &AmoClient{
		Api: apiClient,
	}, nil
}

func (client *AmoClient) updateEntity(url string, id int, body interface{}) error {
	err := client.Api.Patch(fmt.Sprintf("%s/%d", url, id), body, nil)
	return err
}

func (client *AmoClient) GetUser(userId string) (*users.User, error) {
	user := new(users.User)
	err := client.Api.Get(fmt.Sprintf("/api/v4/users/%s", userId), user)
	return user, err
}

func (client *AmoClient) GetLead(leadId string, query string) (*leads.Lead, error) {
	deal := new(leads.Lead)
	resource := fmt.Sprintf("/api/v4/leads/%s", leadId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *AmoClient) UpdateLead(lead *leads.Lead) error {
	return client.updateEntity("/api/v4/leads", lead.ID, lead)
}

func (client *AmoClient) GetCompany(companyId string, query string) (*companies.Company, error) {
	deal := new(companies.Company)
	resource := fmt.Sprintf("/api/v4/companies/%s", companyId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *AmoClient) UpdateCompany(company *companies.Company) error {
	return client.updateEntity("/api/v4/companies", company.ID, company)
}

func (client *AmoClient) GetContact(contactId string, query string) (*contacts.Contact, error) {
	deal := new(contacts.Contact)
	resource := fmt.Sprintf("/api/v4/contacts/%s", contactId)
	if len(query) != 0 {
		resource = resource + "?" + query
	}

	err := client.Api.Get(resource, deal)
	return deal, err
}

func (client *AmoClient) UpdateContact(contact *contacts.Contact) error {
	return client.updateEntity("/api/v4/contacts", contact.ID, contact)
}
