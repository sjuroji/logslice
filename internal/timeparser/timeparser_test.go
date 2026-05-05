package timeparser_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/timeparser"
)

func TestParse_SupportedFormats(t *testing.T) {
	p := timeparser.New(time.UTC)

	cases := []struct {
		input    string
		wantYear int
		wantHour int
	}{
		{"2024-03-15T08:30:00Z", 2024, 8},
		{"2024-03-15T08:30:00", 2024, 8},
		{"2024-03-15 08:30:00", 2024, 8},
		{"2024-03-15 08:30:00.123", 2024, 8},
		{"2024/03/15 08:30:00", 2024, 8},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, format, err := p.Parse(tc.input)
			if err != nil {
				t.Fatalf("Parse(%q) unexpected error: %v", tc.input, err)
			}
			if got.Year() != tc.wantYear {
				t.Errorf("year: got %d, want %d", got.Year(), tc.wantYear)
			}
			if got.Hour() != tc.wantHour {
				t.Errorf("hour: got %d, want %d", got.Hour(), tc.wantHour)
			}
			if format == "" {
				t.Error("expected non-empty format string")
			}
		})
	}
}

func TestParse_UnknownFormat(t *testing.T) {
	p := timeparser.New(time.UTC)
	_, _, err := p.Parse("not-a-timestamp")
	if err == nil {
		t.Fatal("expected error for unrecognized timestamp, got nil")
	}
}

func TestParseWithFormat(t *testing.T) {
	p := timeparser.New(time.UTC)
	got, err := p.ParseWithFormat("2024-03-15 08:30:00", "2006-01-02 15:04:05")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Month() != time.March {
		t.Errorf("month: got %v, want March", got.Month())
	}
}

func TestNew_NilLocation(t *testing.T) {
	p := timeparser.New(nil)
	if p.Location != time.UTC {
		t.Errorf("expected UTC location, got %v", p.Location)
	}
}
