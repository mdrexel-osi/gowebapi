
package gowebapi

import (
	"time"
)

type EventFrame struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	ExtendedProperties map[string]Value `json:"ExtendedProperties,omitempty"`

	StartTime time.Time `json:"StartTime,omitempty"`

	EndTime time.Time `json:"EndTime,omitempty"`

	Severity string `json:"Severity,omitempty"`

	AcknowledgedBy string `json:"AcknowledgedBy,omitempty"`

	AcknowledgedDate time.Time `json:"AcknowledgedDate,omitempty"`

	CanBeAcknowledged bool `json:"CanBeAcknowledged,omitempty"`

	IsAcknowledged bool `json:"IsAcknowledged,omitempty"`

	IsAnnotated bool `json:"IsAnnotated,omitempty"`

	IsLocked bool `json:"IsLocked,omitempty"`

	AreValuesCaptured bool `json:"AreValuesCaptured,omitempty"`

	RefElementWebIds []string `json:"RefElementWebIds,omitempty"`

	Security *Security `json:"Security,omitempty"`

	Links *EventFrameLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
