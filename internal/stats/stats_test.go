package stats_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/stats"
)

func TestNew_InitialisesStartTime(t *testing.T) {
	before := time.Now()
	s := stats.New()
	after := time.Now()

	if s.StartedAt.Before(before) || s.StartedAt.After(after) {
		t.Errorf("StartedAt %v not in expected range [%v, %v]", s.StartedAt, before, after)
	}
}

func TestAddRead_AccumulatesCount(t *testing.T) {
	s := stats.New()
	s.AddRead(10)
	s.AddRead(5)
	if s.LinesRead != 15 {
		t.Errorf("expected LinesRead=15, got %d", s.LinesRead)
	}
}

func TestAddMatched_AccumulatesCount(t *testing.T) {
	s := stats.New()
	s.AddMatched(3)
	s.AddMatched(7)
	if s.LinesMatched != 10 {
		t.Errorf("expected LinesMatched=10, got %d", s.LinesMatched)
	}
}

func TestAddFile_Increments(t *testing.T) {
	s := stats.New()
	s.AddFile()
	s.AddFile()
	if s.FilesRead != 2 {
		t.Errorf("expected FilesRead=2, got %d", s.FilesRead)
	}
}

func TestFinish_SetsFinishedAt(t *testing.T) {
	s := stats.New()
	s.Finish()
	if s.FinishedAt.IsZero() {
		t.Error("FinishedAt should not be zero after Finish()")
	}
}

func TestElapsed_BeforeFinish(t *testing.T) {
	s := stats.New()
	time.Sleep(2 * time.Millisecond)
	elapsed := s.Elapsed()
	if elapsed < time.Millisecond {
		t.Errorf("expected elapsed >= 1ms, got %v", elapsed)
	}
}

func TestElapsed_AfterFinish(t *testing.T) {
	s := stats.New()
	time.Sleep(2 * time.Millisecond)
	s.Finish()
	elapsed := s.Elapsed()
	if elapsed < time.Millisecond {
		t.Errorf("expected elapsed >= 1ms, got %v", elapsed)
	}
}

func TestWrite_ContainsFields(t *testing.T) {
	s := stats.New()
	s.AddFile()
	s.AddRead(100)
	s.AddMatched(42)
	s.Finish()

	var buf bytes.Buffer
	s.Write(&buf)
	out := buf.String()

	for _, want := range []string{"files read", "lines read", "lines matched", "elapsed", "100", "42", "1"} {
		if !strings.Contains(out, want) {
			t.Errorf("Write output missing %q\nGot:\n%s", want, out)
		}
	}
}
