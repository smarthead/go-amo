package webhooks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Webhook_MoveIntoStage(t *testing.T) {
	requestString := "leads[status][0][id]=2050297&" +
		"leads[status][0][status_id]=35573056&" +
		"leads[status][0][pipeline_id]=3643927&" +
		"leads[status][0][old_status_id]=35572897&" +
		"leads[status][0][old_pipeline_id]=3643927&" +
		"account[id]=29085955&" +
		"account[subdomain]=domain"

	expected := &WebhookRequest{
		Leads: Leads{
			Status: []Status{
				{
					Id:            "2050297",
					StatusId:      "35573056",
					PipelineId:    "3643927",
					OldStatusId:   "35572897",
					OldPipelineId: "3643927",
				},
			},
		},
		Account: Account{
			Id:        "29085955",
			SubDomain: "domain",
		},
	}

	webhook, err := NewFromString(requestString)

	if err != nil {
		t.Fail()
	}

	if assert.NotNil(t, webhook) {
		assert.Equal(t, webhook, expected)
	}
}

func Test_Webhook_CreateIntoStage(t *testing.T) {
	requestString := "leads[add][0][id]=2232929&" +
		"leads[add][0][status_id]=35573056&" +
		"leads[add][0][pipeline_id]=3643927&" +
		"account[id]=29085955&account[subdomain]=domain"

	expected := &WebhookRequest{
		Leads: Leads{
			Add: []Status{
				{
					Id:         "2232929",
					StatusId:   "35573056",
					PipelineId: "3643927",
				},
			},
		},
		Account: Account{
			Id:        "29085955",
			SubDomain: "domain",
		},
	}

	webhook, err := NewFromString(requestString)

	if err != nil {
		t.Fail()
	}

	if assert.NotNil(t, webhook) {
		assert.Equal(t, webhook, expected)
	}
}
