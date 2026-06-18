package main

import "fmt"

func main() {
	// Signed integers
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	// Unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	// Floating-point numbers
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	// Complex numbers
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i

	// Print the values
	fmt.Println("Signed integers:")
	fmt.Printf("int8	:%v\n", i8)
	fmt.Printf("int16	:%v\n", i16)
	fmt.Printf("int32	:%v\n", i32)
	fmt.Printf("int64	:%v\n", i64)

	fmt.Println("\nUnsigned integers:")
	fmt.Printf("uint8	:%v\n", u8)
	fmt.Printf("uint16	:%v\n", u16)
	fmt.Printf("uint32	:%v\n", u32)
	fmt.Printf("uint64	:%v\n", u64)

	fmt.Println("\nFloating-point numbers:")
	fmt.Printf("float32	:%v\n", f32)
	fmt.Printf("float64	:%v\n", f64)

	fmt.Println("\nComplex numbers:")
	fmt.Printf("complex64	:%v\n", c64)
	fmt.Printf("complex128	:%v\n", c128)

}