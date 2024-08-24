package main

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	var client *modbus.ModbusClient
	var err error

	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://127.0.0.1:5020",
		Timeout: 1 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Open()
	if err != nil {
		fmt.Println(err)
	}
	err = client.WriteRegisters(100, []uint16{
		0x0100, 0x0200, 0x0300, 0x0400,
	})
	if err != nil {
		fmt.Println(err)
	}
	reg16s, err := client.ReadBytes(100, 6, modbus.HOLDING_REGISTER)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reg16s)
	client.Close()
}
