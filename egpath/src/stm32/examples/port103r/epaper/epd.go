package main

import (
	"delay"

	"stm32/hal/gpio"
	"stm32/hal/spi"
	"stm32/hal/system"
)

// Max. SPI SCK frequency: 10 MHz. IL037F spec., table on page 32:
//
//  tSCYCW >= 100 ns.
//
// CSN can be set low whole time. IL037F spec., figure on page 24:
//
//  CSB *CAN* be "H" between parameter-command and command-parameter. SCL and
//  SDA during CSB="H"is invalid.

type EPD struct {
	width, height uint16
	loop20ns      byte

	spi  *spi.Driver
	cs   gpio.Pin
	dc   gpio.Pin
	rst  gpio.Pin
	busy gpio.Pin
}

func (epd *EPD) Reset() {
	epd.cs.Set()
	epd.rst.Clear()
	delay.Millisec(200)
	epd.rst.Set()
	delay.Millisec(200)
	epd.loop20ns = byte((system.Core.Clock()*2 + 1e8 - 1) / 1e8)
}

func (epd *EPD) Begin() {
	epd.cs.Clear()
}

func (epd *EPD) End() {
	epd.cs.Set()
}

type Command byte

const (
	SetPanel          Command = 0x00
	SetPower          Command = 0x01
	PowerOff          Command = 0x02
	SetPowerOffSeq    Command = 0x03
	PowerOn           Command = 0x04
	PowerOnMesure     Command = 0x05
	BoosterSoftStart  Command = 0x06
	DisplayStartTx    Command = 0x10
	DataStop          Command = 0x11
	DisplayRefresh    Command = 0x12
	DisplayStartTxRed Command = 0x13
	SetVcomLUT        Command = 0x20
	SetWhiteLUT       Command = 0x21
	SetBlackLUT       Command = 0x22
	SetGray1LUT       Command = 0x23
	SetGray2LUT       Command = 0x24
	SetVcomRedLUT     Command = 0x25
	SetRed0LUT        Command = 0x26
	SetRed1LUT        Command = 0x27
	SetPLLControl     Command = 0x30
	SetVcomAndDataInt Command = 0x50
	SetTCON           Command = 0x60
	SetResolution     Command = 0x61
	SetVCMDC          Command = 0x82
)

func (epd *EPD) wait20ns() {
	delay.Loop(int(epd.loop20ns))
}

func (epd *EPD) Cmd(cmd Command) {
	epd.dc.Clear()
	//epd.cs.Clear()
	epd.spi.WriteReadByte(byte(cmd))
	epd.dc.Set()
}

func (epd *EPD) WriteByte(b byte) {
	epd.spi.WriteReadByte(b)
}

func (epd *EPD) Write(s []byte) {
	epd.spi.WriteRead(s, nil)
}

func (epd *EPD) Wait() {
	for epd.busy.Load() == 0 {
		delay.Millisec(4)
	}
}

//emgo:const
var (
	lut_vcom0 = [...]byte{
		0x0E, 0x14, 0x01, 0x0A, 0x06, 0x04, 0x0A, 0x0A,
		0x0F, 0x03, 0x03, 0x0C, 0x06, 0x0A, 0x00,
	}
	lut_w = [...]byte{
		0x0E, 0x14, 0x01, 0x0A, 0x46, 0x04, 0x8A, 0x4A,
		0x0F, 0x83, 0x43, 0x0C, 0x86, 0x0A, 0x04,
	}
	lut_b = [...]byte{
		0x0E, 0x14, 0x01, 0x8A, 0x06, 0x04, 0x8A, 0x4A,
		0x0F, 0x83, 0x43, 0x0C, 0x06, 0x4A, 0x04,
	}
	lut_g1 = [...]byte{
		0x8E, 0x94, 0x01, 0x8A, 0x06, 0x04, 0x8A, 0x4A,
		0x0F, 0x83, 0x43, 0x0C, 0x06, 0x0A, 0x04,
	}
	lut_g2 = [...]byte{
		0x8E, 0x94, 0x01, 0x8A, 0x06, 0x04, 0x8A, 0x4A,
		0x0F, 0x83, 0x43, 0x0C, 0x06, 0x0A, 0x04,
	}
	lut_vcom1 = [...]byte{
		0x03, 0x1D, 0x01, 0x01, 0x08, 0x23, 0x37, 0x37,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	lut_red0 = [...]byte{
		0x83, 0x5D, 0x01, 0x81, 0x48, 0x23, 0x77, 0x77,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	lut_red1 = [...]byte{
		0x03, 0x1D, 0x01, 0x01, 0x08, 0x23, 0x37, 0x37,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
)

/*
	black := make([]byte, epd.width*epd.height/4)
	//red := make([]byte, epd.width*epd.height/8)

	for i := 0; i < int(epd.width); i++ {
		x, y := i, i
		o := y*int(epd.width)/4 + x/4
		black[o] = 0xFF
	}
*/
