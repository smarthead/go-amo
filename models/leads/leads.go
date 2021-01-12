package leads

import "github.com/qdimka/go-amo/models/common"

type Lead struct {
	ID                     int                         `json:"id"`
	Name                   string                      `json:"name,omitempty"`
	Price                  int                         `json:"price,omitempty"`
	ResponsibleUserID      int                         `json:"responsible_user_id,omitempty"`
	GroupID                int                         `json:"group_id,omitempty"`
	StatusID               int                         `json:"status_id,omitempty"`
	PipelineID             int                         `json:"pipeline_id,omitempty"`
	LossReasonID           int                         `json:"loss_reason_id,omitempty"`
	SourceID               interface{}                 `json:"source_id,omitempty"`
	CreatedBy              int                         `json:"created_by,omitempty"`
	UpdatedBy              int                         `json:"updated_by,omitempty"`
	CreatedAt              int                         `json:"created_at,omitempty"`
	UpdatedAt              int                         `json:"updated_at,omitempty"`
	ClosedAt               int                         `json:"closed_at,omitempty"`
	ClosestTaskAt          interface{}                 `json:"closest_task_at,omitempty"`
	IsDeleted              bool                        `json:"is_deleted,omitempty"`
	CustomFieldsValues     []*common.CustomFieldsValue `json:"custom_fields_values,omitempty"`
	Score                  interface{}                 `json:"score,omitempty"`
	AccountID              int                         `json:"account_id,omitempty"`
	IsPriceModifiedByRobot bool                        `json:"is_price_modified_by_robot,omitempty"`
	Links                  Links                       `json:"_links,omitempty"`
	Embedded               Embedded                    `json:"_embedded,omitempty"`
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
