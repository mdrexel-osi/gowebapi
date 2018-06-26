
package gowebapi

type ItemsItemElement struct {
	Items []ItemElement `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
