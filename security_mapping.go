
package gowebapi

type SecurityMapping struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Account string `json:"Account,omitempty"`

	SecurityIdentityWebId string `json:"SecurityIdentityWebId,omitempty"`

	Links *SecurityMappingLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
