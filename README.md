# imgui_xl

**Dear ImGui XL** — An enhanced community fork of [Dear ImGui](https://github.com/ocornut/imgui), focused on backend improvements, CI/CD fixes, and extended examples.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Upstream](https://img.shields.io/badge/upstream-ocornut%2Fimgui-blue)](https://github.com/ocornut/imgui)
[![Branch: master](https://img.shields.io/badge/branch-master-green)](https://github.com/Turtle-PB/imgui_xl/tree/master)

> Dear ImGui is a bloat-free graphical user interface library for C++. It outputs optimized vertex buffers that you can render anytime in your 3D-pipeline-enabled application. It is fast, portable, renderer-agnostic, and self-contained (no external dependencies).

## Why This Fork?

This fork (`imgui_xl`) serves as a staging ground for community contributions to Dear ImGui. It tracks upstream `master` and `docking` branches closely and adds:

- **Backend fixes** — Vulkan use-after-free fix for multi-viewport dynamic rendering (#9390)
- **CI/CD improvements** — Android build fixed via pinned Gradle wrapper (#8878)
- **Enhanced metadata** — Topics, issues, and discussions enabled for community engagement
- **Contributing guide** — Clear path for new contributors

## Sync Status

This fork is kept in sync with [ocornut/imgui](https://github.com/ocornut/imgui):

| Branch | Tracks | Status |
|--------|--------|--------|
| `master` | `upstream/master` | ✅ Synced |
| `docking` | `upstream/docking` | ✅ Synced |

## Active Contribution Branches

| Branch | Issue | Description |
|--------|-------|-------------|
| `fix/vulkan-dynamic-rendering-use-after-free` | [#9390](https://github.com/ocornut/imgui/issues/9390) | Fix use-after-free in Vulkan multi-viewport dynamic rendering |
| `fix/android-ci-gradle-wrapper` | [#8878](https://github.com/ocornut/imgui/issues/8878) | Pin Gradle 9.4.1 via wrapper for AGP 9.2.0 compatibility |

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

## Web Demo: Pixel OS Model

This fork includes an interactive **Pixel OS Model Demo** — a working Pixel 8 Pro phone simulator that showcases Pixel OS capabilities in any browser:

- Lock screen, home screen, 7 functional apps (Phone, Messages, Camera, Settings, Clock, Photos, Browser)
- Quick Settings, Recents, Dark Mode, Theme switcher
- Full menu bar with keyboard shortcuts
- 100% self-contained HTML — no build step, no dependencies

**Run it:** Open [`examples/example_emscripten_pixel/pixel_os_demo.html`](examples/example_emscripten_pixel/pixel_os_demo.html) in any browser, or visit the [live demo on GitHub Pages](https://turtle-pb.github.io/imgui_xl/examples/example_emscripten_pixel/pixel_os_demo.html).

## Upstream Resources

- **Dear ImGui Homepage**: [github.com/ocornut/imgui](https://github.com/ocornut/imgui)
- **Documentation**: [docs/](docs/) directory
- **FAQ**: [docs/FAQ.md](docs/FAQ.md)
- **Wiki**: [github.com/ocornut/imgui/wiki](https://github.com/ocornut/imgui/wiki)
- **Issues**: [github.com/ocornut/imgui/issues](https://github.com/ocornut/imgui/issues)

## License

MIT License — see [LICENSE.txt](LICENSE.txt) for details.

## Credits

Dear ImGui is developed by [Omar Cornut](https://github.com/ocornut) and every direct or indirect [contributor](https://github.com/ocornut/imgui/graphs/contributors) to the GitHub repository.
