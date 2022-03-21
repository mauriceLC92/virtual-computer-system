// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// Opcode is the numeric code which represents an instruction
type Opcode uint64

const (
	HALT Opcode = iota
	NOOP
)

type Gmachine struct {
	// P is the program counter and holds the memory address of the next instruction
	// to execute.
	P uint64
	// A is the accumulator which will keep track of our current result
	A      uint64
	Memory []uint64
}

// New instantiates a new Gmachine in it's initial state.
func New() *Gmachine {
	mem := make([]uint64, DefaultMemSize)
	return &Gmachine{
		Memory: mem,
	}
}

func (gm *Gmachine) Run() {
	for {
		nextInstuction := gm.Memory[gm.P]
		gm.P += 1
		switch nextInstuction {
		case uint64(HALT):
			return
		case uint64(NOOP):
		}
	}
}
