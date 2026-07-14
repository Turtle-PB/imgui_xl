# imgui_xl

**Dear ImGui XL** — An enhanced community fork of [Dear ImGui](https://github.com/ocornut/imgui), focused on backend improvements, CI/CD fixes, sensors, and extended examples.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Upstream](https://img.shields.io/badge/upstream-ocornut%2Fimgui-blue)](https://github.com/ocornut/imgui)
[![Live Demo](https://img.shields.io/badge/demo-GitHub%20Pages-green)](https://turtle-pb.github.io/imgui_xl/)
[![PRs](https://img.shields.io/badge/PRs-3%20open-orange)](https://github.com/ocornut/imgui/pulls?q=is%3Aopen+author%3ATurtle-PB)

> Dear ImGui is a bloat-free graphical user interface library for C++. It outputs optimized vertex buffers that you can render anytime in your 3D-pipeline-enabled application. It is fast, portable, renderer-agnostic, and self-contained (no external dependencies).

## Live Demos

| Demo | Description | Link |
|------|-------------|------|
| **Showcase** | Landing page with all demos + PR list | [turtle-pb.github.io/imgui_xl](https://turtle-pb.github.io/imgui_xl/) |
| **Echonation Visual** | 5-mode audio-reactive canvas demo | [echonation.html](https://turtle-pb.github.io/imgui_xl/echonation.html) |
| **Pixel OS Model** | Interactive Pixel 8 Pro phone simulator | [pixel_os_demo.html](https://turtle-pb.github.io/imgui_xl/pixel_os_demo.html) |

## Upstream Contributions

This fork tracks upstream and contributes back via PRs:

| PR | Issue | Description | Status |
|----|-------|-------------|--------|
| [#9467](https://github.com/ocornut/imgui/pull/9467) | [#8878](https://github.com/ocornut/imgui/issues/8878) | Android CI: Pin Gradle 9.4.1 via setup-gradle action (no binaries) | Open |
| [#9468](https://github.com/ocornut/imgui/pull/9468) | [#9390](https://github.com/ocornut/imgui/issues/9390) | Vulkan: Fixed use-after-free in multi-viewport dynamic rendering | Open |
| [#9469](https://github.com/ocornut/imgui/pull/9469) | [#3446](https://github.com/ocornut/imgui/issues/3446) | Android: Move JNI keyboard/clipboard into backend, add sensors + display | Open |

### Backend Improvements (PR #9469)

- **JNI keyboard/clipboard** — Moved ~120 lines of JNI boilerplate from the example into the backend. App code is now just `Init → Loop → Shutdown`.
- **8 sensor types** — Accelerometer, gyroscope, magnetometer, light, proximity, pressure, humidity, ambient temperature via NDK ASensor API (no JNI).
- **Display metrics** — DPI, density, xdpi/ydpi, refresh rate, orientation queried via JNI. Auto-scales ImGui style to actual device density.
- **Clipboard** — Full clipboard support via JNI to Android ClipboardManager.
- **Zero binaries** — No jar, no wrapper, no .gitignore changes.

### CI Fix (PR #9467)

- Uncommented and updated `setup-gradle@v6` action to pin Gradle 9.4.1 (required by AGP 9.2.0).
- Single 4-line change to `build.yml`, no binaries added.

### Vulkan Fix (PR #9468)

- Deep-copy `SurfaceFormat.format` into persistent buffer in `ImGui_ImplVulkan_CreateWindow()` to prevent use-after-free when viewport is destroyed.
- Follows the same deep-copy pattern already used in `CreateMainPipeline()`.

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
