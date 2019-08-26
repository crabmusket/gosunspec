package modbus

import (
	"errors"

	"github.com/goburrow/modbus"
)

var ErrNotSupported = errors.New("not supported")
var ErrBadAddress = errors.New("bad address")

type simulator struct {
	memorymap   []byte
	baseAddress uint16
}

// OpenSimulator creates a modbus simulator atop the memory map defined
// by memory map. The baseAddress is specified by baseAddress.
//
// The only modbus.Client methods supported are ReadHoldingRegisters() and
// WriteMultipleRegisters() - all other methods return ErrNotSupported
func OpenSimulator(memorymap []byte, baseAddress uint16) (modbus.Client, error) {
	return &simulator{memorymap: memorymap, baseAddress: baseAddress}, nil
}

func (s *simulator) checkbounds(address uint16, quantity uint16) error {
	if address < s.baseAddress || (address+quantity > s.baseAddress+uint16(len(s.memorymap))/2) {
		return ErrBadAddress
	}
	return nil
}

func (s *simulator) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	if err := s.checkbounds(address, quantity); err != nil {
		return nil, err
	}
	byteOffset := (address - s.baseAddress) * 2
	return s.memorymap[byteOffset : byteOffset+quantity*2], nil
}

func (s *simulator) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	if err := s.checkbounds(address, quantity); err != nil {
		return nil, err
	}
	byteOffset := (address - s.baseAddress) * 2
	copy(s.memorymap[byteOffset:byteOffset+quantity*2], value)
	return s.memorymap[byteOffset : byteOffset+quantity*2], nil
}

func (s *simulator) ReadCoils(address, quantity uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) ReadDiscreteInputs(address, quantity uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) WriteSingleCoil(address, value uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) ReadInputRegisters(address, quantity uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}

func (s *simulator) ReadFIFOQueue(address uint16) (results []byte, err error) {
	return nil, ErrNotSupported
}
