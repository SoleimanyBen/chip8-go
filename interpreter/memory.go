package interpreter

const (
	ReservedAddressSpaceStart = 0x000 // Start of Chip-8 RAM
	ReservedAddressSpaceEnd   = 0x200 // Start of most Chip-8 programs
	ProgramAddressSpaceEnd    = 0x600 // Start of ETI-600 Chip-8 programs
	Chip8AddressSpaceEnd      = 0xFFF // End of Chip-8 RAM
)

type RAM struct {
	mem [Chip8AddressSpaceEnd]byte
}
