
package gowebapi

type ItemAttribute struct {
	Identifier string `json:"Identifier,omitempty"`

	IdentifierType string `json:"IdentifierType,omitempty"`

	Object *Attribute `json:"Object,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`
}
