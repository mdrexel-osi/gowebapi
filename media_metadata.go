
package gowebapi

import (
	"time"
)

type MediaMetadata struct {
	Author string `json:"Author,omitempty"`

	ChangeDate time.Time `json:"ChangeDate,omitempty"`

	Description string `json:"Description,omitempty"`

	Name string `json:"Name,omitempty"`

	Size float32 `json:"Size,omitempty"`

	Links *MediaMetadataLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
