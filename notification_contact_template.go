
package gowebapi

type NotificationContactTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Available bool `json:"Available,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	ContactType string `json:"ContactType,omitempty"`

	PlugInName string `json:"PlugInName,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	MaximumRetries int32 `json:"MaximumRetries,omitempty"`

	MinimumAcknowledgements int32 `json:"MinimumAcknowledgements,omitempty"`

	NotifyWhenInstanceEnded bool `json:"NotifyWhenInstanceEnded,omitempty"`

	EscalationTimeout string `json:"EscalationTimeout,omitempty"`

	RetryInterval string `json:"RetryInterval,omitempty"`

	Links *NotificationContactTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
