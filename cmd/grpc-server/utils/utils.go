package utils

import (
	"fmt"
	"os"
	"time"
)

func GetHowOldInSec() int64 {
	i, _ := os.Stat(os.Args[0])

	btime := i.ModTime().Unix()
	return time.Now().Unix() - btime
}

// get current time ( unix )
func CurrentUnixTime() int64 {
	return time.Now().Unix()
}

// get build unix time
func GetBuildUnixTime() int64 {
	i, _ := os.Stat(os.Args[0])
	btime := i.ModTime().Unix()
	return btime
}

// get build time ( string )
func GetBuildTime() string {

	i, _ := os.Stat(os.Args[0])

	dubai := time.FixedZone("Dubai", 4*60*60)

	dubaiTime := i.ModTime().In(dubai).Format(time.RFC3339)

	return fmt.Sprintf("%v", dubaiTime)
}
