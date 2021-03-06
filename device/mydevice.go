package device

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
)

type ISerial interface {
	Config()
	OpenPort()
}

type MySerial struct {
	ISerial
	config *serial.Config
	port   *serial.Port
	isMock bool
}

func MockDevice() *MySerial {
	myserial := &MySerial{}
	myserial.isMock = true

	return myserial
}

func RealDevice(name string, baud int) *MySerial {
	config := serial.Config{Name: name, Baud: baud}

	myserial := &MySerial{}
	myserial.config = &config
	myserial.isMock = false

	serial.OpenPort(myserial.config)
	return myserial
}

func (ms *MySerial) OpenPort() {
	if ms.isMock {
		println("Using Mock!")
		return
	}

	var err error
	ms.port, err = serial.OpenPort(ms.config)
	if err != nil {
		log.Fatal(err)
	}
}

func (ms *MySerial) Read() (bytes []byte, err error) {
	if ms.isMock {
		return GetData()
	}

	reader := bufio.NewReader(ms.port)
	reply, err := reader.ReadBytes('\x0a')

	return reply, nil
}
