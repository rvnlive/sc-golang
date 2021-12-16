package electric

import (
	"context"
	"fmt"
	"github.com/smart-core-os/sc-golang/pkg/time/clock"
	"log"

	"github.com/smart-core-os/sc-api/go/traits"
)

func ExampleMemoryDevice() {
	mem := NewMemory(clock.Real())
	device := NewMemoryDevice(mem)

	client := Wrap(device)
	settings := WrapMemorySettings(device)

	ctx := context.Background()
	_, err := settings.CreateMode(ctx, &CreateModeRequest{
		Name: "foo",
		Mode: &traits.ElectricMode{
			Title:       "Normal mode",
			Description: "Normal mode",
			Segments: []*traits.ElectricMode_Segment{
				{Magnitude: 1},
			},
			Normal: true,
		},
	})
	if err != nil {
		log.Println("create mode failed:", err)
		return
	}

	_, err = client.ClearActiveMode(ctx, &traits.ClearActiveModeRequest{
		Name: "foo",
	})
	if err != nil {
		log.Println("clear mode failed:", err)
	}

	mode, err := client.GetActiveMode(ctx, &traits.GetActiveModeRequest{Name: "foo"})
	if err != nil {
		log.Println("GetActiveMode failed:", err)
		return
	}
	fmt.Println(mode.Title)
	// Output: Normal mode
}
