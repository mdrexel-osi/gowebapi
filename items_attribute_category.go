
package gowebapi

type ItemsAttributeCategory struct {
	Items []AttributeCategory `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
