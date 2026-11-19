package main

import (
	"fmt"
	"strings"
)

// Item is a component that can be priced and displayed.
type Item interface {
	Price() float64
	Display() string
}

// Glass is an individual window component.
type Glass struct {
	Cost float64
}

func (g Glass) Price() float64 {
	return g.Cost
}

func (g Glass) Display() string {
	return fmt.Sprintf("Glass ($%.2f)", g.Price())
}

// Frame is an individual window component.
type Frame struct {
	Cost float64
}

func (f Frame) Price() float64 {
	return f.Cost
}

func (f Frame) Display() string {
	return fmt.Sprintf("Frame ($%.2f)", f.Price())
}

// Window is a composite item containing other items.
type Window struct {
	Components []Item
}

func (w Window) Price() float64 {
	var total float64
	for _, component := range w.Components {
		total += component.Price()
	}
	return total
}

func (w Window) Display() string {
	var display strings.Builder
	fmt.Fprintf(&display, "Window ($%.2f)", w.Price())
	for _, component := range w.Components {
		for _, line := range strings.Split(component.Display(), "\n") {
			fmt.Fprintf(&display, "\n  %s", line)
		}
	}
	return display.String()
}
