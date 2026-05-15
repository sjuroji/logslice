package rate

import (
	"testing"
	"time"
)

func TestNew_ValidRate(t *testing.T) {
	l, err := New(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if l.Rate() != 10 {
		t.Fatalf("expected rate 10, got %d", l.Rate())
	}
}

func TestNew_ZeroReturnsError(t *testing.T) {
	_, err := New(0)
	if err == nil {
		t.Fatal("expected error for zero rate")
	}
}

func TestNew_NegativeReturnsError(t *testing.T) {
	_, err := New(-5)
	if err == nil {
		t.Fatal("expected error for negative rate")
	}
}

func TestAccept_NilAlwaysTrue(t *testing.T) {
	var l *Limiter
	for i := 0; i < 1000; i++ {
		if !l.Accept("line") {
			t.Fatal("nil Limiter should always accept")
		}
	}
}

func TestAccept_AllowsUpToLimit(t *testing.T) {
	l, _ := New(3)
	fixed := time.Now()
	l.nowFn = func() time.Time { return fixed }

	for i := 0; i < 3; i++ {
		if !l.Accept("line") {
			t.Fatalf("line %d should be accepted", i)
		}
	}
}

func TestAccept_DropsOverLimit(t *testing.T) {
	l, _ := New(3)
	fixed := time.Now()
	l.nowFn = func() time.Time { return fixed }

	for i := 0; i < 3; i++ {
		l.Accept("line")
	}
	if l.Accept("overflow") {
		t.Fatal("4th line in same window should be dropped")
	}
}

func TestAccept_ResetsAfterWindow(t *testing.T) {
	l, _ := New(2)
	fixed := time.Now()
	l.nowFn = func() time.Time { return fixed }

	l.Accept("a")
	l.Accept("b")
	if l.Accept("c") {
		t.Fatal("3rd line should be dropped in same window")
	}

	// advance clock by 1 second
	next := fixed.Add(time.Second)
	l.nowFn = func() time.Time { return next }

	if !l.Accept("d") {
		t.Fatal("first line in new window should be accepted")
	}
}

func TestRate_NilReturnsZero(t *testing.T) {
	var l *Limiter
	if l.Rate() != 0 {
		t.Fatalf("expected 0 for nil Limiter, got %d", l.Rate())
	}
}
