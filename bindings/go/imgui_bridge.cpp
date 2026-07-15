// imgui_bridge.cpp — Thin C API bridge implementation for Dear ImGui
// Copyright (c) 2024 OwlLogics / imgui_xl
// Licensed under MIT

#include "imgui.h"
#include "imgui_bridge.h"

// Compile the full ImGui implementation here
#define IMGUI_IMPLEMENTATION
#include "imgui.cpp"
#include "imgui_demo.cpp"
#include "imgui_draw.cpp"
#include "imgui_tables.cpp"
#include "imgui_widgets.cpp"

// === Core ===
extern "C" void ig_CreateContext(void) { ImGui::CreateContext(); }
extern "C" void ig_DestroyContext(void) { ImGui::DestroyContext(); }
extern "C" void* ig_GetIO(void) { return (void*)&ImGui::GetIO(); }
extern "C" void* ig_GetStyle(void) { return (void*)&ImGui::GetStyle(); }
extern "C" void ig_NewFrame(void) { ImGui::NewFrame(); }
extern "C" void ig_Render(void) { ImGui::Render(); }
extern "C" void* ig_GetDrawData(void) { return (void*)ImGui::GetDrawData(); }
extern "C" const char* ig_GetVersion(void) { return ImGui::GetVersion(); }

// === Style ===
extern "C" void ig_StyleColorsDark(void) { ImGui::StyleColorsDark(); }
extern "C" void ig_StyleColorsLight(void) { ImGui::StyleColorsLight(); }
extern "C" void ig_StyleColorsClassic(void) { ImGui::StyleColorsClassic(); }

// === Windows ===
extern "C" int ig_Begin(const char* name, int* p_open) {
    return ImGui::Begin(name, p_open ? (bool*)p_open : nullptr) ? 1 : 0;
}
extern "C" void ig_End(void) { ImGui::End(); }
extern "C" int ig_BeginChild(const char* str_id, float w, float h, int border) {
    return ImGui::BeginChild(str_id, ImVec2(w, h), border != 0) ? 1 : 0;
}
extern "C" void ig_EndChild(void) { ImGui::EndChild(); }

// === Widgets ===
extern "C" void ig_Text(const char* fmt) { ImGui::Text("%s", fmt); }
extern "C" void ig_TextColored(float r, float g, float b, float a, const char* fmt) {
    ImGui::TextColored(ImVec4(r, g, b, a), "%s", fmt);
}
extern "C" void ig_BulletText(const char* fmt) { ImGui::BulletText("%s", fmt); }
extern "C" int ig_Button(const char* label, float w, float h) {
    return ImGui::Button(label, ImVec2(w, h)) ? 1 : 0;
}
extern "C" int ig_InvisibleButton(const char* str_id, float w, float h) {
    return ImGui::InvisibleButton(str_id, ImVec2(w, h)) ? 1 : 0;
}
extern "C" int ig_Checkbox(const char* label, int* v) {
    bool b = *v != 0;
    int ret = ImGui::Checkbox(label, &b) ? 1 : 0;
    *v = b ? 1 : 0;
    return ret;
}
extern "C" int ig_SliderFloat(const char* label, float* v, float v_min, float v_max) {
    return ImGui::SliderFloat(label, v, v_min, v_max) ? 1 : 0;
}
extern "C" int ig_SliderInt(const char* label, int* v, int v_min, int v_max) {
    return ImGui::SliderInt(label, v, v_min, v_max) ? 1 : 0;
}
extern "C" int ig_ColorEdit3(const char* label, float* col) {
    return ImGui::ColorEdit3(label, col) ? 1 : 0;
}
extern "C" int ig_ColorEdit4(const char* label, float* col, int flags) {
    return ImGui::ColorEdit4(label, col, flags) ? 1 : 0;
}
extern "C" int ig_InputText(const char* label, char* buf, int buf_size) {
    return ImGui::InputText(label, buf, (size_t)buf_size) ? 1 : 0;
}
extern "C" int ig_Combo(const char* label, int* current_item, const char* items_separated_by_zeros, int items_count) {
    return ImGui::Combo(label, current_item, items_separated_by_zeros, items_count) ? 1 : 0;
}

// === Layout ===
extern "C" void ig_SameLine(float offset, float spacing) { ImGui::SameLine(offset, spacing); }
extern "C" void ig_Separator(void) { ImGui::Separator(); }
extern "C" void ig_Spacing(void) { ImGui::Spacing(); }
extern "C" void ig_NewLine(void) { ImGui::NewLine(); }
extern "C" void ig_PushItemWidth(float w) { ImGui::PushItemWidth(w); }
extern "C" void ig_PopItemWidth(void) { ImGui::PopItemWidth(); }

// === Drawing ===
extern "C" void ig_ShowDemoWindow(int* p_open) { ImGui::ShowDemoWindow(p_open ? (bool*)p_open : nullptr); }

// === Metrics / Info ===
extern "C" float ig_GetFramerate(void) { return ImGui::GetIO().Framerate; }
extern "C" float ig_GetTime(void) { return (float)ImGui::GetTime(); }
extern "C" int ig_GetFrameCount(void) { return ImGui::GetFrameCount(); }

// === Display ===
extern "C" void ig_GetDisplaySize(float* w, float* h) {
    ImVec2 size = ImGui::GetIO().DisplaySize;
    *w = size.x; *h = size.y;
}
extern "C" void ig_GetDisplayFramebufferScale(float* x, float* y) {
    ImVec2 scale = ImGui::GetIO().DisplayFramebufferScale;
    *x = scale.x; *y = scale.y;
}
