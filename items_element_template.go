
package gowebapi

type ItemsElementTemplate struct {
	Items []ElementTemplate `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
