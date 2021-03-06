package sdcard

type Response [4]uint32

func (r Response) R1() CardStatus {
	return CardStatus(r[0])
}

func (r Response) R2CID() CID {
	return CID(r)
}

func (r Response) R2CSD() CSD {
	return CSD(r)
}

// R3 contains Operating Condition Register.
func (r Response) R3() OCR {
	return OCR(r[0])
}

// R6 contains Relative Card Address and Card Status bits 23, 22, 19, 12:0.
func (r Response) R6() (rca uint16, status CardStatus) {
	rca = uint16(r[0] >> 16)
	status = CardStatus(r[0]>>14&3<<22 | r[0]>>13&1<<19 | r[0]&0x1FFF)
	return
}

func (r Response) R7() (vhs VHS, pattern byte) {
	return VHS(r[0] >> 8 & 0xF), byte(r[0])
}

type CardStatus uint32

const (
	AKE_SEQ_ERROR      CardStatus = 1 << 3
	APP_CMD            CardStatus = 1 << 5
	FX_EVENT           CardStatus = 1 << 6
	READY_FOR_DATA     CardStatus = 1 << 8
	CURRENT_STATE      CardStatus = 15 << 9
	StateIdle          CardStatus = 0 << 9
	StateReady         CardStatus = 1 << 9
	StateIdent         CardStatus = 2 << 9
	StateStby          CardStatus = 3 << 9
	StateTran          CardStatus = 4 << 9
	StateData          CardStatus = 5 << 9
	StateRcv           CardStatus = 6 << 9
	StatePrg           CardStatus = 7 << 9
	StateDis           CardStatus = 8 << 9
	StateIOOnly        CardStatus = 15 << 9
	ERASE_RESET        CardStatus = 1 << 13
	CARD_ECC_DISABLED  CardStatus = 1 << 14
	WP_ERASE_SKIP      CardStatus = 1 << 15
	CSD_OVERWRITE      CardStatus = 1 << 16
	ERROR              CardStatus = 1 << 19
	CC_ERROR           CardStatus = 1 << 20
	CARD_ECC_FAILED    CardStatus = 1 << 21
	ILLEGAL_COMMAND    CardStatus = 1 << 22
	COM_CRC_ERROR      CardStatus = 1 << 23
	LOCK_UNLOCK_FAILED CardStatus = 1 << 24
	CARD_IS_LOCKED     CardStatus = 1 << 25
	WP_VIOLATION       CardStatus = 1 << 26
	ERASE_PARAM        CardStatus = 1 << 27
	ERASE_SEQ_ERROR    CardStatus = 1 << 28
	BLOCK_LEN_ERROR    CardStatus = 1 << 29
	ADDRESS_ERROR      CardStatus = 1 << 30
	OUT_OF_RANGE       CardStatus = 1 << 31
)

type OCR uint32

const (
	DVC   OCR = 1 << 7  // Dual Voltage Card
	V28   OCR = 1 << 15 // 2.7-2.8 V
	V29   OCR = 1 << 16 // 2.8-2.9 V
	V30   OCR = 1 << 17 // 2.9-3.0 V
	V31   OCR = 1 << 18 // 3.0-3.1 V
	V32   OCR = 1 << 19 // 3.1-3.2 V
	V33   OCR = 1 << 20 // 3.2-3.3 V
	V34   OCR = 1 << 21 // 3.3-3.4 V
	V35   OCR = 1 << 22 // 3.4-3.5 V
	V36   OCR = 1 << 23 // 3.5-3.6 V
	S18   OCR = 1 << 24 // Switching to 1.8V
	XPC   OCR = 1 << 18 // SDXC maximum performance
	UHSII OCR = 1 << 29 // UHS-II Card Status
	HCXC  OCR = 1 << 30 // Card Capacity Status (set fot SDHC, SDXC).
	PWUP  OCR = 1 << 31 // Card in power up state (^Busy).
)


// CID - Card Identification Register
type CID [4]uint32

// MID returns the manufacturer ID.
func (cid CID) MID() byte {
	return byte(cid[3] >> 24)
}

// OID returns the OEM/Application ID.
func (cid CID) OID() [2]byte {
	return [2]byte{byte(cid[3] >> 16), byte(cid[3] >> 8)}
}

// PNM returns the product name.
func (cid CID) PNM() [5]byte {
	return [5]byte{
		byte(cid[3]), byte(cid[2] >> 24), byte(cid[2] >> 16), byte(cid[2] >> 8),
		byte(cid[2]),
	}
}

// PRV returns the product revision.
func (cid CID) PRV() byte {
	return byte(cid[1] >> 24)
}

// PSN returns the product serial number.
func (cid CID) PSN() uint32 {
	return cid[1]<<8 | cid[0]>>24
}

// MDT returns the manufacturing date.
func (cid CID) MDT() (year, month int) {
	mdt := int(cid[0] >> 8 & 0xFFF)
	return mdt>>4 + 2000, mdt & 15
}

// CRC returns the 7-bit CRC field.
func (cid CID) CRC() byte {
	return byte(cid[0] >> 1)
}

// CSD - Card Specific Data register
type CSD [4]uint32

func (csd CSD) Version() int {
	return int(csd[3]>>30 + 1)
}

//emgo:const
var frac = [16]byte{
	0, 10, 12, 13, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 70, 80,
}

func pow10(exp uint32) int {
	pow := 1
	for exp > 0 {
		pow *= 10
		exp--
	}
	return pow
}

// TAAC returns asynchronous part of the data access time [ns].
func (csd CSD) TAAC() int {
	exp := csd[3] >> 16 & 7
	f := int(frac[csd[3]>>19&15])
	if exp > 0 {
		return f * pow10(exp-1)
	}
	return f / 10
}

// NSAC returns the worst case for the clock-dependent factor of the data
// access time [clk].
func (csd CSD) NSAC() int {
	return int(csd[3] >> 8 & 255 * 100)
}

// TRAN_SPEED returns the maximum data transfer rate per one data line [kb/s].
func (csd CSD) TRAN_SPEED() int {
	exp := csd[3] & 7
	f := int(frac[csd[3]>>3&15])
	return f * 10 * pow10(exp)
}

// CCC returns SDMC command set as 12-bit bitfield.
func (csd CSD) CCC() uint {
	return uint(csd[2] >> 20)
}

// READ_BL_LEN returns the maximum read data block length (bytes).
func (csd CSD) READ_BL_LEN() int {
	return 1 << (csd[2] >> 16 & 15)
}

// READ_BL_PARTIAL reports whether Partial Block Read is allowed.
func (csd CSD) READ_BL_PARTIAL() bool {
	return csd[2]>>15&1 != 0
}

// WRITE_BLK_MISALIGN reports whether the data block to be written by one
// command can be spread over more than one physical block of the memory device.
func (csd CSD) WRITE_BLK_MISALIGN() bool {
	return csd[2]>>14&1 != 0
}

// READ_BLK_MISALIGN reports whether the data block to be read by one command
// can be spread over more than one physical block of the memory device.
func (csd CSD) READ_BLK_MISALIGN() bool {
	return csd[2]>>13&1 != 0
}

// DSR_IMP reports whether the configurable driver stage is integrated on the
// card and Driver Stage Register is implemented.
func (csd CSD) DSR_IMP() bool {
	return csd[2]>>12&1 != 0
}

// C_SIZE returns the user's data card capacity as number of 512 B blocks.
func (csd CSD) C_SIZE() int64 {
	if csd.Version() == 2 {
		return int64(csd[2]&63<<16+csd[1]>>16+1) << 10
	}
	read_bl_len := csd[2] >> 16 & 15
	c_size := csd[2]&0x3FF<<2 + csd[1]>>30
	c_size_mult := csd[1] >> 15 & 7
	return int64((c_size + 1) << (read_bl_len + c_size_mult + 2 - 9))
}

// ERASE_BLK_EN reports whether the memory card supports erasing of one or
// multiple 512 B blocks. If ERASE_BLK_EN is false it supports erasing of one
// or more sectors of SECTOR_SIZE * WRITE_BL_LEN bytes.
func (csd CSD) ERASE_BLK_EN() bool {
	return csd[1]>>14&1 != 0
}

// SECTOR_SIZE returns the size of an erasable sector as number of blocks of
// WRITE_BL_LEN bytes.
func (csd CSD) SECTOR_SIZE() int {
	return int(csd[1]>>7&127 + 1)
}

// WP_GRP_SIZE returns the size of a write protected group.
func (csd CSD) WP_GRP_SIZE() int {
	return int(csd[1]&127 + 1)
}

// WP_GRP_ENABLE reports whether the write protection is possible
func (csd CSD) WP_GRP_ENABLE() bool {
	return csd[0]>>31&1 != 0
}

// R2W_FACTOR returns the typical block program time as a multiple of the read
// access time.
func (csd CSD) R2W_FACTOR() int {
	return int(csd[0]>>26&7 + 1)
}

// WRITE_BL_LEN returns the maximum write data block length (bytes).
func (csd CSD) WRITE_BL_LEN() int {
	return 1 << (csd[0] >> 22 & 15)
}

// WRITE_BL_PARTIAL reports whether the partial block sizes can be used in
// block write commands
func (csd CSD) WRITE_BL_PARTIAL() bool {
	return csd[0]>>21&1 != 0
}

type FileFormat byte

const (
	HardDisk  FileFormat = 0 // Hard disk-like file system with partition table.
	DOSFloppy FileFormat = 1 // DOS FAT (floppy-like) without partition table.
	UFF       FileFormat = 2 // Universal File Format.
	OtherFF   FileFormat = 3 // Other/unknown.
)

// FILE_FORMAT returns the file format on the card (3-bit field: the MS bit
// represents format group, two LS bits represents format).
func (csd CSD) FILE_FORMAT() FileFormat {
	return FileFormat(csd[3]>>13&4 | csd[3]>>10&3)
}

// COPY reports whether the contents has been copied (not original).
func (csd CSD) COPY() bool {
	return csd[3]>>14&1 != 0
}

// PERM_WRITE_PROTECT reports whether the entrie card content is permanently
// protected against overwriting or erasing.
func (csd CSD) PERM_WRITE_PROTECT() bool {
	return csd[3]>>13&1 != 0
}

// TMP_WRITE_PROTECT reports whether the entrie card content is temporarily
// protected against overwriting or erasing.
func (csd CSD) TMP_WRITE_PROTECT() bool {
	return csd[3]>>12&1 != 0
}

// CRC returns the checksum for the CSD content.
func (csd CSD) CRC() byte {
	return byte(csd[3] >> 1 & 127)
}
