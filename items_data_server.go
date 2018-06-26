
package gowebapi

type ItemsDataServer struct {
	Items []DataServer `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
