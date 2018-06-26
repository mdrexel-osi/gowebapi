
package gowebapi

type ItemsPointAttribute struct {
	Items []PointAttribute `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
