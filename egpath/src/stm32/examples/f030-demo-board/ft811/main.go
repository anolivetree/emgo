package main

import (
	"delay"
	"rtos"

	"display/eve"
	"display/eve/ft81"

	"stm32/evedci"
	"stm32/hal/dma"
	"stm32/hal/exti"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/spi"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var dci *evedci.SPI

func init() {
	system.SetupPLL(8, 1, 48/8)
	systick.Setup(2e6)

	// GPIO

	gpio.A.EnableClock(true)
	spiport, sck, miso, mosi := gpio.A, gpio.Pin5, gpio.Pin6, gpio.Pin7
	irqn := gpio.A.Pin(9)
	pdn := gpio.A.Pin(10)

	gpio.B.EnableClock(true)
	csn := gpio.B.Pin(1)

	// EVE control lines

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.High}
	pdn.Setup(&cfg)
	csn.Setup(&cfg)
	irqn.Setup(&gpio.Config{Mode: gpio.In})
	irqline := exti.Lines(irqn.Mask())
	irqline.Connect(irqn.Port())
	irqline.EnableFallTrig()
	irqline.EnableIRQ()
	rtos.IRQ(irq.EXTI4_15).Enable()

	// EVE SPI

	spiport.Setup(sck|mosi, &gpio.Config{Mode: gpio.Alt, Speed: gpio.High})
	spiport.Setup(miso, &gpio.Config{Mode: gpio.AltIn})
	spiport.SetAltFunc(sck|miso|mosi, gpio.SPI1)
	d := dma.DMA1
	d.EnableClock(true)
	spidrv := spi.NewDriver(spi.SPI1, d.Channel(3, 0), d.Channel(2, 0))
	rtos.IRQ(irq.SPI1).Enable()
	rtos.IRQ(irq.DMA1_Channel2_3).Enable()

	dci = evedci.NewSPI(spidrv, csn, pdn)
	dci.Setup(11e6)
}

func main() {
	dci.SetPDN(0)
	delay.Millisec(20)
	dci.SetPDN(1)
	delay.Millisec(20)

	lcd := eve.NewDriver(dci, 128)
	lcd.HostCmd(ft81.CLKEXT, 0)
	lcd.HostCmd(ft81.ACTIVE, 0)
	delay.Millisec(300)

	dci.Setup(30e6)

	cfg := &eve.Default800x480
	lcd.WriteUint32(ft81.REG_HCYCLE, uint32(cfg.Hcycle))
	lcd.WriteUint32(ft81.REG_HOFFSET, uint32(cfg.Hoffset))
	lcd.WriteUint32(ft81.REG_HSIZE, uint32(cfg.Hsize))
	lcd.WriteUint32(ft81.REG_HSYNC0, uint32(cfg.Hsync0))
	lcd.WriteUint32(ft81.REG_HSYNC1, uint32(cfg.Hsync1))
	lcd.WriteUint32(ft81.REG_VCYCLE, uint32(cfg.Vcycle))
	lcd.WriteUint32(ft81.REG_VOFFSET, uint32(cfg.Voffset))
	lcd.WriteUint32(ft81.REG_VSIZE, uint32(cfg.Vsize))
	lcd.WriteUint32(ft81.REG_VSYNC0, uint32(cfg.Vsync0))
	lcd.WriteUint32(ft81.REG_VSYNC1, uint32(cfg.Vsync1))
	lcd.WriteUint32(ft81.REG_PCLK_POL, uint32(cfg.ClkPol))

	lcd.WriteUint32(ft81.REG_GPIO, 0x80)
	lcd.WriteUint32(ft81.REG_PCLK, 5)

	var x, y int16
	for {
		lcd.WriteUint32(ft81.RAM_DL, eve.CLEAR|eve.CST)
		lcd.WriteUint32(ft81.RAM_DL+4, eve.BEGIN|eve.POINTS)
		lcd.WriteUint32(ft81.RAM_DL+8, eve.POINT_SIZE|150*16)
		lcd.WriteUint32(ft81.RAM_DL+12, eve.VERTEX2F|400*16<<15|240*16)
		lcd.WriteUint32(ft81.RAM_DL+16, eve.POINT_SIZE|100*16)
		lcd.WriteUint32(ft81.RAM_DL+20, eve.COLOR_RGB|0x8800AA)
		lcd.WriteUint32(ft81.RAM_DL+24, eve.COLOR_A|128)
		lcd.WriteUint32(
			ft81.RAM_DL+28,
			eve.VERTEX2F|uint32(x)&0x7FFF*16<<15|uint32(y)&0x7FFF*16,
		)
		lcd.WriteUint32(ft81.RAM_DL+32, eve.DISPLAY)
		
		lcd.WriteUint32(ft81.REG_DLSWAP, eve.DLSWAP_FRAME)

		for {
			xy := lcd.ReadUint32(ft81.REG_TOUCH_SCREEN_XY)
			x, y = int16(xy>>16), int16(xy)
			if x >= 0 && y >= 0 {
				break
			}
		}
	}
}

func lcdSPIISR() {
	dci.SPI().ISR()
}

func lcdDMAISR() {
	dci.SPI().DMAISR(dci.SPI().RxDMA())
	dci.SPI().DMAISR(dci.SPI().TxDMA())
}

func exti4_15ISR() {
	pending := exti.Pending()
	pending &= exti.L4 | exti.L5 | exti.L6 | exti.L7 | exti.L8 | exti.L9 |
		exti.L10 | exti.L11 | exti.L12 | exti.L13 | exti.L14 | exti.L15
	pending.ClearPending()
	if pending&exti.L9 != 0 {
		dci.ISR()
	}
}

//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.SPI1:            lcdSPIISR,
	irq.DMA1_Channel2_3: lcdDMAISR,
	irq.EXTI4_15:        exti4_15ISR,
}
