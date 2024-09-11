package vm

import "fmt"

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

const (
	MEMORY_SIZE   = 256
	NUM_REGISTERS = 3
)

type Computer struct {
	memory    *[MEMORY_SIZE]byte
	registers [NUM_REGISTERS]byte
	pc        int
}

func NewComputer() *Computer {
	return &Computer{pc: 8}
}

func (c *Computer) LoadProgram(program *[256]byte) {
	c.memory = program
}

func (c *Computer) Run() {
	for {
		opcode := c.memory[c.pc]

		c.pc++
		register := c.memory[c.pc]

		c.pc++
		operand := c.memory[c.pc]
		c.pc++

		switch opcode {
		case Halt:
			return
		case Load:
			c.registers[register] = c.memory[operand]
		case Store:
			c.memory[operand] = c.registers[register]
		case Add:
			c.registers[register] += c.registers[operand]
		case Sub:
			c.registers[register] -= c.registers[operand]
		default:
			fmt.Printf("Unknown opcode: %d\n", opcode)
			return
		}
	}
}

func (c *Computer) PrintState() {
	fmt.Println("Registers:", c.registers)
	fmt.Println("Memory (first 16 bytes):", c.memory[:16])
}

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
func Compute(memory []byte) {
	vm := NewComputer()
	vm.LoadProgram((*[256]byte)(memory))
	vm.PrintState()
	vm.Run()
	vm.PrintState()
	//registers := [3]byte{8, 0, 0} // PC, R1 and R2
}
