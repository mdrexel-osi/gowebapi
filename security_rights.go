
package gowebapi

type SecurityRights struct {
	OwnerWebId string `json:"OwnerWebId,omitempty"`

	SecurityItem string `json:"SecurityItem,omitempty"`

	UserIdentity string `json:"UserIdentity,omitempty"`

	Links *SecurityRightsLinks `json:"Links,omitempty"`

	CanAnnotate bool `json:"CanAnnotate,omitempty"`

	CanDelete bool `json:"CanDelete,omitempty"`

	CanExecute bool `json:"CanExecute,omitempty"`

	CanRead bool `json:"CanRead,omitempty"`

	CanReadData bool `json:"CanReadData,omitempty"`

	CanSubscribe bool `json:"CanSubscribe,omitempty"`

	CanSubscribeOthers bool `json:"CanSubscribeOthers,omitempty"`

	CanWrite bool `json:"CanWrite,omitempty"`

	CanWriteData bool `json:"CanWriteData,omitempty"`

	HasAdmin bool `json:"HasAdmin,omitempty"`

	Rights []string `json:"Rights,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
