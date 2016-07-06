# FooVM

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

Instruction | Mnemonic | Op Code | Description
----------- | -------- | ------- | -----------
Nil | nil | 0x00 | Nil pointer/no-op
Push | push | 0x01 | Push data onto stack
Pop | pop | 0x02 | Pop data off of stack
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
