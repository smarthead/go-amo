package webhooks

import (
	"github.com/gorilla/schema"
	"log"
	"net/url"
	"strings"
)

type WebhookRequest struct {
	Leads   Leads   `schema:"leads"`
	Account Account `schema:"account"`
}

type Leads struct {
	Status []Status `schema:"status"`
	Add    []Status `schema:"add"`
}

type Status struct {
	Id            string `schema:"id"`
	StatusId      string `schema:"status_id"`
	PipelineId    string `schema:"pipeline_id"`
	OldStatusId   string `schema:"old_status_id"`
	OldPipelineId string `schema:"old_pipeline_id"`
}

type Account struct {
	Id        string `schema:"id"`
	SubDomain string `schema:"subdomain"`
}

func NewFromString(body string) (*WebhookRequest, error) {
	decodedBody, err := url.QueryUnescape(body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	replacer := strings.NewReplacer("][", ".", "[", ".", "]", "")
	decodedBody = replacer.Replace(decodedBody)

	bodyMap := make(map[string][]string)

	for _, value := range strings.Split(decodedBody, "&") {
		parameter := strings.Split(value, "=")
		bodyMap[parameter[0]] = []string{parameter[1]}
	}

	webhook := new(WebhookRequest)

	err = schema.NewDecoder().Decode(webhook, bodyMap)

	return webhook, err
}

func (hook *WebhookRequest) GetLeadId() string {
	if hook.Leads.Status != nil {
		return hook.Leads.Status[0].Id
	}
	return hook.Leads.Add[0].Id
}
