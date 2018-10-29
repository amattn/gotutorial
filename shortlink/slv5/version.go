package main

import (
	"fmt"
	"log"
	"time"
)

const (
	internal_BUILD_TIMESTAMP = 1500000000
	internal_BUILD_NUMBER    = 5
	internal_VERSION_STRING  = "0.5.0"
)

func BuildDate() time.Time {
	return time.Unix(internal_BUILD_TIMESTAMP, 0)
}
func BuildNumber() int64 {
	return internal_BUILD_NUMBER
}
func Version() string {
	return internal_VERSION_STRING
}

func VersionInfo() string {
	return fmt.Sprintf("shortlink (%v, build %v, build date:%v)", Version(), BuildNumber(), BuildDate())
}

func LogVersionInfo() {
	log.Printf(VersionInfo())
}
