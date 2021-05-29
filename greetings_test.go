package greetings

import (
	"regexp"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	testName := "testName"
	actualResp, err := Hello(testName)

	if err != nil || !contains(actualResp, testName) {
		t.Error("somethings wrong .. i can feel it")
	}
}

func TestHelloWithEmptyString(t *testing.T) {
	testName := ""
	actualResp, err := Hello(testName)

	if err == nil || actualResp != "" {
		t.Error("somethings wrong .. i can feel it")
	}
}

func contains(actualResp string, testName string) bool {
	want := regexp.MustCompile(`\b` + testName + `\b`)
	return want.MatchString(actualResp)
}
