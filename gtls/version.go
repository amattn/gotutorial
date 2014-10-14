package main

const (
	internal_BUILD_NUMBER   = 12
	internal_VERSION_STRING = "0.1.0"
)

func BuildNumber() int64 {
	return internal_BUILD_NUMBER
}
func Version() string {
	return internal_VERSION_STRING
}
