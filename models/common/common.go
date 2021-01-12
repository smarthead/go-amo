package common

type Values struct {
	Value  interface{} `json:"value,omitempty"`
	EnumID int         `json:"enum_id,omitempty"`
	Enum   string      `json:"enum,omitempty"`
}
type CustomFieldsValue struct {
	FieldID   int      `json:"field_id"`
	FieldName string   `json:"field_name,omitempty"`
	FieldCode string   `json:"field_code,omitempty"`
	FieldType string   `json:"field_type,omitempty"`
	Values    []Values `json:"values"`
}
