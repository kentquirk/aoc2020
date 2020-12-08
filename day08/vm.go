package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Opcode is a type to represent the individual opcodes
type Opcode string

// These are the Opcodes that the VM knows about
const (
	ACC Opcode = "acc"
	JMP Opcode = "jmp"
	NOP Opcode = "nop"
)

// ErrInfiniteLoop is returned when the VM detects an infinite loop
var ErrInfiniteLoop = errors.New("infinite loop")

// ErrAccessOutOfBounds is returned when the VM detects an access beyond the memory bounds
var ErrAccessOutOfBounds = errors.New("access out of bounds")

// ErrBadInstruction is returned when the VM gets an instruction it doesn't understand
var ErrBadInstruction = errors.New("bad opcode")

// Instruction represents a single instruction for our VM
type Instruction struct {
	opcode   Opcode
	operand  int
	runcount int
}

// VM is the entire VM object
type VM struct {
	ip                   int
	accumulator          int
	memory               []Instruction
	instructionsExecuted int
}

func parseInstruction(s string) (Instruction, error) {
	splits := strings.Split(s, " ")
	operand, err := strconv.Atoi(splits[1])
	return Instruction{
		opcode:  Opcode(splits[0]),
		operand: operand,
	}, err
}

// Accumulator is an accessor for the value of the accumulator
func (vm VM) Accumulator() int {
	return vm.accumulator
}

// InstructionsExecuted returns the total number of instructions the VM executed
func (vm VM) InstructionsExecuted() int {
	return vm.instructionsExecuted
}

// Size returns the number of instructions currently loaded into the VM
func (vm VM) Size() int {
	if vm.memory == nil {
		return 0
	}
	return len(vm.memory)
}

// Reset sets the VM to its starting state but does not modify the code
// that it is running.
func (vm *VM) Reset() {
	vm.ip = 0
	vm.accumulator = 0
	vm.instructionsExecuted = 0
	for i := range vm.memory {
		vm.memory[i].runcount = 0
	}
}

// InstructionAt returns a pointer to the instruction at the given offset
// This will allow it to be modified.
func (vm VM) InstructionAt(index int) (*Instruction, error) {
	if index < 0 || index >= len(vm.memory) {
		return nil, ErrAccessOutOfBounds
	}
	return &(vm.memory[index]), nil
}

// Load takes a file of instructions, parses it, and loads it into a VM
func (vm *VM) Load(filename string) error {
	// Initialize all the parts
	vm.memory = []Instruction{}

	// load and parse the file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		ins, err := parseInstruction(line)
		if err != nil {
			return err
		}
		vm.memory = append(vm.memory, ins)
	}
	vm.Reset()
	return nil
}

// Run executes the VM starting from the current IP
// If an error occurs, the IP is pointing at the instruction with the error
func (vm *VM) Run() error {
	for {
		// when we fall off the end of memory, that's a normal termination
		if vm.ip >= len(vm.memory) {
			return nil
		}
		if vm.ip < 0 {
			return ErrAccessOutOfBounds
		}
		vm.instructionsExecuted++
		vm.memory[vm.ip].runcount++
		if vm.memory[vm.ip].runcount > 1 {
			return ErrInfiniteLoop
		}
		switch vm.memory[vm.ip].opcode {
		case ACC:
			vm.accumulator += vm.memory[vm.ip].operand
			vm.ip++
		case JMP:
			vm.ip += vm.memory[vm.ip].operand
		case NOP:
			vm.ip++
		default:
			return ErrBadInstruction
		}
	}
	// return nil
}
