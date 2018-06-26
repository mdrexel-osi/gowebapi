
package gowebapi

import (
	"time"
)

type Annotation struct {
	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Value *interface{} `json:"Value,omitempty"`

	Creator string `json:"Creator,omitempty"`

	CreationDate time.Time `json:"CreationDate,omitempty"`

	Modifier string `json:"Modifier,omitempty"`

	ModifyDate time.Time `json:"ModifyDate,omitempty"`

	Links *AnnotationLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
