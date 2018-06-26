
package gowebapi

import (
	"time"
)

type AssetServer struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	IsConnected bool `json:"IsConnected,omitempty"`

	ServerVersion string `json:"ServerVersion,omitempty"`

	ServerTime time.Time `json:"ServerTime,omitempty"`

	ExtendedProperties map[string]Value `json:"ExtendedProperties,omitempty"`

	Links *AssetServerLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
