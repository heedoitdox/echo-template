package utils

import (
	"testing"
)

func TestSendMessageToSlack(t *testing.T) {
	if err := SendMessageToSlack(); err != nil {
		t.Log(err)
	}
}
