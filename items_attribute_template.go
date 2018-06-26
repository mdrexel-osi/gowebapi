
package gowebapi

type ItemsAttributeTemplate struct {
	Items []AttributeTemplate `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
