package main

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
)

var test_server *httptest.Server

func init() {
	test_server = httptest.NewServer(getHandler())
}

func assertEqual(t *testing.T, expected, candidate interface{}, printargs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printargs...)
		t.Errorf("Expected != Candidate\n%s\nExpected:\n%+v\nCandidate:\n%+v", extra, expected, candidate)
	}
}
