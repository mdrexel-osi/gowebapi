
package gowebapi

type TimedValues struct {
	Items []TimedValue `json:"Items,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
