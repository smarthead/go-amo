package companies

import "github.com/qdimka/go-amo/models/common"

type Company struct {
	ID                 int                        `json:"id"`
	Name               string                     `json:"name,omitempty"`
	ResponsibleUserID  int                        `json:"responsible_user_id,omitempty"`
	GroupID            int                        `json:"group_id,omitempty"`
	CreatedBy          int                        `json:"created_by,omitempty"`
	UpdatedBy          int                        `json:"updated_by,omitempty"`
	CreatedAt          int                        `json:"created_at,omitempty"`
	UpdatedAt          int                        `json:"updated_at,omitempty"`
	ClosestTaskAt      interface{}                `json:"closest_task_at,omitempty"`
	CustomFieldsValues []common.CustomFieldsValue `json:"custom_fields_values,omitempty"`
	AccountID          int                        `json:"account_id,omitempty"`
	Links              Links                      `json:"_links,omitempty"`
	Embedded           Embedded                   `json:"_embedded,omitempty"`
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
