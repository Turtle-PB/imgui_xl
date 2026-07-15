# Go Bindings for Dear ImGui (imgui_xl)

This directory provides Go (cgo) bindings for Dear ImGui via a thin C API bridge.

## Architecture

```
imgui.go          → Go package (cgo wrapper)
imgui_bridge.h    → C API declarations (extern "C")
imgui_bridge.cpp  → C API implementation (wraps C++ ImGui)
go.mod            → Go module definition
```

The Go package calls C functions via cgo, which in turn call the C++ ImGui API through the bridge layer.

## Prerequisites

- **Go 1.21+** (with `CGO_ENABLED=1`)
- **C++ compiler** (g++ or clang++) on PATH

## Usage

```go
package main

import (
    "fmt"
    imgui "github.com/Turtle-PB/imgui_xl/bindings/go"
)

func main() {
    imgui.CreateContext()
    defer imgui.DestroyContext()
    
    imgui.StyleColorsDark()
    
    imgui.NewFrame()
    visible := imgui.Begin("Hello, Go!", nil)
    if visible {
        imgui.Text("Hello from Go!")
        imgui.Text(fmt.Sprintf("FPS: %.1f", imgui.GetFramerate()))
    }
    imgui.End()
    imgui.Render()
}
```

## Building

```bash
# Set up the Go module (from the example directory)
cd examples/example_golang
go mod tidy

# Build (requires C++ compiler)
CGO_ENABLED=1 go build -o imgui_go_example

# Run
./imgui_go_example
```

## Available Functions

### Core
- `CreateContext()`, `DestroyContext()`
- `NewFrame()`, `Render()`
- `Version()`, `GetFramerate()`, `GetTime()`, `GetFrameCount()`
- `GetDisplaySize()`

### Style
- `StyleColorsDark()`, `StyleColorsLight()`, `StyleColorsClassic()`

### Windows
- `Begin(name, &open)`, `End()`
- `ShowDemoWindow(&open)`

### Widgets
- `Text(s)`, `TextColored(r,g,b,a,s)`, `BulletText(s)`
- `Button(label, w, h)`, `Checkbox(label, &v)`
- `SliderFloat(label, &v, min, max)`, `SliderInt(label, &v, min, max)`
- `ColorEdit3(label, &col)`, `InputText(label, &buf, size)`
- `Combo(label, &current, items, count)`

### Layout
- `SameLine(offset, spacing)`, `Separator()`, `Spacing()`, `NewLine()`
- `PushItemWidth(w)`, `PopItemWidth()`

## License

MIT — same as Dear ImGui
