package main

import "vm/vm"

func main() {
	memory := make([]byte, 256)
	vm.Compute(memory)
}
