
package gowebapi

type ItemsElementCategory struct {
	Items []ElementCategory `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
