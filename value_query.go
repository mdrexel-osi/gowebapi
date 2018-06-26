
package gowebapi

type ValueQuery struct {
	AttributeName string `json:"AttributeName,omitempty"`

	AttributeUOM string `json:"AttributeUOM,omitempty"`

	AttributeValue *interface{} `json:"AttributeValue,omitempty"`

	SearchOperator string `json:"SearchOperator,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
