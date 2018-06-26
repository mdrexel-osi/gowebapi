
package gowebapi

type ItemsTable struct {
	Items []Table `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
