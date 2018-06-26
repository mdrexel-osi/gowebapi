
package gowebapi

type SecurityIdentity struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	IsEnabled bool `json:"IsEnabled,omitempty"`

	Links *SecurityIdentityLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
