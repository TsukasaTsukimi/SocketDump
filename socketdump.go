package SocketDump

import (
	"fmt"

	"github.com/imgk/divert-go"
	"github.com/shirou/gopsutil/process"
)

type Handle struct {
	windivert *divert.Handle
	dict      map[SOCKET]string
}

func NewSocketDump() (*Handle, error) {
	windivert, err := divert.Open("outbound", divert.LayerSocket, int16(0), divert.FlagSniff|divert.FlagRecvOnly)
	if err != nil {
		//log.Fatalf("Open windivert Handle Error: %v", err)
		return nil, err
	}
	handle := &Handle{
		windivert: windivert,
		dict:      make(map[SOCKET]string),
	}
	return handle, nil
}

func (handle *Handle) Process() error {
	buffer := make([]byte, BUFFER_SIZE)
	for {
		address := divert.Address{}
		_, err := handle.windivert.Recv(buffer, &address)
		if err != nil {
			//log.Fatalf("Failed to receive packet: %v", err)
			continue
		}

		switch address.Event() {
		case divert.EventSocketConnect:
			socket := SOCKET{
				Protocol:      address.Socket().Protocol,
				LocalAddress:  address.Socket().LocalAddress,
				LocalPort:     address.Socket().LocalPort,
				RemoteAddress: address.Socket().RemoteAddress,
				RemotePort:    address.Socket().RemotePort,
			}
			ProcessID := int32(address.Socket().ProcessID)
			Process, err := process.NewProcess(ProcessID)
			if err != nil {
				//log.Fatalf("Failed to create NewProcess: %v", err)
				continue
			}
			handle.dict[socket], _ = Process.Name()
			fmt.Println(socket.RemoteAddress)

		case divert.EventSocketClose:
			socket := SOCKET{
				Protocol:      address.Socket().Protocol,
				LocalAddress:  address.Socket().LocalAddress,
				LocalPort:     address.Socket().LocalPort,
				RemoteAddress: address.Socket().RemoteAddress,
				RemotePort:    address.Socket().RemotePort,
			}
			delete(handle.dict, socket)
		}
	}
}
