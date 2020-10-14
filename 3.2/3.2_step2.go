package main

import (
	"fmt"
	"math"
)

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Printf("%v %T\n", bigIndex, bigIndex) // 15 int32

	// По аналогии выше легко понять как конвертировать в другие типы:
	var a int32 = 22
	var b = uint64(a)
	fmt.Printf("%v %T\n", b, b) // 22 uint64

	var big int64 = 64
	var little = int8(big)
	fmt.Printf("%v %T\n", little, little) // 64 int8

	var big2 int64 = 129
	var little2 = int8(big2)
	fmt.Printf("%v %T\n", little2, little2) // -127 int8

	fmt.Println()

	fmt.Println("MaxInt8: ", int8(math.MaxInt8))
	fmt.Println("MinInt8: ", int8(math.MinInt8))
	fmt.Println("MaxInt16: ", int16(math.MaxInt16))
	fmt.Println("MinInt16: ", int16(math.MinInt16))
	fmt.Println("MaxInt32: ", int32(math.MaxInt32))
	fmt.Println("MinInt32: ", int32(math.MinInt32))
	fmt.Println("MaxInt64: ", int64(math.MaxInt64))
	fmt.Println("MinInt64: ", int64(math.MinInt64))
	fmt.Println("MaxUint8: ", uint8(math.MaxUint8))
	fmt.Println("MaxUint16: ", uint16(math.MaxUint16))
	fmt.Println("MaxUint32: ", uint32(math.MaxUint32))
	fmt.Println("MaxUint64: ", uint64(math.MaxUint64))
}
