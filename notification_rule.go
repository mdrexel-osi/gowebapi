
package gowebapi

type NotificationRule struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AutoCreated bool `json:"AutoCreated,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	Criteria string `json:"Criteria,omitempty"`

	MultiTriggerEventOption string `json:"MultiTriggerEventOption,omitempty"`

	NonrepetitionInterval string `json:"NonrepetitionInterval,omitempty"`

	ResendInterval string `json:"ResendInterval,omitempty"`

	Status string `json:"Status,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
