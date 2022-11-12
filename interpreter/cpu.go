package interpreter

type CPU struct {
	memory [0xFFF]uint8

	registers [0xF]uint8 // general purpose registers
	i         uint16     // 16-bit register
	delay     uint8
	sound     uint8

	oc    uint8
	sp    uint8
	pc    uint16
	stack [0x0F]uint16
}
