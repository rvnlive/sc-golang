package segmentpb

import (
	"testing"
	"time"

	"github.com/smart-core-os/sc-api/go/traits"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestMaxMagnitude(t *testing.T) {
	tests := []struct {
		name    string
		args    []*traits.ElectricMode_Segment
		wantMax float32
	}{
		{"empty", segs(), 0},
		{"one-cap", segs(s{10, 0}), 10},
		{"one-inf", segs(s{10, 10}), 10},
		{"few", segs(s{2, 10}, s{6, 10}, s{4, 0}), 6},
		{"inf in the middle", segs(s{2, 10}, s{4, 0}, s{6, 10}), 6},
		{"zero-length segment", []*traits.ElectricMode_Segment{
			{Magnitude: 2, Length: durationpb.New(10)},
			{Magnitude: 6, Length: durationpb.New(0)}, // not counted
			{Magnitude: 4, Length: durationpb.New(10)},
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := MaxMagnitude(tt.args...); gotMax != tt.wantMax {
				t.Errorf("MaxMagnitude() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestSumMagnitude(t *testing.T) {
	tests := []struct {
		name    string
		args    []*traits.ElectricMode_Segment
		wantSum float32
	}{
		{"empty", segs(), 0},
		{"one-cap", segs(s{10, 10}), 10},
		{"one-inf", segs(s{10, 0}), 10},
		{"some", segs(s{1, 10}, s{2, 10}, s{3, 10}, s{5, 10}), 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumMagnitude(tt.args...); gotSum != tt.wantSum {
				t.Errorf("SumMagnitude() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestMagnitudeAt(t *testing.T) {
	tests := []struct {
		name      string
		d         time.Duration
		segments  []*traits.ElectricMode_Segment
		wantLevel float32
		wantOk    bool
	}{
		{"empty", 0, segs(), 0, false},
		{"empty+1", 1 * time.Second, segs(), 0, false},
		{"one@-1", -1, segs(s{10, 10}), 0, false},
		{"one@0", 0, segs(s{10, 10}), 10, true},
		{"one@1", 1, segs(s{10, 10}), 10, true},
		{"one@10", 10, segs(s{10, 10}), 0, false},
		{"inf@-1", -1, segs(s{10, 0}), 0, false},
		{"inf@0", 0, segs(s{10, 0}), 10, true},
		{"inf@1", 1, segs(s{10, 0}), 10, true},
		{"inf@10", 10, segs(s{10, 0}), 10, true},
		{"inf@1000", 1000, segs(s{10, 0}), 10, true},
		{"few@-1", -1, segs(s{10, 10}, s{4, 10}), 0, false},
		{"few@0", 0, segs(s{10, 10}, s{4, 10}), 10, true},
		{"few@1", 1, segs(s{10, 10}, s{4, 10}), 10, true},
		{"few@10", 10, segs(s{10, 10}, s{4, 10}), 4, true},
		{"few@11", 11, segs(s{10, 10}, s{4, 10}), 4, true},
		{"few@20", 20, segs(s{10, 10}, s{4, 10}), 0, false},
		{"few+inf@-1", -1, segs(s{10, 10}, s{4, 0}), 0, false},
		{"few+inf@0", 0, segs(s{10, 10}, s{4, 0}), 10, true},
		{"few+inf@1", 1, segs(s{10, 10}, s{4, 0}), 10, true},
		{"few+inf@10", 10, segs(s{10, 10}, s{4, 0}), 4, true},
		{"few+inf@11", 11, segs(s{10, 10}, s{4, 0}), 4, true},
		{"few+inf@20", 20, segs(s{10, 10}, s{4, 0}), 4, true},
		{"few+inf@2000", 2000, segs(s{10, 10}, s{4, 0}), 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLevel, gotOk := MagnitudeAt(tt.d, tt.segments...)
			if gotLevel != tt.wantLevel {
				t.Errorf("MagnitudeAt() gotLevel = %v, want %v", gotLevel, tt.wantLevel)
			}
			if gotOk != tt.wantOk {
				t.Errorf("MagnitudeAt() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
