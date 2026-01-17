package pubsub

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeMsg(t *testing.T) {
	// Test data
	originalMsg := MsgPayload{
		topic:  "test/topic",
		sender: "bob",
		data:   []byte("Hello, world!"),
	}

	// Encode the message
	encoded := encodeMsg(originalMsg)

	// Decode the message
	decoded := decodeMsg(encoded)

	// Verify the decoded message matches the original
	if decoded.topic != originalMsg.topic {
		t.Errorf("Expected topic %q, got %q", originalMsg.topic, decoded.topic)
	}

	if decoded.sender != originalMsg.sender {
		t.Errorf("Expected sender %q, got %q", originalMsg.sender, decoded.sender)
	}

	if !bytes.Equal(decoded.data, originalMsg.data) {
		t.Errorf("Expected data %q, got %q", originalMsg.data, decoded.data)
	}
}

func TestEncodeMsgSize(t *testing.T) {
	// Test with maximum allowed data size
	maxData := make([]byte, 1400) // Close to MAX_PAYLOAD_LEN
	for i := range maxData {
		maxData[i] = byte(i % 256)
	}

	msg := MsgPayload{
		topic:  "large/topic",
		sender: "bob",
		data:   maxData,
	}

	encoded := encodeMsg(msg)

	// Should not panic and should be within limits
	if len(encoded) > 1472 {
		t.Errorf("Encoded message exceeds MAX_PAYLOAD_LEN: %d > 1472", len(encoded))
	}

	// Verify it can be decoded correctly
	decoded := decodeMsg(encoded)

	if decoded.topic != msg.topic {
		t.Errorf("Topic mismatch after encoding large message")
	}

	if len(decoded.data) != len(maxData) {
		t.Errorf("Data length mismatch: expected %d, got %d", len(maxData), len(decoded.data))
	}
}

func TestEncodeDecodeEmptyData(t *testing.T) {
	msg := MsgPayload{
		topic:  "empty/data",
		sender: "charlie",
		data:   []byte{},
	}

	encoded := encodeMsg(msg)
	decoded := decodeMsg(encoded)

	if decoded.topic != msg.topic {
		t.Errorf("Empty data test: topic mismatch")
	}

	if decoded.sender != msg.sender {
		t.Errorf("Empty data test: topic mismatch")
	}

	if len(decoded.data) != 0 {
		t.Errorf("Empty data test: expected empty data, got %d bytes", len(decoded.data))
	}
}

func TestEncodeDecodeSpecialCharacters(t *testing.T) {
	specialData := []byte("Special chars: \x00\xff\x01\x02")
	msg := MsgPayload{
		topic:  "special/chars",
		sender: "user@domain.com",
		data:   specialData,
	}

	encoded := encodeMsg(msg)
	decoded := decodeMsg(encoded)

	if decoded.topic != msg.topic {
		t.Errorf("Special chars test: topic mismatch")
	}

	if decoded.sender != msg.sender {
		t.Errorf("Special chars test: sender mismatch")
	}

	if !bytes.Equal(decoded.data, specialData) {
		t.Errorf("Special chars test: data mismatch")
	}
}
