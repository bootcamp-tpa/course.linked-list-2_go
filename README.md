# Another linked list implementation in Go

See `main.go` and finish the implementation.

The output should look like:

```
{ value: one, next: 0xc0000a4020 } -> { value: two, next: 0xc0000a4040 } -> { value: three, next: 0xc0000a4060 } -> { value: 4, next: 0xc0000a4080 } -> { value: 5, next: 0xc0000a40a0 } -> { value: 6, next: 0x0 } -> 
---
{ value: 6, next: 0xc0000a4080 } -> { value: 5, next: 0xc0000a4060 } -> { value: 4, next: 0xc0000a4040 } -> { value: three, next: 0xc0000a4020 } -> { value: two, next: 0xc0000a4000 } -> { value: one, next: 0x0 } -> 
---
{ value: one, next: 0xc0000a4020 } -> { value: two, next: 0xc0000a4040 } -> { value: three, next: 0xc0000a4060 } -> { value: 4, next: 0xc0000a4080 } -> { value: 5, next: 0xc0000a40a0 } -> { value: 6, next: 0x0 } -> 
```
