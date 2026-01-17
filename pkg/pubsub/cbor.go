package pubsub

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type MsgPayload struct {
	topic  string
	sender string
	data   []byte
}

func encodeMsg(msg MsgPayload) []byte {
	arr := []interface{}{msg.topic, msg.sender, msg.data}
	cborData, err := cbor.Marshal(arr)
	if err != nil {
		panic(err)
	}

	const MAX_PAYLOAD_LEN = 1472
	if len(cborData) > MAX_PAYLOAD_LEN {
		output := fmt.Sprintf("cborData had length of %d", len(cborData))
		panic(output)
	}

	return cborData
}

func decodeMsg(payload []byte) MsgPayload {
	var arr []interface{}
	err := cbor.Unmarshal(payload, &arr)
	if err != nil {
		panic(err)
	}

	return MsgPayload{
		topic:  arr[0].(string),
		sender: arr[1].(string),
		data:   arr[2].([]byte),
	}
}
