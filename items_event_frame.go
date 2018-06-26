
package gowebapi

type ItemsEventFrame struct {
	Items []EventFrame `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
