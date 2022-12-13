package wework

import (
	workwx "github.com/xen0n/go-workwx" // package workwx
)

var (
	corpID           = "ww4bfd1ff0eb790d78"
	corpSecret       = "BCgfXmnUXagVKt7HMW2SFRwaEL0l764wMtAxjJjqwFg"
	agentID    int64 = 1000154
)

// Send ...
func Send(user, msg string) error {
	wxAppp := workwx.New(corpID).WithApp(corpSecret, agentID)
	rec := workwx.Recipient{
		UserIDs:  []string{user},
		PartyIDs: nil,
		TagIDs:   nil,
		ChatID:   "",
	}
	return wxAppp.SendTextMessage(&rec, msg, false)
}

// SendMarkdown ,,,
func SendMarkdown(user, msg string) error {
	wxAppp := workwx.New(corpID).WithApp(corpSecret, agentID)
	rec := workwx.Recipient{
		UserIDs:  []string{user},
		PartyIDs: nil,
		TagIDs:   nil,
		ChatID:   "",
	}
	return wxAppp.SendMarkdownMessage(&rec, msg, false)
}
