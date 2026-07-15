// imgui_bridge.h — Thin C API bridge for Dear ImGui
// This file provides C-callable wrappers around the C++ Dear ImGui API.
// It allows languages like Go (via cgo), Rust, Python, etc. to use ImGui.
//
// Compile with: g++ -c imgui_bridge.cpp -I.. -I../backends -DIMGUI_IMPLEMENTATION
// The resulting .o can be linked from any C-compatible FFI.
//
// Copyright (c) 2024 OwlLogics / imgui_xl
// Licensed under MIT

#pragma once

#ifdef __cplusplus
extern "C" {
#endif

// === Core ===
void        ig_CreateContext(void);
void        ig_DestroyContext(void);
void*       ig_GetIO(void);             // Returns ImGuiIO* as void*
void*       ig_GetStyle(void);           // Returns ImGuiStyle* as void*
void        ig_NewFrame(void);
void        ig_Render(void);
void*       ig_GetDrawData(void);        // Returns ImDrawData* as void*
const char* ig_GetVersion(void);

// === Style ===
void        ig_StyleColorsDark(void);
void        ig_StyleColorsLight(void);
void        ig_StyleColorsClassic(void);

// === Windows ===
int         ig_Begin(const char* name, int* p_open);  // Returns bool (1=visible)
void        ig_End(void);
int         ig_BeginChild(const char* str_id, float w, float h, int border);  // bool
void        ig_EndChild(void);

// === Widgets ===
void        ig_Text(const char* fmt);
void        ig_TextColored(float r, float g, float b, float a, const char* fmt);
void        ig_BulletText(const char* fmt);
int         ig_Button(const char* label, float w, float h);  // bool
int         ig_InvisibleButton(const char* str_id, float w, float h);  // bool
int         ig_Checkbox(const char* label, int* v);  // bool, v is bool*
int         ig_SliderFloat(const char* label, float* v, float v_min, float v_max);  // bool
int         ig_SliderInt(const char* label, int* v, int v_min, int v_max);  // bool
int         ig_ColorEdit3(const char* label, float* col);  // bool, col is float[3]
int         ig_ColorEdit4(const char* label, float* col, int flags);  // bool, col is float[4]
int         ig_InputText(const char* label, char* buf, int buf_size);  // bool
int         ig_Combo(const char* label, int* current_item, const char* items_separated_by_zeros, int items_count);  // bool

// === Layout ===
void        ig_SameLine(float offset, float spacing);
void        ig_Separator(void);
void        ig_Spacing(void);
void        ig_NewLine(void);
void        ig_PushItemWidth(float w);
void        ig_PopItemWidth(void);

// === Drawing ===
void        ig_ShowDemoWindow(int* p_open);

// === Metrics / Info ===
float       ig_GetFramerate(void);
float       ig_GetTime(void);
int         ig_GetFrameCount(void);

// === Display ===
void        ig_GetDisplaySize(float* w, float* h);
void        ig_GetDisplayFramebufferScale(float* x, float* y);

#ifdef __cplusplus
}
#endif
