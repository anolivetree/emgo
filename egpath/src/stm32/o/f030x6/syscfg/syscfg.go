// Peripheral: SYSCFG_Periph  SysTem Configuration.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  CFGR1     Configuration register 1.
//  0x08 32  EXTICR[4] External interrupt configuration register.
//  0x18 32  CFGR2     Configuration register 2.
// Import:
//  stm32/o/f030x6/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_MODE         CFGR1_Bits = 0x03 << 0  //+ SYSCFG_Memory Remap Config.
	DMA_RMP          CFGR1_Bits = 0x1F << 8  //+ DMA remap mask.
	ADC_DMA_RMP      CFGR1_Bits = 0x01 << 8  //  ADC DMA remap.
	USART1TX_DMA_RMP CFGR1_Bits = 0x02 << 8  //  USART1 TX DMA remap.
	USART1RX_DMA_RMP CFGR1_Bits = 0x04 << 8  //  USART1 RX DMA remap.
	TIM16_DMA_RMP    CFGR1_Bits = 0x08 << 8  //  Timer 16 DMA remap.
	TIM17_DMA_RMP    CFGR1_Bits = 0x10 << 8  //  Timer 17 DMA remap.
	I2C_FMP_PB6      CFGR1_Bits = 0x01 << 16 //+ I2C PB6 Fast mode plus.
	I2C_FMP_PB7      CFGR1_Bits = 0x01 << 17 //+ I2C PB7 Fast mode plus.
	I2C_FMP_PB8      CFGR1_Bits = 0x01 << 18 //+ I2C PB8 Fast mode plus.
	I2C_FMP_PB9      CFGR1_Bits = 0x01 << 19 //+ I2C PB9 Fast mode plus.
	I2C_FMP_I2C1     CFGR1_Bits = 0x01 << 20 //+ Enable Fast Mode Plus on PB10, PB11, PF6 and PF7.
	I2C_FMP_PA9      CFGR1_Bits = 0x01 << 22 //+ Enable Fast Mode Plus on PA9.
	I2C_FMP_PA10     CFGR1_Bits = 0x01 << 23 //+ Enable Fast Mode Plus on PA10.
)

const (
	MEM_MODEn     = 0
	DMA_RMPn      = 8
	I2C_FMP_PB6n  = 16
	I2C_FMP_PB7n  = 17
	I2C_FMP_PB8n  = 18
	I2C_FMP_PB9n  = 19
	I2C_FMP_I2C1n = 20
	I2C_FMP_PA9n  = 22
	I2C_FMP_PA10n = 23
)

const (
	EXTI0    EXTICR_Bits = 0x0F << 0  //+ EXTI 0 configuration.
	EXTI1    EXTICR_Bits = 0x0F << 4  //+ EXTI 1 configuration.
	EXTI2    EXTICR_Bits = 0x0F << 8  //+ EXTI 2 configuration.
	EXTI3    EXTICR_Bits = 0x0F << 12 //+ EXTI 3 configuration.
	EXTI0_PA EXTICR_Bits = 0x00 << 12 //  PA[0] pin.
	EXTI0_PB EXTICR_Bits = 0x01 << 0  //  PB[0] pin.
	EXTI0_PC EXTICR_Bits = 0x02 << 0  //  PC[0] pin.
	EXTI0_PD EXTICR_Bits = 0x03 << 0  //  PD[0] pin.
	EXTI0_PF EXTICR_Bits = 0x05 << 0  //  PF[0] pin.
	EXTI1_PA EXTICR_Bits = 0x00 << 12 //  PA[1] pin.
	EXTI1_PB EXTICR_Bits = 0x01 << 4  //  PB[1] pin.
	EXTI1_PC EXTICR_Bits = 0x02 << 4  //  PC[1] pin.
	EXTI1_PD EXTICR_Bits = 0x03 << 4  //  PD[1] pin.
	EXTI1_PF EXTICR_Bits = 0x05 << 4  //  PF[1] pin.
	EXTI2_PA EXTICR_Bits = 0x00 << 12 //  PA[2] pin.
	EXTI2_PB EXTICR_Bits = 0x01 << 8  //  PB[2] pin.
	EXTI2_PC EXTICR_Bits = 0x02 << 8  //  PC[2] pin.
	EXTI2_PD EXTICR_Bits = 0x03 << 8  //  PD[2] pin.
	EXTI2_PF EXTICR_Bits = 0x05 << 8  //  PF[2] pin.
	EXTI3_PA EXTICR_Bits = 0x00 << 12 //  PA[3] pin.
	EXTI3_PB EXTICR_Bits = 0x01 << 12 //  PB[3] pin.
	EXTI3_PC EXTICR_Bits = 0x02 << 12 //  PC[3] pin.
	EXTI3_PD EXTICR_Bits = 0x03 << 12 //  PD[3] pin.
	EXTI3_PF EXTICR_Bits = 0x05 << 12 //  PF[3] pin.
)

const (
	EXTI0n = 0
	EXTI1n = 4
	EXTI2n = 8
	EXTI3n = 12
)

const (
	LOCKUP_LOCK      CFGR2_Bits = 0x01 << 0 //+ Enables and locks the LOCKUP (Hardfault) output of CortexM0 with Break Input of TIMER1.
	SRAM_PARITY_LOCK CFGR2_Bits = 0x01 << 1 //+ Enables and locks the SRAM_PARITY error signal with Break Input of TIMER1.
	SRAM_PEF         CFGR2_Bits = 0x01 << 8 //+ SRAM Parity error flag.
)

const (
	LOCKUP_LOCKn      = 0
	SRAM_PARITY_LOCKn = 1
	SRAM_PEFn         = 8
)
