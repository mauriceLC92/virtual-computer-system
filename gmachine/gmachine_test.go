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

func TestINCA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.Memory[0] = uint64(gmachine.INCA)

	g.Run()

	wantA := 1
	gotA := g.A

	if wantA != int(gotA) {
		t.Errorf("wanted %d but got %d", wantA, gotA)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.A = 2
	g.Memory[0] = uint64(gmachine.DECA)

	g.Run()

	wantA := 1
	gotA := g.A

	if wantA != int(gotA) {
		t.Errorf("wanted %d but got %d", wantA, gotA)
	}
}

// TestSubtraction tests for the machine to be able to subtract 2 from 3
func TestSubtraction(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.A = 3
	g.Memory[0] = uint64(gmachine.DECA)
	g.Memory[1] = uint64(gmachine.DECA)

	g.Run()

	wantA := 1
	gotA := g.A

	if wantA != int(gotA) {
		t.Errorf("wanted %d but got %d", wantA, gotA)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		uint64(gmachine.SETA), 5,
		uint64(gmachine.HALT),
	})
	if g.A != 5 {
		t.Errorf("want A == 5, got %d", g.A)
	}
	if g.P != 3 {
		t.Errorf("want P == 3, got %d", g.P)
	}
}

func TestSub3(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input, want uint64
	}{
		{input: 3, want: 1},
		{input: 100, want: 98},
		{input: 5, want: 3},
	}
	for _, tc := range tcs {
		g := gmachine.New()
		g.RunProgram([]uint64{
			uint64(gmachine.SETA),
			tc.input,
			uint64(gmachine.DECA),
			uint64(gmachine.DECA),
			uint64(gmachine.HALT),
		})
		if g.A != tc.want {
			t.Errorf("want A == %d, got %d", tc.want, g.A)
		}
	}
}
