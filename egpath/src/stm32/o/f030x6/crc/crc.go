// Peripheral: CRC_Periph  CRC calculation unit.
// Instances:
//  CRC  mmap.CRC_BASE
// Registers:
//  0x00 32  DR        Data register.
//  0x04  8  IDR       Independent data register.
//  0x08 32  CR        Control register.
//  0x10 32  INIT      Initial CRC value register.
//  0x14 32  RESERVED3 Reserved,                                                    0x14.
// Import:
//  stm32/o/f030x6/mmap
package crc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	RESET   CR_Bits = 0x01 << 0 //+ RESET the CRC computation unit bit.
	REV_IN  CR_Bits = 0x03 << 5 //+ REV_IN Reverse Input Data bits.
	REV_OUT CR_Bits = 0x01 << 7 //+ REV_OUT Reverse Output Data bits.
)

const (
	RESETn   = 0
	REV_INn  = 5
	REV_OUTn = 7
)
