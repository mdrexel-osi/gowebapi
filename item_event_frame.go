
package gowebapi

type ItemEventFrame struct {
	Identifier string `json:"Identifier,omitempty"`

	IdentifierType string `json:"IdentifierType,omitempty"`

	Object *EventFrame `json:"Object,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`
}
