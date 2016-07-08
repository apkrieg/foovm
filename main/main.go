package main

import (
	fvm "github.com/apkrieg/foovm"
)

var (
	Program = []byte{
		// fvm.Push, 16,
		// fvm.Push, 0,
		// fvm.Push, 12,
		// fvm.Push, 0,
		// fvm.Push, 255,
		// fvm.Call,
		// fvm.Push, 1,
		// fvm.Push, 255,
		// fvm.Call,
		// 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!',
		//
		fvm.Push, 10,   // 0, 1
		fvm.Push, 0,   // 2, 3
		fvm.Call,      // 4
		fvm.Push, 1,   // 5, 6
		fvm.Push, 255, // 7, 8
		fvm.Call,      // 9
		fvm.Push, 22,   // 10, 11
		fvm.Push, 0,   // 12, 13
		fvm.Push, 1,   // 14, 15
		fvm.Push, 0,   // 16, 17
		fvm.Push, 255, // 18, 19
		fvm.Call,      // 20
		fvm.Ret,       // 21
		'A',           // 22
	}
)

func main() {
	vm := fvm.New()
	for i, v := range Program {
		vm.Heap[i] = v
	}
	vm.Exec()
}
