// +build f411xe

// Peripheral: ADC_Periph  Analog to Digital Converter.
// Instances:
//  ADC1  mmap.ADC1_BASE
// Registers:
//  0x00 32  SR     Status register.
//  0x04 32  CR1    Control register 1.
//  0x08 32  CR2    Control register 2.
//  0x0C 32  SMPR1  Sample time register 1.
//  0x10 32  SMPR2  Sample time register 2.
//  0x14 32  JOFR1  Injected channel data offset register 1.
//  0x18 32  JOFR2  Injected channel data offset register 2.
//  0x1C 32  JOFR3  Injected channel data offset register 3.
//  0x20 32  JOFR4  Injected channel data offset register 4.
//  0x24 32  HTR    Watchdog higher threshold register.
//  0x28 32  LTR    Watchdog lower threshold register.
//  0x2C 32  SQR1   Regular sequence register 1.
//  0x30 32  SQR2   Regular sequence register 2.
//  0x34 32  SQR3   Regular sequence register 3.
//  0x38 32  JSQR   Injected sequence register.
//  0x3C 32  JDR[4] Injected data registers.
//  0x4C 32  DR     Regular data register.
// Import:
//  stm32/o/f411xe/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	AWD   SR = 0x01 << 0 //+ Analog watchdog flag.
	EOC   SR = 0x01 << 1 //+ End of conversion.
	JEOC  SR = 0x01 << 2 //+ Injected channel end of conversion.
	JSTRT SR = 0x01 << 3 //+ Injected channel Start flag.
	STRT  SR = 0x01 << 4 //+ Regular channel Start flag.
	OVR   SR = 0x01 << 5 //+ Overrun flag.
)

const (
	AWDn   = 0
	EOCn   = 1
	JEOCn  = 2
	JSTRTn = 3
	STRTn  = 4
	OVRn   = 5
)

const (
	AWDCH   CR1 = 0x1F << 0  //+ AWDCH[4:0] bits (Analog watchdog channel select bits).
	EOCIE   CR1 = 0x01 << 5  //+ Interrupt enable for EOC.
	AWDIE   CR1 = 0x01 << 6  //+ AAnalog Watchdog interrupt enable.
	JEOCIE  CR1 = 0x01 << 7  //+ Interrupt enable for injected channels.
	SCAN    CR1 = 0x01 << 8  //+ Scan mode.
	AWDSGL  CR1 = 0x01 << 9  //+ Enable the watchdog on a single channel in scan mode.
	JAUTO   CR1 = 0x01 << 10 //+ Automatic injected group conversion.
	DISCEN  CR1 = 0x01 << 11 //+ Discontinuous mode on regular channels.
	JDISCEN CR1 = 0x01 << 12 //+ Discontinuous mode on injected channels.
	DISCNUM CR1 = 0x07 << 13 //+ DISCNUM[2:0] bits (Discontinuous mode channel count).
	JAWDEN  CR1 = 0x01 << 22 //+ Analog watchdog enable on injected channels.
	AWDEN   CR1 = 0x01 << 23 //+ Analog watchdog enable on regular channels.
	RES     CR1 = 0x03 << 24 //+ RES[2:0] bits (Resolution).
	OVRIE   CR1 = 0x01 << 26 //+ overrun interrupt enable.
)

const (
	AWDCHn   = 0
	EOCIEn   = 5
	AWDIEn   = 6
	JEOCIEn  = 7
	SCANn    = 8
	AWDSGLn  = 9
	JAUTOn   = 10
	DISCENn  = 11
	JDISCENn = 12
	DISCNUMn = 13
	JAWDENn  = 22
	AWDENn   = 23
	RESn     = 24
	OVRIEn   = 26
)

const (
	ADON     CR2 = 0x01 << 0  //+ A/D Converter ON / OFF.
	CONT     CR2 = 0x01 << 1  //+ Continuous Conversion.
	DMA      CR2 = 0x01 << 8  //+ Direct Memory access mode.
	DDS      CR2 = 0x01 << 9  //+ DMA disable selection (Single ADC).
	EOCS     CR2 = 0x01 << 10 //+ End of conversion selection.
	ALIGN    CR2 = 0x01 << 11 //+ Data Alignment.
	JEXTSEL  CR2 = 0x0F << 16 //+ JEXTSEL[3:0] bits (External event select for injected group).
	JEXTEN   CR2 = 0x03 << 20 //+ JEXTEN[1:0] bits (External Trigger Conversion mode for injected channelsp).
	JSWSTART CR2 = 0x01 << 22 //+ Start Conversion of injected channels.
	EXTSEL   CR2 = 0x0F << 24 //+ EXTSEL[3:0] bits (External Event Select for regular group).
	EXTEN    CR2 = 0x03 << 28 //+ EXTEN[1:0] bits (External Trigger Conversion mode for regular channelsp).
	SWSTART  CR2 = 0x01 << 30 //+ Start Conversion of regular channels.
)

const (
	ADONn     = 0
	CONTn     = 1
	DMAn      = 8
	DDSn      = 9
	EOCSn     = 10
	ALIGNn    = 11
	JEXTSELn  = 16
	JEXTENn   = 20
	JSWSTARTn = 22
	EXTSELn   = 24
	EXTENn    = 28
	SWSTARTn  = 30
)

const (
	SMP10 SMPR1 = 0x07 << 0  //+ SMP10[2:0] bits (Channel 10 Sample time selection).
	SMP11 SMPR1 = 0x07 << 3  //+ SMP11[2:0] bits (Channel 11 Sample time selection).
	SMP12 SMPR1 = 0x07 << 6  //+ SMP12[2:0] bits (Channel 12 Sample time selection).
	SMP13 SMPR1 = 0x07 << 9  //+ SMP13[2:0] bits (Channel 13 Sample time selection).
	SMP14 SMPR1 = 0x07 << 12 //+ SMP14[2:0] bits (Channel 14 Sample time selection).
	SMP15 SMPR1 = 0x07 << 15 //+ SMP15[2:0] bits (Channel 15 Sample time selection).
	SMP16 SMPR1 = 0x07 << 18 //+ SMP16[2:0] bits (Channel 16 Sample time selection).
	SMP17 SMPR1 = 0x07 << 21 //+ SMP17[2:0] bits (Channel 17 Sample time selection).
	SMP18 SMPR1 = 0x07 << 24 //+ SMP18[2:0] bits (Channel 18 Sample time selection).
)

const (
	SMP10n = 0
	SMP11n = 3
	SMP12n = 6
	SMP13n = 9
	SMP14n = 12
	SMP15n = 15
	SMP16n = 18
	SMP17n = 21
	SMP18n = 24
)

const (
	SMP0 SMPR2 = 0x07 << 0  //+ SMP0[2:0] bits (Channel 0 Sample time selection).
	SMP1 SMPR2 = 0x07 << 3  //+ SMP1[2:0] bits (Channel 1 Sample time selection).
	SMP2 SMPR2 = 0x07 << 6  //+ SMP2[2:0] bits (Channel 2 Sample time selection).
	SMP3 SMPR2 = 0x07 << 9  //+ SMP3[2:0] bits (Channel 3 Sample time selection).
	SMP4 SMPR2 = 0x07 << 12 //+ SMP4[2:0] bits (Channel 4 Sample time selection).
	SMP5 SMPR2 = 0x07 << 15 //+ SMP5[2:0] bits (Channel 5 Sample time selection).
	SMP6 SMPR2 = 0x07 << 18 //+ SMP6[2:0] bits (Channel 6 Sample time selection).
	SMP7 SMPR2 = 0x07 << 21 //+ SMP7[2:0] bits (Channel 7 Sample time selection).
	SMP8 SMPR2 = 0x07 << 24 //+ SMP8[2:0] bits (Channel 8 Sample time selection).
	SMP9 SMPR2 = 0x07 << 27 //+ SMP9[2:0] bits (Channel 9 Sample time selection).
)

const (
	SMP0n = 0
	SMP1n = 3
	SMP2n = 6
	SMP3n = 9
	SMP4n = 12
	SMP5n = 15
	SMP6n = 18
	SMP7n = 21
	SMP8n = 24
	SMP9n = 27
)

const (
	JOFFSET1 JOFR1 = 0xFFF << 0 //+ Data offset for injected channel 1.
)

const (
	JOFFSET1n = 0
)

const (
	JOFFSET2 JOFR2 = 0xFFF << 0 //+ Data offset for injected channel 2.
)

const (
	JOFFSET2n = 0
)

const (
	JOFFSET3 JOFR3 = 0xFFF << 0 //+ Data offset for injected channel 3.
)

const (
	JOFFSET3n = 0
)

const (
	JOFFSET4 JOFR4 = 0xFFF << 0 //+ Data offset for injected channel 4.
)

const (
	JOFFSET4n = 0
)

const (
	HT HTR = 0xFFF << 0 //+ Analog watchdog high threshold.
)

const (
	HTn = 0
)

const (
	LT LTR = 0xFFF << 0 //+ Analog watchdog low threshold.
)

const (
	LTn = 0
)

const (
	SQ13 SQR1 = 0x1F << 0  //+ SQ13[4:0] bits (13th conversion in regular sequence).
	SQ14 SQR1 = 0x1F << 5  //+ SQ14[4:0] bits (14th conversion in regular sequence).
	SQ15 SQR1 = 0x1F << 10 //+ SQ15[4:0] bits (15th conversion in regular sequence).
	SQ16 SQR1 = 0x1F << 15 //+ SQ16[4:0] bits (16th conversion in regular sequence).
	L    SQR1 = 0x0F << 20 //+ L[3:0] bits (Regular channel sequence length).
)

const (
	SQ13n = 0
	SQ14n = 5
	SQ15n = 10
	SQ16n = 15
	Ln    = 20
)

const (
	SQ7  SQR2 = 0x1F << 0  //+ SQ7[4:0] bits (7th conversion in regular sequence).
	SQ8  SQR2 = 0x1F << 5  //+ SQ8[4:0] bits (8th conversion in regular sequence).
	SQ9  SQR2 = 0x1F << 10 //+ SQ9[4:0] bits (9th conversion in regular sequence).
	SQ10 SQR2 = 0x1F << 15 //+ SQ10[4:0] bits (10th conversion in regular sequence).
	SQ11 SQR2 = 0x1F << 20 //+ SQ11[4:0] bits (11th conversion in regular sequence).
	SQ12 SQR2 = 0x1F << 25 //+ SQ12[4:0] bits (12th conversion in regular sequence).
)

const (
	SQ7n  = 0
	SQ8n  = 5
	SQ9n  = 10
	SQ10n = 15
	SQ11n = 20
	SQ12n = 25
)

const (
	SQ1 SQR3 = 0x1F << 0  //+ SQ1[4:0] bits (1st conversion in regular sequence).
	SQ2 SQR3 = 0x1F << 5  //+ SQ2[4:0] bits (2nd conversion in regular sequence).
	SQ3 SQR3 = 0x1F << 10 //+ SQ3[4:0] bits (3rd conversion in regular sequence).
	SQ4 SQR3 = 0x1F << 15 //+ SQ4[4:0] bits (4th conversion in regular sequence).
	SQ5 SQR3 = 0x1F << 20 //+ SQ5[4:0] bits (5th conversion in regular sequence).
	SQ6 SQR3 = 0x1F << 25 //+ SQ6[4:0] bits (6th conversion in regular sequence).
)

const (
	SQ1n = 0
	SQ2n = 5
	SQ3n = 10
	SQ4n = 15
	SQ5n = 20
	SQ6n = 25
)

const (
	JSQ1 JSQR = 0x1F << 0  //+ JSQ1[4:0] bits (1st conversion in injected sequence).
	JSQ2 JSQR = 0x1F << 5  //+ JSQ2[4:0] bits (2nd conversion in injected sequence).
	JSQ3 JSQR = 0x1F << 10 //+ JSQ3[4:0] bits (3rd conversion in injected sequence).
	JSQ4 JSQR = 0x1F << 15 //+ JSQ4[4:0] bits (4th conversion in injected sequence).
	JL   JSQR = 0x03 << 20 //+ JL[1:0] bits (Injected Sequence length).
)

const (
	JSQ1n = 0
	JSQ2n = 5
	JSQ3n = 10
	JSQ4n = 15
	JLn   = 20
)

const (
	JDATA JDR = 0xFFFF << 0 //+ Injected data.
)

const (
	JDATAn = 0
)

const (
	DATA     DR = 0xFFFF << 0  //+ Regular data.
	ADC2DATA DR = 0xFFFF << 16 //+ ADC2 data.
)

const (
	DATAn     = 0
	ADC2DATAn = 16
)
