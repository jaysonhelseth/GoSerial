package main

import (
	"GoSerial/device"
	"fmt"
)

func main() {
	//println(state.Color)

	mockDevice := device.MockDevice()
	runDevice(mockDevice)

	rDevice := device.RealDevice("tty", 115200)
	runDevice(rDevice)
}

func runDevice(device *device.MySerial) {
	device.OpenPort()
	buf := make([]byte, 128)
	count, _ := device.Read(buf)
	fmt.Printf("%s\n", string(buf[:count]))
}
