// ref: github.com/cockroachdb/pebble

package main

import (
	"log"

	hconsts "github.com/ava-labs/hypersdk/consts"
	// "github.com/ava-labs/hypersdk/pebble"
)

func CheckBit(b uint8, v uint8) bool {
	if v > hconsts.MaxUint8Offset {
		return false
	}
	marker := uint8(1 << v)
	return b&marker != 0
}
func main() {
	log.Printf("main")
	log.Printf("CheckBit(255,9): %v", CheckBit(255, 9))
	log.Printf("CheckBit(255,10): %v", CheckBit(255, 10))
	log.Printf("CheckBit(9,255): %v", CheckBit(9, 255))
	log.Printf("CheckBit(9,9): %v", CheckBit(9, 9))
	log.Printf("CheckBit(255,255): %v", CheckBit(255, 255))
}
