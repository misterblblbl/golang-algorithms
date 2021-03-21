package main

import (
	"fmt"
	"unsafe"
)

type t struct {
	b   uint8   // 1 byte
	i   int     // 8 bytes, platform dependent 64-bits wide on a 64-bit system
	pi  *int    // 8 bytes, platform dependent 64-bits wide on a 64-bit system
	i32 int32   // 4 bytes
	i64 int64   // 8 bytes
	f64 float64 // 8 bytes

	r   rune      // 4 bytes
	s   string    // 16 bytes
	ss  []string  // 24 bytes
	ss0 [0]string // 0 bytes
	ss1 [1]string // 16 bytes
}

func main() {
	fmt.Println("PEW")

	data := t{}
	ptr := &data.ss

	fmt.Println("sizeof(uint8)", unsafe.Sizeof(data.b), "bytes")
	fmt.Println("sizeof(int)", unsafe.Sizeof(data.i), "bytes")
	fmt.Println("sizeof(*int)", unsafe.Sizeof(data.pi), "bytes")

	fmt.Println("sizeof(int32)", unsafe.Sizeof(data.i32), "bytes")
	fmt.Println("sizeof(int64)", unsafe.Sizeof(data.i64), "bytes")
	fmt.Println("sizeof(float64)", unsafe.Sizeof(data.f64), "bytes")

	fmt.Println("sizeof(rune)", unsafe.Sizeof(data.r), "bytes")
	fmt.Println("sizeof(string)", unsafe.Sizeof(data.s), "bytes")

	fmt.Println("sizeof([]string)", unsafe.Sizeof(data.ss), "bytes")
	fmt.Println("sizeof([0]string)", unsafe.Sizeof(data.ss0), "bytes")
	fmt.Println("sizeof([1]string)", unsafe.Sizeof(data.ss1), "bytes")
	fmt.Println("sizeof(*[]string)", unsafe.Sizeof(ptr), "bytes")

	fmt.Println("sizeof(t)", unsafe.Sizeof(data), "bytes")
	fmt.Println("sizeof([4]t)", unsafe.Sizeof([4]t{}), "bytes")
}
