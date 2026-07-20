# Go Bindings for Dear ImGui (imgui_xl)

This directory provides Go (cgo) bindings for Dear ImGui via a thin C API bridge.

## Architecture

```
imgui.go          → Go package (cgo wrapper)
imgui_bridge.h    → C API declarations (extern "C")
imgui_bridge.cpp  → C API implementation (wraps C++ ImGui)
go.mod            → Go module definition
```

The Go package calls C functions via cgo, which calls the C++ ImGui API through the bridge.

---

## Prerequisites

| Requirement | Windows | Linux | macOS |
|---|---|---|---|
| **Go** | 1.21+ | 1.21+ | 1.21+ |
| **C++ compiler** | MinGW-w64 gcc/g++ | g++ | clang++ (Xcode CLT) |
| **CGO** | `CGO_ENABLED=1` | enabled by default | enabled by default |

> **⚠️ Most common failure:** `CGO_ENABLED=0` because Go can't find a C compiler.
> Check with `go env CC CXX CGO_ENABLED`. If `CGO_ENABLED=0`, cgo is disabled and the build will fail with "build constraints exclude all Go files".

### Windows Setup (the tricky one)

Go's cgo on Windows requires a **gcc-compatible** compiler. The default `CC=gcc`, `CXX=g++`.

**Option A — MinGW-w64 (recommended):**
```bash
# Install via MSYS2: https://www.msys2.org/
pacman -S mingw-w64-x86_64-gcc
# Add to PATH: C:\msys64\mingw64\bin

# Verify
gcc --version
g++ --version
go env CC CXX CGO_ENABLED   # should show CGO_ENABLED=1
```

**Option B — TDM-GCC:**
Download from https://jmeubank.github.io/tdm-gcc/ and add `bin` to PATH.

**⚠️ LLVM/Clang on Windows does NOT work out of the box:**
`clang++` targeting `x86_64-pc-windows-msvc` fails with `unsupported option '-mthreads'`.
If you must use clang, you need a MinGW-targeted clang build and:
```bash
set CC=clang
set CXX=clang++
set CGO_CFLAGS=-D__USE_MINGW_ANSI_STDIO
```
(MinGW-w64 gcc is far easier.)

### Linux Setup
```bash
sudo apt install build-essential golang   # Debian/Ubuntu
sudo dnf install gcc-c++ golang           # Fedora
```

### macOS Setup
```bash
xcode-select --install   # provides clang/clang++
brew install go
```

---

## Building the Example

```bash
# Clone the repo
git clone https://github.com/Turtle-PB/imgui_xl.git
cd imgui_xl/examples/example_golang

# Download module deps (uses a local `replace` directive to ../../bindings/go)
go mod tidy

# Build
#   Linux/macOS:
CGO_ENABLED=1 go build -o imgui_go_example .

#   Windows (MinGW on PATH):
set CGO_ENABLED=1
go build -o imgui_go_example.exe .

# Run
./imgui_go_example        # Linux/macOS
imgui_go_example.exe      # Windows
```

Or use the provided Makefile (Linux/macOS/MinGW):
```bash
cd examples/example_golang
make build    # compile
make run      # compile + run
make clean    # remove artifacts
```

---

## Using the Bindings in Your Own Project

The bindings module lives inside this repo, so use a `replace` directive:

```bash
# In your project directory
go mod init myapp
go mod edit -require=github.com/Turtle-PB/imgui_xl/bindings/go@v0.0.0
go mod edit -replace=github.com/Turtle-PB/imgui_xl/bindings/go=/path/to/imgui_xl/bindings/go
go mod tidy
```

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
    if imgui.Begin("Hello, Go!", nil) {
        imgui.Text("Hello from Go!")
        imgui.Text(fmt.Sprintf("FPS: %.1f", imgui.GetFramerate()))
    }
    imgui.End()
    imgui.Render()
}
```

> **Note:** This example runs in "null" mode (no window). For a real window,
> pair with a Go windowing library (GLFW, SDL2) and a renderer backend.
> The cgo wrapper calls into the same C++ backends in `backends/`.

---

## Troubleshooting

| Error | Cause | Fix |
|---|---|---|
| `build constraints exclude all Go files` | `CGO_ENABLED=0` | Set `CGO_ENABLED=1` and install a C compiler |
| `gcc: command not found` / `g++: command not found` | No C++ compiler on PATH | Install MinGW-w64 (Windows) or build-essential (Linux) |
| `unsupported option '-mthreads'` (Windows) | Using LLVM clang targeting MSVC | Use MinGW-w64 gcc instead |
| `undefined reference to 'ig_*'` | Bridge .cpp not compiled | Ensure `imgui_bridge.cpp` is in the same dir as `imgui.go` (it is, in this repo) |
| `cannot find -lstdc++` | Missing C++ stdlib | Install g++ (not just gcc) |
| cgo flag errors mentioning `-std=c++11` | Compiler too old | Use GCC 5+ / Clang 3.4+ |

### Verify your toolchain before building:
```bash
go env CC CXX CGO_ENABLED
# Expected (Windows MinGW): CC=gcc  CXX=g++  CGO_ENABLED=1
# Expected (Linux):         CC=gcc  CXX=g++  CGO_ENABLED=1
# Expected (macOS):         CC=clang CXX=clang++ CGO_ENABLED=1
```

---

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

MIT — same as Dear ImGui. Copyright (c) 2024 OwlLogics / imgui_xl.
