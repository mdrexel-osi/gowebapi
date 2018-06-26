
package gowebapi

type ChannelInstance struct {
	Id string `json:"Id,omitempty"`

	StartTime string `json:"StartTime,omitempty"`

	LastMessageSentTime string `json:"LastMessageSentTime,omitempty"`

	SentMessageCount int32 `json:"SentMessageCount,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
