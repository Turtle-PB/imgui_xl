# Example: Emscripten Pixel OS Demo

Interactive **Pixel phone simulator** showcasing Pixel OS capabilities, built as a self-contained HTML demo.

## What It Is

A working model of a Pixel 8 Pro phone that runs in any browser. It demonstrates:

- **Lock screen** with live clock, notifications, fingerprint/swipe unlock
- **Home screen** with Material You design, Google search bar, weather widget, app grid
- **7 functional apps**: Phone (keypad), Messages (chat with auto-replies), Settings (toggles), Camera (shutter flash), Photos (gallery), Clock (world clock, alarms, stopwatch), Browser (Chrome mockup)
- **Quick Settings** panel with 8 toggle tiles and brightness slider
- **Recents** app switcher
- **Dark mode** toggle
- **Theme switcher** with 6 color themes
- **Menu bar** integration (File, View, Tools, Help) with keyboard shortcuts

## Running

### Option 1: Open Directly
```bash
# Just open the HTML file in any browser
open pixel_os_demo.html        # macOS
xdg-open pixel_os_demo.html    # Linux
start pixel_os_demo.html       # Windows
```

### Option 2: Local Server
```bash
python -m http.server 8765
# Open http://localhost:8765/pixel_os_demo.html
```

### Option 3: GitHub Pages
Visit the [imgui_xl GitHub Pages](https://turtle-pb.github.io/imgui_xl/examples/example_emscripten_pixel/pixel_os_demo.html) to see it live.

## Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Ctrl+P` | Open Phone |
| `Ctrl+M` | Open Messages |
| `Ctrl+B` | Open Browser |
| `Ctrl+,` | Open Settings |
| `Ctrl+L` | Lock Phone |
| `Ctrl+R` | Reload Demo |
| `Ctrl+Q` | Quick Settings |
| `Ctrl+D` | Toggle Dark Mode |
| `Ctrl+T` | Cycle Theme |
| `Ctrl+Tab` | Recents |
| `Home` / `Esc` | Go Home |
| `F1` | Quick Guide |

## Integration with imgui_xl

This demo serves as a reference for web-based UI patterns that complement Dear ImGui's native C++ rendering:

- **Self-contained**: No build step, no dependencies, no server required
- **Cross-platform**: Works on any modern browser (desktop, mobile, tablet)
- **Interactive**: All buttons, toggles, and gestures are functional
- **Extensible**: Add new apps by extending the `views` pattern

## License

MIT License — (c) 2024 Paul Adcock / OwlLogics

Part of [imgui_xl](https://github.com/Turtle-PB/imgui_xl).
