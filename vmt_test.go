package vmt

import (
	"testing"
)

func TestSprintq(t *testing.T) {
	s := "hello ?"
	var person = "bill"

	result:=Sprintq(s, person)

	if result!="hello bill" {
		t.Errorf("format result error [%v] expecting [%v]", result, "hello bill")
	}

}

func TestSprints(t *testing.T) {
	hello:="hello"

	result := Sprints("hello $1", hello)

	if result!="hello hello" {
		t.Errorf("result error, expected [hello hello], get [%v]", result)
	}
}

func TestPrintq(t *testing.T) {
	Printq("hello ? love ?", "I", "You")
}

func TestPrints(t *testing.T) {
	Prints("hello $2 love $1", "You", "I")
}