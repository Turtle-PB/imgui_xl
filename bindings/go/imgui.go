// Package imgui provides Go bindings for Dear ImGui via cgo.
// It uses a thin C bridge (imgui_bridge.h/.cpp) to call the C++ ImGui API.
//
// Requirements:
//   - C++ compiler (g++ or clang++) for building the bridge
//   - CGO_ENABLED=1
//
// Build:
//   CGO_ENABLED=1 go build
//
// Copyright (c) 2024 OwlLogics / imgui_xl
// Licensed under MIT

package imgui

/*
#cgo CXXFLAGS: -std=c++11 -I${SRCDIR}/../../.. -I${SRCDIR}/../../../backends
#cgo LDFLAGS: -lstdc++ -lm

// Include the C bridge header
#include "imgui_bridge.h"

// We need to compile the bridge .cpp via cgo's CXX support.
// The bridge .cpp includes all of ImGui's implementation.
*/
import "C"
import (
	"unsafe"
)

// Version returns the compiled ImGui version string.
func Version() string {
	return C.GoString(C.ig_GetVersion())
}

// CreateContext creates the ImGui context.
func CreateContext() {
	C.ig_CreateContext()
}

// DestroyContext destroys the current ImGui context.
func DestroyContext() {
	C.ig_DestroyContext()
}

// NewFrame starts a new ImGui frame.
func NewFrame() {
	C.ig_NewFrame()
}

// Render finalizes the frame and prepares draw data.
func Render() {
	C.ig_Render()
}

// StyleColorsDark sets the dark color scheme.
func StyleColorsDark() {
	C.ig_StyleColorsDark()
}

// StyleColorsLight sets the light color scheme.
func StyleColorsLight() {
	C.ig_StyleColorsLight()
}

// StyleColorsClassic sets the classic color scheme.
func StyleColorsClassic() {
	C.ig_StyleColorsClassic()
}

// ShowDemoWindow shows the ImGui demo window.
func ShowDemoWindow(open *bool) {
	if open != nil {
		v := 0
		if *open {
			v = 1
		}
		C.ig_ShowDemoWindow(&v)
		*open = v != 0
	} else {
		C.ig_ShowDemoWindow(nil)
	}
}

// GetFramerate returns the current FPS.
func GetFramerate() float32 {
	return float32(C.ig_GetFramerate())
}

// GetTime returns the elapsed time since initialization.
func GetTime() float32 {
	return float32(C.ig_GetTime())
}

// GetFrameCount returns the current frame count.
func GetFrameCount() int {
	return int(C.ig_GetFrameCount())
}

// GetDisplaySize returns the display width and height.
func GetDisplaySize() (float32, float32) {
	var w, h C.float
	C.ig_GetDisplaySize(&w, &h)
	return float32(w), float32(h)
}

// --- Window functions ---

// Begin pushes a new window to the stack.
// Returns true if the window is visible (not collapsed).
func Begin(name string, open *bool) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var pOpen *C.int
	if open != nil {
		v := 0
		if *open {
			v = 1
		}
		pOpen = &v
		defer func() { *open = v != 0 }()
	}
	return C.ig_Begin(cName, pOpen) != 0
}

// End closes the current window.
func End() {
	C.ig_End()
}

// SameLine places the next widget on the same line.
func SameLine(offset, spacing float32) {
	C.ig_SameLine(C.float(offset), C.float(spacing))
}

// Separator draws a horizontal separator.
func Separator() {
	C.ig_Separator()
}

// Spacing adds vertical spacing.
func Spacing() {
	C.ig_Spacing()
}

// --- Widget functions ---

// Text displays a text label.
func Text(s string) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	C.ig_Text(cStr)
}

// TextColored displays colored text.
func TextColored(r, g, b, a float32, s string) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	C.ig_TextColored(C.float(r), C.float(g), C.float(b), C.float(a), cStr)
}

// Button displays a button. Returns true if clicked.
func Button(label string, width, height float32) bool {
	cStr := C.CString(label)
	defer C.free(unsafe.Pointer(cStr))
	return C.ig_Button(cStr, C.float(width), C.float(height)) != 0
}

// Checkbox displays a checkbox.
func Checkbox(label string, v *bool) bool {
	cStr := C.CString(label)
	defer C.free(unsafe.Pointer(cStr))
	var vC C.int
	if *v {
		vC = 1
	}
	ret := C.ig_Checkbox(cStr, &vC) != 0
	*v = vC != 0
	return ret
}

// SliderFloat displays a float slider.
func SliderFloat(label string, v *float32, vMin, vMax float32) bool {
	cStr := C.CString(label)
	defer C.free(unsafe.Pointer(cStr))
	return C.ig_SliderFloat(cStr, (*C.float)(v), C.float(vMin), C.float(vMax)) != 0
}

// SliderInt displays an integer slider.
func SliderInt(label string, v *int, vMin, vMax int) bool {
	cStr := C.CString(label)
	defer C.free(unsafe.Pointer(cStr))
	return C.ig_SliderInt(cStr, (*C.int)(v), C.int(vMin), C.int(vMax)) != 0
}

// ColorEdit3 displays a 3-component color editor.
func ColorEdit3(label string, col *[3]float32) bool {
	cStr := C.CString(label)
	defer C.free(unsafe.Pointer(cStr))
	return C.ig_ColorEdit3(cStr, (*C.float)(unsafe.Pointer(col))) != 0
}

// InputText displays a text input field.
func InputText(label string, buf *string, bufSize int) bool {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	// Allocate a C buffer
	cBuf := C.malloc(C.size_t(bufSize))
	defer C.free(cBuf)
	C.memcpy(cBuf, unsafe.Pointer(C.CString(*buf)), C.size_t(len(*buf)+1))
	ret := C.ig_InputText(cLabel, (*C.char)(cBuf), C.int(bufSize)) != 0
	*buf = C.GoString((*C.char)(cBuf))
	return ret
}

// Combo displays a dropdown combo box.
func Combo(label string, current *int, items string, itemCount int) bool {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	cItems := C.CString(items)
	defer C.free(unsafe.Pointer(cItems))
	return C.ig_Combo(cLabel, (*C.int)(current), cItems, C.int(itemCount)) != 0
}

// BulletText displays a bulleted text item.
func BulletText(s string) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	C.ig_BulletText(cStr)
}

// NewLine inserts a new line.
func NewLine() {
	C.ig_NewLine()
}

// PushItemWidth sets the width for the next widget.
func PushItemWidth(width float32) {
	C.ig_PushItemWidth(C.float(width))
}

// PopItemWidth restores the previous item width.
func PopItemWidth() {
	C.ig_PopItemWidth()
}
