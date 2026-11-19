package main

import "testing"

func TestWindowPrice(t *testing.T) {
	window := Window{
		Components: []Item{
			Glass{Cost: 25},
			Frame{Cost: 10},
			Window{Components: []Item{
				Glass{Cost: 15},
				Frame{Cost: 5},
			}},
		},
	}

	if got, want := window.Price(), 55.0; got != want {
		t.Errorf("Window.Price() = %v, want %v", got, want)
	}
}

func TestWindowDisplayHierarchy(t *testing.T) {
	window := Window{
		Components: []Item{
			Glass{Cost: 25},
			Window{Components: []Item{
				Frame{Cost: 10},
			}},
		},
	}

	const want = "Window ($35.00)\n" +
		"  Glass ($25.00)\n" +
		"  Window ($10.00)\n" +
		"    Frame ($10.00)"
	if got := window.Display(); got != want {
		t.Errorf("Window.Display() =\n%s\nwant:\n%s", got, want)
	}
}
