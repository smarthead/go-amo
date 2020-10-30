package amo

import (
	"fmt"
	"github.com/qdimka/go-amo/api"
	"github.com/qdimka/go-amo/models"
)

type AmoClient struct {
	api *api.Client
}

func NewAmoClient(options *api.ClientOptions) (*AmoClient, error) {
	apiClient, err := api.NewClient(options)
	if err != nil {
		return nil, err
	}

	return &AmoClient{
		api: apiClient,
	}, nil
}

func (client *AmoClient) GetLead(leadId string) (*models.Lead, error) {
	deal := new(models.Lead)
	err := client.api.Get(fmt.Sprintf("/api/v4/leads/%s", leadId), deal)
	return deal, err
}

func (client *AmoClient) GetUser(userId string) (*models.User, error) {
	user := new(models.User)
	err := client.api.Get(fmt.Sprintf("/api/v4/users/%s", userId), user)
	return user, err
}

func (client *AmoClient) UpdateLead(lead *models.Lead) error {
	err := client.api.Patch(fmt.Sprintf("/api/v4/leads/%d", lead.ID), lead, nil)
	return err
}
