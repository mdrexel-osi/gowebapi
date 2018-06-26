
package gowebapi

type UserInfo struct {
	IdentityType string `json:"IdentityType,omitempty"`

	Name string `json:"Name,omitempty"`

	IsAuthenticated bool `json:"IsAuthenticated,omitempty"`

	SID string `json:"SID,omitempty"`

	ImpersonationLevel string `json:"ImpersonationLevel,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
