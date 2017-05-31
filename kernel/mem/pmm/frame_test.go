package pmm

import (
	"testing"

	"github.com/achilleasa/gopher-os/kernel/mem"
)

func TestFrameMethods(t *testing.T) {
	for order := mem.PageOrder(0); order < mem.PageOrder(10); order++ {
		for frameIndex := uint64(0); frameIndex < 128; frameIndex++ {
			frame := Frame(frameIndex | (uint64(order) << 56))

			if !frame.IsValid() {
				t.Errorf("[order %d] expected frame %d to be valid", order, frameIndex)
			}

			if exp, got := uintptr(frameIndex<<mem.PageShift), frame.Address(); got != exp {
				t.Errorf("[order %d] expected frame (%d, index: %d) call to Address() to return %x; got %x", order, frame, frameIndex, exp, got)
			}
		}
	}
}
