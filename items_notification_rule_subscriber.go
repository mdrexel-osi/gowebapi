
package gowebapi

type ItemsNotificationRuleSubscriber struct {
	Items []NotificationRuleSubscriber `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
