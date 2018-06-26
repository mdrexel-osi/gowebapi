
package gowebapi

type ItemsNotificationRuleTemplate struct {
	Items []NotificationRuleTemplate `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
