
package gowebapi

import (
	"time"
)

type ExtendedTimedValue struct {
	Annotations []StreamAnnotation `json:"Annotations,omitempty"`

	Timestamp time.Time `json:"Timestamp,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	Good bool `json:"Good,omitempty"`

	Questionable bool `json:"Questionable,omitempty"`

	Substituted bool `json:"Substituted,omitempty"`

	Annotated bool `json:"Annotated,omitempty"`

	Value *interface{} `json:"Value,omitempty"`

	Errors []PropertyError `json:"Errors,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
