package communication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type customMessage struct {
	Field int
}

func TestJsonPack(t *testing.T) {
	packer := JsonPacker(customMessage{123})
	data := packer()

	assert.JSONEq(t, `{"Field": 123}`, string(data))
}

func TestJsonUnpack(t *testing.T) {
	var message customMessage
	unpacker := JsonUnpacker(&message)
	unpacker([]byte(`{"Field": 123}`))

	assert.Equal(t, customMessage{123}, message)
}

func TestJsonListener(t *testing.T) {
	var messageConsumed customMessage
	listener := JsonListener(func(message customMessage) {
		messageConsumed = message
	})
	listener([]byte(`{"Field": 123}`))

	assert.Exactly(t, customMessage{123}, messageConsumed)
}

type customRequest struct {
	FieldIn string
}

type customResponse struct {
	FieldOut string
}

func TestJsonHandler(t *testing.T) {
	var requestReceived customRequest
	handler := JsonHandler(func(request customRequest) customResponse {
		requestReceived = request
		return customResponse{"RESPONSE"}
	})
	response := handler([]byte(`{"FieldIn": "REQUEST"}`))

	assert.Exactly(t, customRequest{"REQUEST"}, requestReceived)
	assert.JSONEq(t, `{"FieldOut": "RESPONSE"}`, string(response))
}