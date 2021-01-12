package companies

import "github.com/qdimka/go-amo/models/common"

type Company struct {
	ID                 int                        `json:"id"`
	Name               string                     `json:"name"`
	ResponsibleUserID  int                        `json:"responsible_user_id"`
	GroupID            int                        `json:"group_id"`
	CreatedBy          int                        `json:"created_by"`
	UpdatedBy          int                        `json:"updated_by"`
	CreatedAt          int                        `json:"created_at"`
	UpdatedAt          int                        `json:"updated_at"`
	ClosestTaskAt      interface{}                `json:"closest_task_at"`
	CustomFieldsValues []common.CustomFieldsValue `json:"custom_fields_values"`
	AccountID          int                        `json:"account_id"`
	Links              Links                      `json:"_links"`
	Embedded           Embedded                   `json:"_embedded"`
}

type Self struct {
	Href string `json:"href"`
}

type Links struct {
	Self Self `json:"self"`
}

type Embedded struct {
	Tags []interface{} `json:"tags"`
}
