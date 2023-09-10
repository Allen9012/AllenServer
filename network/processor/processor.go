package processor

type IProcessor interface {
	// MsgRoute must goroutine safe
	MsgRoute(clientId uint64, msg interface{}) error
	// UnknownMsgRoute must goroutine safe
	UnknownMsgRoute(clientId uint64, msg interface{})
	// ConnectedRoute connect event
	ConnectedRoute(clientId uint64)
	DisConnectedRoute(clientId uint64)

	// Unmarshal must goroutine safe
	Unmarshal(clientId uint64, data []byte) (interface{}, error)
	// Marshal must goroutine safe
	Marshal(clientId uint64, msg interface{}) ([]byte, error)
}

type IRawProcessor interface {
	IProcessor

	SetByteOrder(littleEndian bool)
	SetRawMsgHandler(handle RawMessageHandler)
	MakeRawMsg(msgType uint16, msg []byte, pbRawPackInfo *PBRawPackInfo)
	SetUnknownMsgHandler(unknownMessageHandler UnknownRawMessageHandler)
	SetConnectedHandler(connectHandler RawConnectHandler)
	SetDisConnectedHandler(disconnectHandler RawConnectHandler)
}
