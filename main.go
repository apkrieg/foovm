package main

import (
	fvm "github.com/apkrieg/leedlebyte/foovm"
)

var (
	Program = []byte {
		fvm.Push, 0x12, // 0, 1
		fvm.Push, 0x00, // 2, 3
		fvm.Push, 0x0c, // 4, 5   Length
		fvm.Push, 0x00, // 6, 7
		fvm.Push, 0xff, // 8, 9
		fvm.Call, 0x00, // a, b   Call Print
		fvm.Push, 0x1e, // c, d
		fvm.Push, 0x00, // e, f
		fvm.Jmp,  0x00, // 10, 11 Jump to Exit
		// Hello World!
		'H', 'e',       // 12, 13
		'l', 'l',       // 14, 15
		'o', ' ',       // 16, 17
		'W', 'o',       // 18, 19
		'r', 'l',       // 1a, 1b
		'd', '!',       // 1c, 1d
		// Exit
		fvm.Push, 0x01, // 1e, 1f
		fvm.Push, 0xff, // 20, 21
		fvm.Call, 0x00, // 22, 23 Call Exit
	}
)

func main() {
	vm := fvm.New()
	for i, v := range Program {
		vm.Heap[i] = v
	}
	vm.Exec()
}
