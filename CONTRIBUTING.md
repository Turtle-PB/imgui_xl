# Contributing to imgui_xl

First and foremost: **Dear ImGui** is developed by [Omar Cornut](https://github.com/ocornut) and the community at [ocornut/imgui](https://github.com/ocornut/imgui). This fork (`imgui_xl`) exists to facilitate community contributions, test improvements, and develop backend fixes before submitting upstream PRs.

## How to Contribute

### Reporting Issues

1. **Check upstream first** — Most issues belong in [ocornut/imgui/issues](https://github.com/ocornut/imgui/issues). Only open issues here if they're specific to this fork's modifications.
2. **Use the issue tracker** — Provide a Minimal, Complete, and Verifiable Example (MCVE).
3. **Include version info** — Branch (master/docking), commit hash, backend, compiler, OS.

### Submitting Pull Requests

1. **Fork this repo** and create a feature branch from `master` or `docking` (match the upstream branch you're targeting).
2. **Keep PRs focused** — One fix/feature per PR. Small PRs are reviewed faster.
3. **Follow the existing code style** — Dear ImGui has a specific style. Read the [upstream FAQ](https://github.com/ocornut/imgui/blob/master/docs/FAQ.md) and [CONTRIBUTING.md](https://github.com/ocornut/imgui/blob/master/docs/CONTRIBUTING.md).
4. **Test your change** — At minimum, build `example_null` to verify compilation:
   ```bash
   # Windows (MinGW)
   mingw32-make -C examples/example_null

   # Windows (MSVC)
   cd examples\example_null
   call "C:\Program Files\Microsoft Visual Studio\...\VC\Auxiliary\Build\vcvars64.bat"
   .\build_win32.bat
   ```
5. **Write a clear commit message** — Use the format: `Area: Description. (#issue)` (e.g., `Vulkan: Fixed use-after-free in multi-viewport dynamic rendering. (#9390)`)

### Areas Needing Help

Check the [upstream "help wanted" issues](https://github.com/ocornut/imgui/issues?q=is%3Aopen+label%3A%22help+wanted%22) for areas where contributions are needed. Current focus areas in this fork:

- **Backend improvements** — Vulkan, DirectX, Metal backends
- **CI/CD fixes** — Android build, cross-platform testing
- **Documentation** — Examples, tutorials, backend guides

### Syncing with Upstream

This fork tracks `ocornut/imgui:master` and `ocornut/imgui:docking`. To sync:

```bash
git remote add upstream https://github.com/ocornut/imgui.git
git fetch upstream
git checkout master
git merge upstream/master
git push origin master
```

### Branch Strategy

| Branch | Tracks | Purpose |
|--------|--------|---------|
| `master` | `upstream/master` | Stable, synced with upstream master |
| `docking` | `upstream/docking` | Synced with upstream docking branch |
| `fix/*` | — | Bug fix branches for PRs |
| `enhance/*` | — | Enhancement/feature branches |

## License

All contributions are subject to the [MIT License](LICENSE.txt), same as upstream Dear ImGui.

## Code of Conduct

Be respectful, constructive, and patient. This is a community effort.
