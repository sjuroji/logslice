package config

import (
	"flag"
	"testing"
)

func newSampleFS(args ...string) *flag.FlagSet {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerSampleFlag(fs)
	_ = fs.Parse(args)
	return fs
}

func TestSampleN_NilFlagSet(t *testing.T) {
	if got := SampleN(nil); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

func TestSampleN_DefaultConfig(t *testing.T) {
	fs := newSampleFS()
	if got := SampleN(fs); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

func TestSampleN_ExplicitValue(t *testing.T) {
	fs := newSampleFS("--sample", "5")
	if got := SampleN(fs); got != 5 {
		t.Errorf("got %d, want 5", got)
	}
}

func TestSampleN_ShortFlag(t *testing.T) {
	fs := newSampleFS("-S", "10")
	if got := SampleN(fs); got != 10 {
		t.Errorf("got %d, want 10", got)
	}
}

func TestSampleN_BelowOneReturnsOne(t *testing.T) {
	fs := newSampleFS("--sample", "0")
	if got := SampleN(fs); got != 1 {
		t.Errorf("got %d, want 1 for n=0", got)
	}
}

func TestSampleN_UnregisteredFlag(t *testing.T) {
	fs := flag.NewFlagSet("bare", flag.ContinueOnError)
	if got := SampleN(fs); got != 1 {
		t.Errorf("got %d, want 1 for unregistered flag", got)
	}
}
