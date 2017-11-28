package vmt

import (
	"testing"
	"fmt"
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
	fmt.Println(result)
}