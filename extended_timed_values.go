
package gowebapi

type ExtendedTimedValues struct {
	Items []ExtendedTimedValue `json:"Items,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
