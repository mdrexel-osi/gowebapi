
package gowebapi

type SystemLanding struct {
	ProductTitle string `json:"ProductTitle,omitempty"`

	ProductVersion string `json:"ProductVersion,omitempty"`

	Links *SystemLandingLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
