package conekta

import (
	"encoding/json"
	"testing"
	"bytes"
)

func TestExtractEventType(t *testing.T) {
	var event Event
	resp := []byte(Fixtures["event"])
	err := json.NewDecoder(bytes.NewReader(resp)).Decode(&event)
	if err != nil {
		t.Error("Could not decode the fixture:", err)
	}

	if event.Type != "charge.paid" {
		t.Errorf("event.Type => %v, want %v", event.Type, "charge.paid")
	}
}
