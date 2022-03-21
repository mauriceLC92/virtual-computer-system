package gmachine_test

import (
	"testing"

	"github.com/mauriceLC92/virtual-computer-system/gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Run()

	wantP := 1
	gotP := g.P
	if wantP != int(gotP) {
		t.Errorf("wanted counter value %d but got %d instead", wantP, gotP)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.P = uint64(gmachine.NOOP)
	g.Run()

	wantP := 2
	gotP := g.P
	if wantP != int(gotP) {
		t.Errorf("wanted counter value %d but got %d instead", wantP, gotP)
	}
}
