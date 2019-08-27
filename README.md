# FooVM

A Simple Bytecode Runtime

`go get github.com/apkrieg/foovm`

### Description
This runtime was built at part of a contest with [circle601](https://github.com/circle601) to see which runtime could run a set of example applications the fastest.

This runtime was the winner as it was the only runtime that could pass all the tests.

### Update

This project hasn't been updated in a while and I don't think I will update it. I've kinda stepped away from Go because I don't like what Google has been doing with the language. I've started to use Rust more and more as my daily driver and have started development on FizzVM which is written in Rust and Assembly.

### Hello World!
```
push 16        ; string heap index
push 0         ; string heap segment
push 12        ; string length
push 0         ; print call index
push 255       ; runtime call segment
call           ; call print
push 1         ; exit call index
push 255       ; runtime call segment
call           ; call exit
"Hello World!" ; "Hello World!" string
```

### Bytecode Reference
Instruction | Mnemonic | Op Code | Description
----------- | -------- | ------- | -----------
Nil | nil | 0x00 | Nil pointer/no-op
Push | push | 0x01 | *stack* <- data
Pop | pop | 0x02 | <- *stack*
Load | load | 0x03 | Load data from heap to stack
Store | store | 0x04 | store data from stack to heap
Add | add | 0x05 | Add pop1 + pop2
Subtract | sub | 0x06 | Subtract pop1 - pop2
Multiply | mul | 0x07 | Multiply pop1 * pop2
Divide | div | 0x08 | Divide pop1 / pop2
Call | call | 0x09 | Call a function
Return | ret | 0x0a | Return from function
Jump | jmp | 0x0b | Jump to address
Compare | cmp | 0x0c | Compare two values
Jump if Equal | jeq | 0x0d | Jump if a == b
Jump if not Equal | jne | 0x0e | Jump if a != b
Jump if Greater Than | jgt | 0x0f | Jump if a > b
Jump if Less Than | jlt | 0x10 | Jump if a < b
Jump if Less Than or Equal | jle | 0x11 | Jump if a <= b
Jump if Greater Than or Equal | jge | 0x12 | Jump if a >= b
