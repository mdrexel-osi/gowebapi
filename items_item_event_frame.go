
package gowebapi

type ItemsItemEventFrame struct {
	Items []ItemEventFrame `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
