
package gowebapi

type ItemsAnnotation struct {
	Items []Annotation `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
