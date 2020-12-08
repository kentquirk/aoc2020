package main

import (
	"errors"
	"fmt"
	"log"
)

func day08a(vm *VM) {
	err := vm.Run()
	if err != nil {
		if errors.Is(err, ErrInfiniteLoop) {
			fmt.Printf("Infinite loop: value of accumulator after %d instructions was %d\n",
				vm.InstructionsExecuted(), vm.Accumulator())
		} else {
			log.Fatalf("Run failed with %s", err)
		}
	} else {
		fmt.Printf("Normal termination: value of accumulator after %d instructions was %d\n",
			vm.InstructionsExecuted(), vm.Accumulator())
	}
}

func day08b(vm *VM) {
	for i := 0; i < vm.Size(); i++ {
		op, _ := vm.InstructionAt(i)
		orig := op.opcode
		switch op.opcode {
		case NOP:
			op.opcode = JMP
		case JMP:
			op.opcode = NOP
		default:
			continue
		}
		vm.Reset()
		if vm.Run() == nil {
			fmt.Printf("Normal termination: value of accumulator after %d instructions was %d\n",
				vm.InstructionsExecuted(), vm.Accumulator())
			break
		}
		op.opcode = orig
	}
}

func main() {
	vm := new(VM)
	err := vm.Load("./input.txt")
	if err != nil {
		log.Fatalf("Load returned %s", err)
	}
	day08a(vm)
	day08b(vm)
}
