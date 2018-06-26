
package gowebapi

type SecurityEntry struct {
	Name string `json:"Name,omitempty"`

	SecurityIdentityName string `json:"SecurityIdentityName,omitempty"`

	AllowRights []string `json:"AllowRights,omitempty"`

	DenyRights []string `json:"DenyRights,omitempty"`

	Links *SecurityEntryLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
