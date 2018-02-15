package ob

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x6/mmap"
)

type OB_Periph struct {
	RDP   RDP
	USER  USER
	DATA0 DATA0
	DATA1 DATA1
	WRP0  WRP0
}

func (p *OB_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OB = (*OB_Periph)(unsafe.Pointer(uintptr(mmap.OB_BASE)))

type RDP_Bits uint16

func (b RDP_Bits) Field(mask RDP_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDP_Bits) J(v int) RDP_Bits {
	return RDP_Bits(bits.Make32(v, uint32(mask)))
}

type RDP struct{ mmio.U16 }

func (r *RDP) Bits(mask RDP_Bits) RDP_Bits { return RDP_Bits(r.U16.Bits(uint16(mask))) }
func (r *RDP) StoreBits(mask, b RDP_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDP) SetBits(mask RDP_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *RDP) ClearBits(mask RDP_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *RDP) Load() RDP_Bits              { return RDP_Bits(r.U16.Load()) }
func (r *RDP) Store(b RDP_Bits)            { r.U16.Store(uint16(b)) }

type RDP_Mask struct{ mmio.UM16 }

func (rm RDP_Mask) Load() RDP_Bits   { return RDP_Bits(rm.UM16.Load()) }
func (rm RDP_Mask) Store(b RDP_Bits) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) RDP() RDP_Mask {
	return RDP_Mask{mmio.UM16{&p.RDP.U16, uint16(RDP)}}
}

func (p *OB_Periph) nRDP() RDP_Mask {
	return RDP_Mask{mmio.UM16{&p.RDP.U16, uint16(nRDP)}}
}

type USER_Bits uint16

func (b USER_Bits) Field(mask USER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask USER_Bits) J(v int) USER_Bits {
	return USER_Bits(bits.Make32(v, uint32(mask)))
}

type USER struct{ mmio.U16 }

func (r *USER) Bits(mask USER_Bits) USER_Bits { return USER_Bits(r.U16.Bits(uint16(mask))) }
func (r *USER) StoreBits(mask, b USER_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *USER) SetBits(mask USER_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *USER) ClearBits(mask USER_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *USER) Load() USER_Bits               { return USER_Bits(r.U16.Load()) }
func (r *USER) Store(b USER_Bits)             { r.U16.Store(uint16(b)) }

type USER_Mask struct{ mmio.UM16 }

func (rm USER_Mask) Load() USER_Bits   { return USER_Bits(rm.UM16.Load()) }
func (rm USER_Mask) Store(b USER_Bits) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) USER() USER_Mask {
	return USER_Mask{mmio.UM16{&p.USER.U16, uint16(USER)}}
}

func (p *OB_Periph) nUSER() USER_Mask {
	return USER_Mask{mmio.UM16{&p.USER.U16, uint16(nUSER)}}
}

type DATA0_Bits uint16

func (b DATA0_Bits) Field(mask DATA0_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DATA0_Bits) J(v int) DATA0_Bits {
	return DATA0_Bits(bits.Make32(v, uint32(mask)))
}

type DATA0 struct{ mmio.U16 }

func (r *DATA0) Bits(mask DATA0_Bits) DATA0_Bits { return DATA0_Bits(r.U16.Bits(uint16(mask))) }
func (r *DATA0) StoreBits(mask, b DATA0_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DATA0) SetBits(mask DATA0_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *DATA0) ClearBits(mask DATA0_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *DATA0) Load() DATA0_Bits                { return DATA0_Bits(r.U16.Load()) }
func (r *DATA0) Store(b DATA0_Bits)              { r.U16.Store(uint16(b)) }

type DATA0_Mask struct{ mmio.UM16 }

func (rm DATA0_Mask) Load() DATA0_Bits   { return DATA0_Bits(rm.UM16.Load()) }
func (rm DATA0_Mask) Store(b DATA0_Bits) { rm.UM16.Store(uint16(b)) }

type DATA1_Bits uint16

func (b DATA1_Bits) Field(mask DATA1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DATA1_Bits) J(v int) DATA1_Bits {
	return DATA1_Bits(bits.Make32(v, uint32(mask)))
}

type DATA1 struct{ mmio.U16 }

func (r *DATA1) Bits(mask DATA1_Bits) DATA1_Bits { return DATA1_Bits(r.U16.Bits(uint16(mask))) }
func (r *DATA1) StoreBits(mask, b DATA1_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DATA1) SetBits(mask DATA1_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *DATA1) ClearBits(mask DATA1_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *DATA1) Load() DATA1_Bits                { return DATA1_Bits(r.U16.Load()) }
func (r *DATA1) Store(b DATA1_Bits)              { r.U16.Store(uint16(b)) }

type DATA1_Mask struct{ mmio.UM16 }

func (rm DATA1_Mask) Load() DATA1_Bits   { return DATA1_Bits(rm.UM16.Load()) }
func (rm DATA1_Mask) Store(b DATA1_Bits) { rm.UM16.Store(uint16(b)) }

type WRP0_Bits uint16

func (b WRP0_Bits) Field(mask WRP0_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP0_Bits) J(v int) WRP0_Bits {
	return WRP0_Bits(bits.Make32(v, uint32(mask)))
}

type WRP0 struct{ mmio.U16 }

func (r *WRP0) Bits(mask WRP0_Bits) WRP0_Bits { return WRP0_Bits(r.U16.Bits(uint16(mask))) }
func (r *WRP0) StoreBits(mask, b WRP0_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *WRP0) SetBits(mask WRP0_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *WRP0) ClearBits(mask WRP0_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *WRP0) Load() WRP0_Bits               { return WRP0_Bits(r.U16.Load()) }
func (r *WRP0) Store(b WRP0_Bits)             { r.U16.Store(uint16(b)) }

type WRP0_Mask struct{ mmio.UM16 }

func (rm WRP0_Mask) Load() WRP0_Bits   { return WRP0_Bits(rm.UM16.Load()) }
func (rm WRP0_Mask) Store(b WRP0_Bits) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) WRP0() WRP0_Mask {
	return WRP0_Mask{mmio.UM16{&p.WRP0.U16, uint16(WRP0)}}
}

func (p *OB_Periph) nWRP0() WRP0_Mask {
	return WRP0_Mask{mmio.UM16{&p.WRP0.U16, uint16(nWRP0)}}
}
