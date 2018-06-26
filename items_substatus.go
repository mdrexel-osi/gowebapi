
package gowebapi

type ItemsSubstatus struct {
	Items []Substatus `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
