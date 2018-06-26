
package gowebapi

type SummaryValue struct {
	Type_ string `json:"Type,omitempty"`

	Value *TimedValue `json:"Value,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
