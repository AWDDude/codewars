# Obfuscated Hello World

The goal of this challenge is make the most difficult to read "Hello World" program.

## How it works

This program uses a carefully crafted slice of integers and performs a specific math operation against a running total. Each integer in the slice has a single math operation applied to it at a time. There are 4 math operation types being performed in order, in a loop (Add, Subtract, Multiply, Divide). Each Math operation has a dedicated [goroutine](https://go.dev/tour/concurrency/1) which uses [channels](https://go.dev/tour/concurrency/2) for input and output. The output channel of the "Add" goroutine is linked to the input of the "Subtract" goroutine, which then links to the "Multiply" goroutine and then to the "Divide" goroutine which links back to the main thread which converts the integer to a character and prints it to the screen. The main thread then links to the "Add" goroutine and the process starts all over.
