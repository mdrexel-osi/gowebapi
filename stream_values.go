
package gowebapi

type StreamValues struct {
	WebId string `json:"WebId,omitempty"`

	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	Items []TimedValue `json:"Items,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	Links *StreamValuesLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
