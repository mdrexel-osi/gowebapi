
package gowebapi

type Landing struct {
	Links *LandingLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
