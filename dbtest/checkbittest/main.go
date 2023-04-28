// ref: github.com/cockroachdb/pebble

package main

import (
	"log"

	hconsts "github.com/ava-labs/hypersdk/consts"
	// "github.com/ava-labs/hypersdk/pebble"
)

func CheckBit(b uint8, v uint8) bool {
	if v > hconsts.MaxUint8Offset { // hconsts.MaxUint8Offset = 7
		return false
	}
	marker := uint8(1 << v)

	return b&marker != 0
}
func main() {
	log.Printf("main")
	log.Printf("CheckBit(255,9): %v", CheckBit(255, 9))     // false
	log.Printf("CheckBit(255,10): %v", CheckBit(255, 10))   // false
	log.Printf("CheckBit(9,255): %v", CheckBit(9, 255))     // false
	log.Printf("CheckBit(9,9): %v", CheckBit(9, 9))         // false
	log.Printf("CheckBit(255,255): %v", CheckBit(255, 255)) // false
	log.Printf("CheckBit(255,5): %v", CheckBit(255, 5))     // true
	log.Printf("CheckBit(4,5): %v", CheckBit(4, 5))         // false
	log.Printf("CheckBit(8,5): %v", CheckBit(8, 5))         // false
	log.Printf("CheckBit(16,5): %v", CheckBit(16, 5))       // false
	log.Printf("CheckBit(32,5): %v", CheckBit(32, 5))       // true
	log.Printf("CheckBit(64,5): %v", CheckBit(64, 5))       // false
	log.Printf("CheckBit(128,5): %v", CheckBit(128, 5))     // false
}
