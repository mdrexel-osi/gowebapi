
package gowebapi

type ItemsTableCategory struct {
	Items []TableCategory `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
