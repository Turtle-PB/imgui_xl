# imgui_xl

**Dear ImGui XL** — An enhanced community fork of [Dear ImGui](https://github.com/ocornut/imgui), focused on backend improvements, CI/CD fixes, and extended tooling.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Upstream](https://img.shields.io/badge/upstream-ocornut%2Fimgui-blue)](https://github.com/ocornut/imgui)
[![Live Demo](https://img.shields.io/badge/demo-GitHub%20Pages-green)](https://turtle-pb.github.io/imgui_xl/)
[![PRs Merged](https://img.shields.io/badge/upstream%20PRs%20merged-3-brightgreen)](https://github.com/ocornut/imgui/pulls?q=is%3Apr+author%3ATurtle-PB)

> Dear ImGui is a bloat-free graphical user interface library for C++. It outputs optimized vertex buffers that you can render anytime in your 3D-pipeline-enabled application. It is fast, portable, renderer-agnostic, and self-contained (no external dependencies).

## Live Demos

| Demo | Description | Link |
|------|-------------|------|
| **Live ImGui (WASM)** | Actual Dear ImGui compiled to WebAssembly via Emscripten | [wasm/index.html](https://turtle-pb.github.io/imgui_xl/wasm/index.html) |
| **Showcase** | Landing page with all demos + PR list | [turtle-pb.github.io/imgui_xl](https://turtle-pb.github.io/imgui_xl/) |
| **Echonation Visual** | 5-mode audio-reactive canvas demo | [echonation.html](https://turtle-pb.github.io/imgui_xl/echonation.html) |
| **Widget Playground** | Visual GUI Builder — drag & drop widget designer (Visual Basic/Delphi style) | [widgets_playground.html](https://turtle-pb.github.io/imgui_xl/widgets_playground.html) |
| **Pixel OS Model** | Interactive Pixel 8 Pro phone simulator with working browser | [pixel_os_demo.html](https://turtle-pb.github.io/imgui_xl/pixel_os_demo.html) |

## Widget Playground (Visual GUI Builder)

A **drag-and-drop GUI designer** for Dear ImGui, inspired by Visual Basic and Delphi form builders:

- **Toolbox** with 12 widget types: Button, Text, InputText, Checkbox, RadioButton, Slider, ProgressBar, ColorEdit, Combo, Separator, Group, SameLine
- **Drag & drop** widgets from toolbox to canvas
- **Move & resize** widgets directly on canvas
- **Properties panel** — edit position, size, text, color, visibility, enabled state
- **Widget list** — hierarchical view of all placed widgets
- **Export** — generate ImGui C++ code or JSON layout to clipboard
- **Console** — command interface with `help`, `clear`, `count`, `list`, `delete`, `export`

### Sample Layouts

Five ready-made layouts via the **Samples** menu: Audio Mixer, Settings Panel, Media Player, Color Picker, Debug Console — plus a full **Pioneer DJ-S11 Mixer** (2 decks, 3-band EQ, crossfader, jog wheels, FX).

### Perform Mode & Web MIDI

- **Perform Mode** (`P` key or ▶ Perform): all placed widgets become live controls — draggable sliders, toggleable buttons, rotatable jog wheels, animated VU meters
- **Web MIDI** (Chrome/Edge/Opera): connect any MIDI controller, **MIDI Learn** (`L` key) to map hardware CC/notes/pitch-bend to on-screen widgets, MIDI OUT feedback for LED rings and motor faders
- **Probe MIDI**: hardware diagnostics with input/output enumeration + live message monitor
- **Numark Mixstream Pro preset** included as a starter mapping (requires Computer/MIDI mode on the unit)

### Built-in Audio Engine (DJ-S11)

The DJ-S11 sample includes a **fully functional audio engine** via Web Audio API:

- **2 royalty-free procedural tracks**, synthesized in-browser (no samples, no licensing, works offline): House 128 BPM on Deck A, Breakbeat 140 BPM on Deck B
- **PLAY / CUE / SYNC** — working sequencer per deck with beatmatching
- **3-band EQ** (HI/MID/LOW) per channel via BiquadFilters
- **Channel faders + equal-power crossfader**
- **Pitch control** ±8% per deck
- **Master volume + real RMS VU meters** (AnalyserNode)
- **Jog wheels** nudge playback phase
- **FX**: ECHO on master (delay + feedback)
- Live playhead progress bars and time counters on each deck

## Upstream Contributions

This fork tracks upstream and contributes back via PRs:

| PR | Issue | Description | Status |
|----|-------|-------------|--------|
| [#9467](https://github.com/ocornut/imgui/pull/9467) | [#8878](https://github.com/ocornut/imgui/issues/8878) | Android CI: Pin Gradle 9.4.1 via setup-gradle action | **MERGED** ✅ |
| [#9473](https://github.com/ocornut/imgui/pull/9473) | [#8802](https://github.com/ocornut/imgui/issues/8802) | OpenGL2/3: Backup and restore GL_UNPACK state in UpdateTexture() | **MERGED** ✅ |
| [#9474](https://github.com/ocornut/imgui/pull/9474) | [#6627](https://github.com/ocornut/imgui/issues/6627) | Android: Clear mouse pos on touch release + SDL2: Restore StartTextInput | **MERGED** ✅ |
| [#9481](https://github.com/ocornut/imgui/pull/9481) | [#8597](https://github.com/ocornut/imgui/issues/8597) | DX9/Allegro5/SDLRenderer2/SDLRenderer3: Defer texture destruction until UnusedFrames > 0 | Open |
| [#9483](https://github.com/ocornut/imgui/pull/9483) | [#6302](https://github.com/ocornut/imgui/issues/6302) | FreeType: Fix CBDT/COLR bitmap color emoji font loading (NotoColorEmoji, Twemoji) | Open |
| [#9468](https://github.com/ocornut/imgui/pull/9468) | [#9390](https://github.com/ocornut/imgui/issues/9390) | Vulkan: Fixed use-after-free in multi-viewport dynamic rendering | Closed |
| [#9469](https://github.com/ocornut/imgui/pull/9469) | [#3446](https://github.com/ocornut/imgui/issues/3446) | Android: Move JNI keyboard/clipboard into backend, add display metrics | Open (under review) |

### Bug Fixes Summary

**Merged upstream (3 PRs):**
- **Android CI** — Pinned Gradle 9.4.1 via `setup-gradle` action, no binaries (#9467)
- **OpenGL GL_UNPACK state** — Backup/restore `GL_UNPACK_ROW_LENGTH` and `GL_UNPACK_ALIGNMENT` in UpdateTexture() to prevent caller state corruption (#9473)
- **Android touch release** — Clear mouse position on `AMOTION_EVENT_ACTION_UP` so buttons don't stay hovered (#6627)
- **SDL2 soft keyboard** — Restored `SDL_StartTextInput()`/`SDL_StopTextInput()` under `#ifdef __ANDROID__` for on-screen keyboard (#7636)

**Open PRs (3):**
- **Texture destroy deferral** — Defer destruction until `UnusedFrames > 0` in DX9, Allegro5, SDLRenderer2, SDLRenderer3, matching all other backends. Addresses multi-threaded rendering safety concerns raised in #8597. (#9481)
- **CBDT/COLR emoji fonts** — Fix FreeType loading of bitmap color emoji fonts (NotoColorEmoji, Twemoji COLR) by detecting CBDT table and using `FT_Select_Size()`. (#9483)
- **Android backend cleanup** — Move JNI keyboard/clipboard into backend, add display metrics, breaking change to Init() signature (#9469, under review by Omar)

## Go Bindings

This fork includes Go (cgo) bindings for Dear ImGui:

```
bindings/go/
  imgui_bridge.h/.cpp  — Thin C API bridge (extern "C")
  imgui.go             — Go cgo wrapper package
  README.md            — Full documentation

examples/example_golang/
  main.go              — Null-mode example
  Makefile             — Build targets
```

## Sync Status

| Branch | Tracks | Status |
|--------|--------|--------|
| `master` | `upstream/master` | Synced |
| `docking` | `upstream/docking` | Synced |

## Quick Start

```bash
# Clone
git clone https://github.com/Turtle-PB/imgui_xl.git
cd imgui_xl

# Build example_null (no backend dependencies needed)
# Windows (MinGW)
mingw32-make -C examples/example_null

# Windows (MSVC)
cd examples\example_null
.\build_win32.bat

# Linux/macOS
cd examples/example_null
make
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on submitting issues, PRs, and syncing with upstream.

## License

MIT License — see [LICENSE.txt](LICENSE.txt) for details.

## Credits

Dear ImGui is developed by [Omar Cornut](https://github.com/ocornut) and every direct or indirect [contributor](https://github.com/ocornut/imgui/graphs/contributors) to the GitHub repository.

Echonation visuals inspired by [@Echonation](https://youtube.com/@echonationmusic). Fork by [Paul Adcock](https://github.com/Turtle-PB) (Turtle-PB).
