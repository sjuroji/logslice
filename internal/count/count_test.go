package count_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/count"
)

func TestNew_InitialisesZero(t *testing.T) {
	c := count.New()
	if c.Value() != 0 {
		t.Fatalf("expected 0, got %d", c.Value())
	}
}

func TestInc_IncrementsOne(t *testing.T) {
	c := count.New()
	c.Inc()
	c.Inc()
	if c.Value() != 2 {
		t.Fatalf("expected 2, got %d", c.Value())
	}
}

func TestAdd_IncrementsBy(t *testing.T) {
	c := count.New()
	c.Add(5)
	c.Add(3)
	if c.Value() != 8 {
		t.Fatalf("expected 8, got %d", c.Value())
	}
}

func TestReset_SetsZero(t *testing.T) {
	c := count.New()
	c.Add(10)
	c.Reset()
	if c.Value() != 0 {
		t.Fatalf("expected 0 after reset, got %d", c.Value())
	}
}

func TestAccumulator_RecordAndTotal(t *testing.T) {
	a := count.NewAccumulator()
	a.Record("a.log", 10)
	a.Record("b.log", 25)
	if a.Total() != 35 {
		t.Fatalf("expected total 35, got %d", a.Total())
	}
}

func TestAccumulator_Files_Snapshot(t *testing.T) {
	a := count.NewAccumulator()
	a.Record("x.log", 7)
	files := a.Files()
	if files["x.log"] != 7 {
		t.Fatalf("expected 7 for x.log, got %d", files["x.log"])
	}
	// mutation of snapshot must not affect accumulator
	files["x.log"] = 999
	if a.Files()["x.log"] != 7 {
		t.Fatal("snapshot mutation affected accumulator")
	}
}

func TestAccumulator_MultipleRecords_SameFile(t *testing.T) {
	a := count.NewAccumulator()
	a.Record("dup.log", 3)
	a.Record("dup.log", 4)
	// last write wins for per-file map; total accumulates both calls
	if a.Total() != 7 {
		t.Fatalf("expected total 7, got %d", a.Total())
	}
}
