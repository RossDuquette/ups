package pubsub

func (conn PubSubConnection) Publish(topic string, data []byte) {
	msg := MsgPayload{
		topic:  topic,
		sender: conn.node_name,
		data:   data,
	}
	cborMsg := encodeMsg(msg)
	sendPacket(conn.broker_address, UDP_PUB_PORT, cborMsg)
}
