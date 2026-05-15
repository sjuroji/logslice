package sample

import (
	"testing"
)

func TestNew_ValidN(t *testing.T) {
	s, err := New(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil Sampler")
	}
}

func TestNew_ZeroReturnsError(t *testing.T) {
	_, err := New(0)
	if err == nil {
		t.Fatal("expected error for n=0")
	}
}

func TestNew_NegativeReturnsError(t *testing.T) {
	_, err := New(-5)
	if err == nil {
		t.Fatal("expected error for n=-5")
	}
}

func TestAccept_EmitsEveryNthLine(t *testing.T) {
	s, _ := New(3)
	results := make([]bool, 9)
	for i := range results {
		results[i] = s.Accept("line")
	}
	expected := []bool{false, false, true, false, false, true, false, false, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("index %d: got %v, want %v", i, got, expected[i])
		}
	}
}

func TestAccept_NilAlwaysTrue(t *testing.T) {
	var s *Sampler
	for i := 0; i < 5; i++ {
		if !s.Accept("line") {
			t.Errorf("nil sampler returned false at iteration %d", i)
		}
	}
}

func TestReset_ResetsCounter(t *testing.T) {
	s, _ := New(2)
	s.Accept("a") // counter=1
	s.Accept("b") // counter=2, accepted
	s.Reset()
	// after reset counter=0, first call should be false
	if s.Accept("c") {
		t.Error("expected false after reset on first call")
	}
	if !s.Accept("d") {
		t.Error("expected true on second call after reset")
	}
}

func TestN_ReturnsInterval(t *testing.T) {
	s, _ := New(7)
	if s.N() != 7 {
		t.Errorf("got %d, want 7", s.N())
	}
}

func TestN_NilReturnsOne(t *testing.T) {
	var s *Sampler
	if s.N() != 1 {
		t.Errorf("got %d, want 1", s.N())
	}
}
