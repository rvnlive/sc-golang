package resource

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/smart-core-os/sc-api/go/traits"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestValue_Pull(t *testing.T) {
	t.Run("SeedValue", func(t *testing.T) {
		now := time.UnixMilli(0)
		clock := clockFunc(func() time.Time {
			return now
		})

		v := NewValue(WithInitialValue(&traits.OnOff{State: traits.OnOff_ON}), WithClock(clock))

		ctx, stop := context.WithCancel(context.Background())
		t.Cleanup(stop)
		changes := v.Pull(ctx, WithBackpressure(false))

		// first value when not using UpdatesOnly should say it's not an update
		seed := waitForChan(t, changes, time.Second)
		want := &ValueChange{
			ChangeTime: now,
			Value:      &traits.OnOff{State: traits.OnOff_ON},
			SeedValue:  true,
		}
		if diff := cmp.Diff(want, seed, protocmp.Transform()); diff != "" {
			t.Fatalf("Seed Value (-want,+got)\n%s", diff)
		}

		// second value should be an update
		v.Set(&traits.OnOff{State: traits.OnOff_OFF})
		next := waitForChan(t, changes, time.Second)
		want = &ValueChange{
			ChangeTime: now,
			Value:      &traits.OnOff{State: traits.OnOff_OFF},
			SeedValue:  false,
		}
		if diff := cmp.Diff(want, next, protocmp.Transform()); diff != "" {
			t.Fatalf("Next Value (-want,+got)\n%s", diff)
		}
	})

	t.Run("SeedValue updatesOnly", func(t *testing.T) {
		now := time.UnixMilli(0)
		clock := clockFunc(func() time.Time {
			return now
		})

		v := NewValue(WithInitialValue(&traits.OnOff{State: traits.OnOff_ON}), WithClock(clock))

		ctx, stop := context.WithCancel(context.Background())
		t.Cleanup(stop)
		changes := v.Pull(ctx, WithBackpressure(false), WithUpdatesOnly(true))

		// with updates only, there should be no waiting event
		noEmitWithin(t, changes, 50*time.Millisecond)

		// first value should be an update
		v.Set(&traits.OnOff{State: traits.OnOff_OFF})
		change := waitForChan(t, changes, time.Second)
		want := &ValueChange{
			ChangeTime: now,
			Value:      &traits.OnOff{State: traits.OnOff_OFF},
			SeedValue:  false,
		}
		if diff := cmp.Diff(want, change, protocmp.Transform()); diff != "" {
			t.Fatalf("Value (-want,+got)\n%s", diff)
		}
	})
}
