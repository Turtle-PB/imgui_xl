# Makefile for the Go ImGui example
# Requirements: g++ (or clang++), Go 1.21+, CGO_ENABLED=1
#
# Usage:
#   make build    - Compile the Go example
#   make run      - Build and run
#   make clean    - Remove build artifacts
#
# Copyright (c) 2024 OwlLogics / imgui_xl
# Licensed under MIT

GO = go
CXX = g++
IMGUI_DIR = ../..
BRIDGE_DIR = $(IMGUI_DIR)/bindings/go

CGO_ENABLED = 1
GOFLAGS = -trimpath
OUTPUT = imgui_go_example

.PHONY: all build run clean vet fmt

all: build

build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build -o $(OUTPUT) . 2>&1 || echo "Build requires a C++ compiler (g++ or clang++) on PATH"

run: build
	./$(OUTPUT)

vet:
	$(GO) vet ./...

fmt:
	$(GO) fmt ./...

clean:
	rm -f $(OUTPUT) *.o *.exe
