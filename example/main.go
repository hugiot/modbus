package main

import (
	"github.com/hugiot/modbus"
	"log"
)

func main() {

}

func StartRTUSlave() {
	// config
	config := modbus.DefaultSerialConfig()
	config.Address = "/dev/pts/8"

	// handler
	handler := modbus.NewRTUServerHandler(config)

	// slave
	slave := modbus.NewSlave(handler)

	// function code handlers
	slave.HandleFuncCode(modbus.FuncCodeReadHoldingRegisters, func(c *modbus.Context) {
		//
		values, err := c.HoldingRegisters.Read(c.ProtocolMessage.Address, c.ProtocolMessage.Quantity)
		_ = c.Response(values, err)
	})
	slave.HandleFuncCode(modbus.FuncCodeWriteMultipleRegisters, func(c *modbus.Context) {
		err := c.HoldingRegisters.Write(c.ProtocolMessage.Address, c.GetUint16Slice())
		_ = c.Response(nil, err)
	})

	// listen
	if err := slave.Listen(); err != nil {
		log.Fatal(err)
	}
	defer slave.Close()
}
