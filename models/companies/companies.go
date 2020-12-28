package companies

type Company struct {
	ID                 int                  `json:"id"`
	Name               string               `json:"name"`
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
type Embedded struct {
	Tags []interface{} `json:"tags"`
}
