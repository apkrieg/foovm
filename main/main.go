package main

import (
	fvm "github.com/apkrieg/foovm"
)

var (
	Program = []byte{
		fvm.Push, 16,
		fvm.Push, 0,
		fvm.Push, 12,
		fvm.Push, 0,
		fvm.Push, 255,
		fvm.Call,
		fvm.Push, 1,
		fvm.Push, 255,
		fvm.Call,
		'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!',
	}
)

func main() {
	vm := fvm.New()
	for i, v := range Program {
		vm.Heap[i] = v
	}
	vm.Exec()
}
