package fieldextract_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/fieldextract"
)

func TestNew_DefaultSeparator(t *testing.T) {
	e, err := fieldextract.New("", []int{1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e == nil {
		t.Fatal("expected non-nil extractor")
	}
}

func TestNew_InvalidIndex(t *testing.T) {
	_, err := fieldextract.New(" ", []int{0})
	if err == nil {
		t.Fatal("expected error for index 0")
	}
}

func TestNew_NegativeIndex(t *testing.T) {
	_, err := fieldextract.New(",", []int{-1})
	if err == nil {
		t.Fatal("expected error for negative index")
	}
}

func TestApply_NilReturnsOriginal(t *testing.T) {
	var e *fieldextract.Extractor
	got := e.Apply("hello world")
	if got != "hello world" {
		t.Errorf("got %q, want %q", got, "hello world")
	}
}

func TestApply_NoIndicesReturnsOriginal(t *testing.T) {
	e, _ := fieldextract.New(" ", nil)
	got := e.Apply("a b c")
	if got != "a b c" {
		t.Errorf("got %q, want %q", got, "a b c")
	}
}

func TestApply_SingleField(t *testing.T) {
	e, _ := fieldextract.New(" ", []int{2})
	got := e.Apply("alpha beta gamma")
	if got != "beta" {
		t.Errorf("got %q, want %q", got, "beta")
	}
}

func TestApply_MultipleFields(t *testing.T) {
	e, _ := fieldextract.New(",", []int{1, 3})
	got := e.Apply("foo,bar,baz")
	if got != "foo,baz" {
		t.Errorf("got %q, want %q", got, "foo,baz")
	}
}

func TestApply_OutOfRangeFieldSkipped(t *testing.T) {
	e, _ := fieldextract.New(" ", []int{1, 99})
	got := e.Apply("only one")
	// field 99 doesn't exist; only field 1 returned
	if got != "only" {
		t.Errorf("got %q, want %q", got, "only")
	}
}

func TestApply_AllOutOfRange(t *testing.T) {
	e, _ := fieldextract.New(" ", []int{10})
	got := e.Apply("a b")
	if got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}

func TestFields_CountsCorrectly(t *testing.T) {
	e, _ := fieldextract.New(":", nil)
	if n := e.Fields("a:b:c"); n != 3 {
		t.Errorf("got %d, want 3", n)
	}
}

func TestFields_NilReturnsZero(t *testing.T) {
	var e *fieldextract.Extractor
	if n := e.Fields("x y z"); n != 0 {
		t.Errorf("got %d, want 0", n)
	}
}
