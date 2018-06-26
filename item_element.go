
package gowebapi

type ItemElement struct {
	Identifier string `json:"Identifier,omitempty"`

	IdentifierType string `json:"IdentifierType,omitempty"`

	Object *Element `json:"Object,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`
}
