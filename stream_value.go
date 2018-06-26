
package gowebapi

type StreamValue struct {
	WebId string `json:"WebId,omitempty"`

	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	Value *TimedValue `json:"Value,omitempty"`

	Links *StreamValueLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
