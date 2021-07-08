package main

import (
	"GoSerial/device"
	"fmt"
)

func main() {
	//println(state.Color)

	mockDevice := device.MockDevice()
	runDevice(mockDevice)

	rDevice := device.RealDevice("/dev/tty.usbserial-0001", 9600)
	runDevice(rDevice)
}

func runDevice(device *device.MySerial) {
	device.OpenPort()
	bytes, _ := device.Read()
	fmt.Printf("%s\n", string(bytes))
}
