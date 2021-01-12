package contacts

import "github.com/qdimka/go-amo/models/common"

type Contact struct {
	ID                 int                        `json:"id"`
	Name               string                     `json:"name,omitempty"`
	FirstName          string                     `json:"first_name,omitempty"`
	LastName           string                     `json:"last_name,omitempty"`
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
type Values struct {
	Value  string `json:"value"`
	EnumID int    `json:"enum_id"`
	Enum   string `json:"enum"`
}

type Self struct {
	Href string `json:"href"`
}

type Links struct {
	Self Self `json:"self"`
}

type Leads struct {
	ID    int   `json:"id"`
	Links Links `json:"_links"`
}

type Customers struct {
	ID    int   `json:"id"`
	Links Links `json:"_links"`
}

type Companies struct {
	ID    int   `json:"id"`
	Links Links `json:"_links"`
}

type Embedded struct {
	Tags            []interface{} `json:"tags"`
	Leads           []Leads       `json:"leads"`
	Customers       []Customers   `json:"customers"`
	CatalogElements []interface{} `json:"catalog_elements"`
	Companies       []Companies   `json:"companies"`
}
