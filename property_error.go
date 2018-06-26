
package gowebapi

type PropertyError struct {
	FieldName string `json:"FieldName,omitempty"`

	Message []string `json:"Message,omitempty"`
}
