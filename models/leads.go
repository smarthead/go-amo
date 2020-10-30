package models

type Lead struct {
	ID                     int                  `json:"id"`
	Name                   string               `json:"name"`
	Price                  int                  `json:"price"`
	ResponsibleUserID      int                  `json:"responsible_user_id"`
	GroupID                int                  `json:"group_id"`
	StatusID               int                  `json:"status_id"`
	PipelineID             int                  `json:"pipeline_id"`
	LossReasonID           int                  `json:"loss_reason_id,omitempty"`
	SourceID               interface{}          `json:"source_id"`
	CreatedBy              int                  `json:"created_by"`
	UpdatedBy              int                  `json:"updated_by"`
	CreatedAt              int                  `json:"created_at"`
	UpdatedAt              int                  `json:"updated_at"`
	ClosedAt               int                  `json:"closed_at"`
	ClosestTaskAt          interface{}          `json:"closest_task_at"`
	IsDeleted              bool                 `json:"is_deleted"`
	CustomFieldsValues     []*CustomFieldsValue `json:"custom_fields_values"`
	Score                  interface{}          `json:"score"`
	AccountID              int                  `json:"account_id"`
	IsPriceModifiedByRobot bool                 `json:"is_price_modified_by_robot"`
	Links                  Links                `json:"_links"`
	Embedded               Embedded             `json:"_embedded"`
}

type Self struct {
	Href string `json:"href"`
}

type Links struct {
	Self Self `json:"self"`
}

type Tags struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Metadata struct {
	Quantity  int `json:"quantity"`
	CatalogID int `json:"catalog_id"`
}

type CatalogElements struct {
	ID       int      `json:"id"`
	Metadata Metadata `json:"metadata"`
}

type LossReason struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Links     Links  `json:"_links"`
}

type Companies struct {
	ID    int   `json:"id"`
	Links Links `json:"_links"`
}

type Contacts struct {
	ID     int   `json:"id"`
	IsMain bool  `json:"is_main"`
	Links  Links `json:"_links"`
}

type Embedded struct {
	Tags            []*Tags            `json:"tags"`
	CatalogElements []*CatalogElements `json:"catalog_elements"`
	LossReason      []*LossReason      `json:"loss_reason"`
	Companies       []*Companies       `json:"companies"`
	Contacts        []*Contacts        `json:"contacts"`
}

type Value struct {
	Value interface{} `json:"value"`
}

type CustomFieldsValue struct {
	FieldID   int      `json:"field_id"`
	FieldCode string   `json:"field_code"`
	Values    []*Value `json:"values"`
}
