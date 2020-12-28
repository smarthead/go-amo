package contacts

type Contact struct {
	ID                 int                  `json:"id"`
	Name               string               `json:"name"`
	FirstName          string               `json:"first_name"`
	LastName           string               `json:"last_name"`
	ResponsibleUserID  int                  `json:"responsible_user_id"`
	GroupID            int                  `json:"group_id"`
	CreatedBy          int                  `json:"created_by"`
	UpdatedBy          int                  `json:"updated_by"`
	CreatedAt          int                  `json:"created_at"`
	UpdatedAt          int                  `json:"updated_at"`
	ClosestTaskAt      interface{}          `json:"closest_task_at"`
	CustomFieldsValues []CustomFieldsValues `json:"custom_fields_values"`
	AccountID          int                  `json:"account_id"`
	Links              Links                `json:"_links"`
	Embedded           Embedded             `json:"_embedded"`
}
type Values struct {
	Value  string `json:"value"`
	EnumID int    `json:"enum_id"`
	Enum   string `json:"enum"`
}
type CustomFieldsValues struct {
	FieldID   int      `json:"field_id"`
	FieldName string   `json:"field_name"`
	FieldCode string   `json:"field_code"`
	FieldType string   `json:"field_type"`
	Values    []Values `json:"values"`
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
