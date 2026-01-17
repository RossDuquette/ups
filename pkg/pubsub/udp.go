package pubsub

import (
	"net"
)

const (
	UDP_PUB_PORT = "54230"
	UDP_SUB_PORT = "54231"
)

type PubSubConnection struct {
	node_name      string
	broker_address string
}

func Connect(node_name string, broker_address string) *PubSubConnection {
	connection := PubSubConnection{
		node_name:      node_name,
		broker_address: broker_address,
	}
	return &connection
}

func sendPacket(dest_addr string, dest_port string, data []byte) {
	addr, err := net.ResolveUDPAddr("udp", dest_addr+":"+dest_port)
	if err != nil {
		panic(err)
	}

	connection, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	_, err = connection.Write(data)
	if err != nil {
		panic(err)
	}
}
