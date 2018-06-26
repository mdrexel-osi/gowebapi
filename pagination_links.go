
package gowebapi

type PaginationLinks struct {
	First string `json:"First,omitempty"`

	Previous string `json:"Previous,omitempty"`

	Next string `json:"Next,omitempty"`

	Last string `json:"Last,omitempty"`
}
