package main

import (
	"testing"
)

func TestMessagePool(t *testing.T) {
	msg0 := Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}

	msg0.Count = 1
	Instance().AddMsg(msg0)
	msg1 := Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg0.Count)
	}
}

func TestMessagePool2(t *testing.T) {
	msg0 := Instance2().GetMsg()
	msg1 := Instance2().GetMsg()

	if msg0 != msg1 {
		t.Errorf("expect msg0 is same to msg1, but actual is not")
	}
}
