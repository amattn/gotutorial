package main

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
)

var test_server *httptest.Server
var router *LoggingRouter

func init() {
	router = NewLoggingRouter()
	test_server = httptest.NewServer(router)
}

func TestNothing(t *testing.T) {
	// test test harness by uncommenting next line:
	// t.Error("93420033538 intentionally induced error")
}

func assertEqual(t *testing.T, expected, candidate interface{}, printargs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printargs...)
		t.Errorf("Expected != Candidate\n%s\nExpected:\n%+v\nCandidate:\n%+v", extra, expected, candidate)
	}
}
