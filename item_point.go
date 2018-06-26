
package gowebapi

type ItemPoint struct {
	Identifier string `json:"Identifier,omitempty"`

	IdentifierType string `json:"IdentifierType,omitempty"`

	Object *Point `json:"Object,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`
}
