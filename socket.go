package SocketDump

type SOCKET struct {
	Protocol      uint8
	LocalAddress  [4]uint8
	LocalPort     uint16
	RemoteAddress [4]uint8
	RemotePort    uint16
}

const BUFFER_SIZE uint32 = 65535
