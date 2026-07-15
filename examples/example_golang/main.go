// Package main demonstrates using Dear ImGui from Go.
// This example runs ImGui in a headless "null" mode — it creates a context,
// runs 20 frames of rendering, and prints the output.
//
// For a full graphical example, pair this with a Go windowing library
// (e.g., GLFW, SDL2) and a renderer backend (OpenGL3, Vulkan).
//
// Build requirements:
//   - C++ compiler (g++ or clang++) on PATH
//   - CGO_ENABLED=1 (default when a C compiler is available)
//
// Build:
//   cd examples/example_golang
//   CGO_ENABLED=1 go build -o imgui_go_example
//   ./imgui_go_example
//
// Copyright (c) 2024 OwlLogics / imgui_xl
// Licensed under MIT

package main

import (
	"fmt"

	imgui "github.com/Turtle-PB/imgui_xl/bindings/go"
)

func main() {
	fmt.Println("=== Dear ImGui Go Example ===")
	fmt.Printf("ImGui Version: %s\n", imgui.Version())
	fmt.Println()

	// Create context
	imgui.CreateContext()
	defer imgui.DestroyContext()

	// Use dark theme
	imgui.StyleColorsDark()

	// Run 20 frames in "null" mode (no rendering output)
	showDemo := true
	showAnother := false
	floatVal := 0.0
	counter := 0

	for frame := 0; frame < 20; frame++ {
		imgui.NewFrame()

		// Show the ImGui demo window
		imgui.ShowDemoWindow(&showDemo)

		// Create a custom window
		visible := imgui.Begin("Hello from Go!", &showDemo)
		if visible {
			imgui.Text(fmt.Sprintf("Frame: %d", frame))
			imgui.Text(fmt.Sprintf("FPS: %.1f", imgui.GetFramerate()))
			imgui.Separator()

			imgui.Checkbox("Another Window", &showAnother)

			f32Val := float32(floatVal)
			if imgui.SliderFloat("Float", &f32Val, 0.0, 1.0) {
				floatVal = float64(f32Val)
			}

			if imgui.Button("Click Me", 0, 0) {
				counter++
				fmt.Printf("  Button clicked! Count: %d\n", counter)
			}

			imgui.SameLine(0, 4)
			imgui.Text(fmt.Sprintf("counter = %d", counter))
		}
		imgui.End()

		// Show another window toggled by checkbox
		if showAnother {
			visible := imgui.Begin("Another Window", &showAnother)
			if visible {
				imgui.Text("Hello from the second window!")
				imgui.Text("This window can be closed via the checkbox.")
			}
			imgui.End()
		}

		imgui.Render()

		if frame == 0 || frame == 9 || frame == 19 {
			w, h := imgui.GetDisplaySize()
			fmt.Printf("Frame %2d: Render complete (display: %.0fx%.0f, FPS: %.1f)\n",
				frame, w, h, imgui.GetFramerate())
		}
	}

	fmt.Println()
	fmt.Println("=== Done! ===")
	fmt.Println("20 frames rendered in null mode.")
	fmt.Println()
	fmt.Println("To build a graphical Go app with ImGui:")
	fmt.Println("  1. Install a Go windowing lib:  go get github.com/go-gl/glfw/v3.3/glfw")
	fmt.Println("  2. Install an OpenGL Go lib:     go get github.com/go-gl/gl/v4.1-core/gl")
	fmt.Println("  3. Create a window, init OpenGL,")
	fmt.Println("     call imgui.CreateContext() + ImGui_ImplOpenGL3_Init()")
	fmt.Println("  4. Loop: poll events -> NewFrame -> draw UI -> Render -> RenderDrawData")
}
