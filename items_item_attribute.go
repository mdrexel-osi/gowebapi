
package gowebapi

type ItemsItemAttribute struct {
	Items []ItemAttribute `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
