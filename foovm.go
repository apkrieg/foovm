package foovm

import (
	"fmt"
)

// Instructions
const (
	Nil   byte = iota // 0x00
	Push  byte = iota // 0x01
	Pop   byte = iota // 0x02
	Load  byte = iota // 0x03
	Store byte = iota // 0x04
	Add   byte = iota // 0x05
	Sub   byte = iota // 0x06
	Mul   byte = iota // 0x07
	Div   byte = iota // 0x08
	Call  byte = iota // 0x09
	Ret   byte = iota // 0x0a
	Jmp   byte = iota // 0x0b
	Cmp   byte = iota // 0x0c
	Jeq   byte = iota // 0x0d
	Jne   byte = iota // 0x0e
	Jgt   byte = iota // 0x0f
	Jlt   byte = iota // 0x10
	Jle   byte = iota // 0x11
	Jge   byte = iota // 0x12
)

type FooVM struct {
	Cmp1 byte
	Cmp2 byte
	// 256 B Stack
	RSP   uint16
	Stack []byte
	// 64 KB Heap
	PHP  uint16
	RHP  uint16
	Heap []byte
}

func New() *FooVM {
	fvm := new(FooVM)
	fvm.Stack = make([]byte, 256)
	fvm.Heap = make([]byte, 256*256)
	return fvm
}

func PrintDebug(fvm *FooVM) {
	fmt.Println("\n\nDebug:")
	fmt.Printf("RSP=%#x, RHP=%#x, PHP=%#x\n", fvm.RSP, fvm.RHP, fvm.PHP)
	fmt.Println("Stack:")
	for i, v := range fvm.Stack {
		if i%16 == 0 && i != 0 {
			fmt.Println("")
		}
		fmt.Printf("%#x ", v)
	}
	fmt.Println("\n\nHeap:")
	for i, v := range fvm.Heap {
		if i%32 == 0 && i != 0 {
			fmt.Println("")
		}
		if i%256 == 0 && i != 0 {
			fmt.Println("-----")
		}
		if i == 5*256 {
			break
		}
		fmt.Printf("%#x ", v)
	}
	return
}

// Execute bytecode
func (fvm *FooVM) Exec() {
	for {
		switch inst := fvm.Heap[fvm.RHP]; {
		case inst == Nil:
			fvm.RHP++
			break
		case inst == Push:
			fvm.Stack[fvm.RSP] = fvm.Heap[fvm.RHP+1]
			fvm.RSP++
			fvm.RHP += 2
			break
		case inst == Pop:
			fvm.RSP--
			fvm.RHP--
			break
		case inst == Load:
			fvm.Stack[fvm.RSP] = fvm.Heap[uint16(fvm.Stack[fvm.RSP-1])*256+uint16(fvm.Stack[fvm.RSP-2])]
			fvm.RSP -= 2
			fvm.RHP++
			break
		case inst == Store:
			fvm.Heap[uint16(fvm.Stack[fvm.RSP-1])*256+uint16(fvm.Stack[fvm.RSP-2])] = fvm.Stack[fvm.RSP-3]
			fvm.RSP -= 3
			fvm.RHP++
			break
		case inst == Add:
			fvm.Stack[fvm.RSP-2] = fvm.Stack[fvm.RSP-1] + fvm.Stack[fvm.RSP-2]
			fvm.RSP--
			fvm.RHP++
			break
		case inst == Sub:
			fvm.Stack[fvm.RSP-2] = fvm.Stack[fvm.RSP-1] - fvm.Stack[fvm.RSP-2]
			fvm.RSP--
			fvm.RHP++
			break
		case inst == Mul:
			fvm.Stack[fvm.RSP-2] = fvm.Stack[fvm.RSP-1] * fvm.Stack[fvm.RSP-2]
			fvm.RSP--
			fvm.RHP++
			break
		case inst == Div:
			fvm.Stack[fvm.RSP-2] = fvm.Stack[fvm.RSP-1] / fvm.Stack[fvm.RSP-2]
			fvm.RSP--
			fvm.RHP++
			break
		case inst == Call:
			fvm.PHP = fvm.RHP
			if fvm.Stack[fvm.RSP-1] == 0xff {
				if fvm.Stack[fvm.RSP-2] == 0x00 {
					temp := uint16(fvm.Stack[fvm.RSP-4])*256 + uint16(fvm.Stack[fvm.RSP-5])
					for _, v := range fvm.Heap[temp : temp+uint16(fvm.Stack[fvm.RSP-3])] {
						fmt.Printf("%c", v)
					}
					fvm.RSP -= 5
					fvm.RHP++
				} else if fvm.Stack[fvm.RSP-2] == 0x01 {
					return
				}
			} else {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP-1])*256 + uint16(fvm.Stack[fvm.RSP-2])
				fvm.RSP -= 2
			}
			break
		case inst == Ret:
			fvm.RHP = fvm.PHP
			fvm.RHP++
			break
		case inst == Jmp:
			fvm.RHP = uint16(fvm.Stack[fvm.RSP-1])*256 + uint16(fvm.Stack[fvm.RSP-2])
			fvm.RSP -= 2
			break
		case inst == Cmp:
			fvm.Cmp1 = fvm.Stack[fvm.RSP-1]
			fvm.Cmp2 = fvm.Stack[fvm.RSP-2]
			fvm.RSP -= 2
			fvm.RHP++
			break
		case inst == Jeq:
			if fvm.Cmp1 == fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		case inst == Jne:
			if fvm.Cmp1 != fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		case inst == Jgt:
			if fvm.Cmp1 > fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		case inst == Jlt:
			if fvm.Cmp1 < fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		case inst == Jle:
			if fvm.Cmp1 <= fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		case inst == Jge:
			if fvm.Cmp1 >= fvm.Cmp2 {
				fvm.RHP = uint16(fvm.Stack[fvm.RSP]-1)*256 + uint16(fvm.Stack[fvm.RSP-2])
			}
			fvm.RSP -= 2
			break
		default:
			return
		}
	}
}
