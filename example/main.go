package main

import (
	"github.com/TsukasaTsukimi/SocketDump"
)

func main() {
	handle, _ := SocketDump.NewSocketDump()
	handle.Process()
}
