// Copyright (c) 2018 Beta Kuang
//
// This software is provided 'as-is', without any express or implied
// warranty.  In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.

package nanovgo

/*
#cgo darwin LDFLAGS: -framework OpenGL
#define NANOVG_GL3_IMPLEMENTATION
#include <OpenGL/gl3.h>
#include "nanovg/src/nanovg.h"
#include "nanovg/src/nanovg_gl.h"
#include "nanovg/src/nanovg_gl_utils.h"
*/
import "C"

// CreateFlag is a flag for creating NanoVGo contexts.
type CreateFlag int

// Create flags.
const (
	// Antialias indicates if geometry based anti-aliasing is used (may not be
	// needed when using MSAA).
	Antialias CreateFlag = C.NVG_ANTIALIAS
	// StencilStrokes indicates if strokes should be drawn using stencil buffer.
	// The rendering will be a little slower, but path overlaps (i.e.
	// self-intersecting or sharp turns) will be drawn just once.
	StencilStrokes CreateFlag = C.NVG_STENCIL_STROKES
	// Debug indicates that additional debug checks are done.
	Debug CreateFlag = C.NVG_DEBUG
)

// CreateContext creates a NanoVGo context for OpenGL 3. flags should be a
// combination of Antialias, StencilStrokes and Debug.
func CreateContext(flags CreateFlag) *Context {
	return (*Context)(C.nvgCreateGL3(C.int(flags)))
}

// Delete deletes a NanoVGo context.
func (ctx *Context) Delete() {
	C.nvgDeleteGL3((*C.NVGcontext)(ctx))
}

// Color is an RGBA color, each color channel represented as a float32.
type Color C.NVGcolor

func (color Color) c() C.NVGcolor {
	return C.NVGcolor(color)
}

// Paint is a paint style used for painting.
type Paint struct {
	XForm      [6]float32
	Extent     [2]float32
	Radius     float32
	Feather    float32
	InnerColor Color
	OuterColor Color
	Image      int
}

func (paint Paint) c() C.NVGpaint {
	cPaint := C.NVGpaint{
		radius:     C.float(paint.Radius),
		feather:    C.float(paint.Feather),
		innerColor: paint.InnerColor.c(),
		outerColor: paint.OuterColor.c(),
		image:      C.int(paint.Image),
	}
	for i := 0; i < 6; i++ {
		cPaint.xform[i] = C.float(paint.XForm[i])
	}
	for i := 0; i < 2; i++ {
		cPaint.extent[i] = C.float(paint.Extent[i])
	}
	return cPaint
}

// Winding specifies the direction of winding.
type Winding int

// Windings.
const (
	// CCW is the winding for solid shapes.
	CCW Winding = C.NVG_CCW
	// CW is the winding for holes.
	CW Winding = C.NVG_CW
)

// Solidity specifies whether a shape is solid has a hole.
type Solidity int

// Solidities.
const (
	// Solid is for CCW.
	Solid Solidity = C.NVG_SOLID
	// Hole is for CW.
	Hole Solidity = C.NVG_HOLE
)

// LineCap specifies the style of line caps.
type LineCap int

// Line caps.
const (
	Butt   LineCap = C.NVG_BUTT
	Round  LineCap = C.NVG_ROUND
	Square LineCap = C.NVG_SQUARE
	Bevel  LineCap = C.NVG_BEVEL
	Miter  LineCap = C.NVG_MITER
)

// Align indicates how text should be aligned horizontally or vertically.
type Align int

// Align styles.
const (
	// Horizontal aligns.
	AlignLeft   Align = C.NVG_ALIGN_LEFT
	AlignCenter Align = C.NVG_ALIGN_CENTER
	AlignRight  Align = C.NVG_ALIGN_RIGHT
	// Vertical align.
	AlignTop      Align = C.NVG_ALIGN_TOP
	AlignMiddle   Align = C.NVG_ALIGN_MIDDLE
	AlignBottom   Align = C.NVG_ALIGN_BOTTOM
	AlignBaseline Align = C.NVG_ALIGN_BASELINE
)

// BlendFactor indicates how blending should work.
type BlendFactor int

// Blend factors.
const (
	Zero             BlendFactor = C.NVG_ZERO
	One              BlendFactor = C.NVG_ONE
	SrcColor         BlendFactor = C.NVG_SRC_COLOR
	OneMinusSrcColor BlendFactor = C.NVG_ONE_MINUS_SRC_COLOR
	DstColor         BlendFactor = C.NVG_DST_COLOR
	OneMinusDstColor BlendFactor = C.NVG_ONE_MINUS_DST_COLOR
	SrcAlpha         BlendFactor = C.NVG_SRC_ALPHA
	OneMinusSrcAlpha BlendFactor = C.NVG_ONE_MINUS_SRC_ALPHA
	DstAlpha         BlendFactor = C.NVG_DST_ALPHA
	OneMinusDstAlpha BlendFactor = C.NVG_ONE_MINUS_DST_ALPHA
	SrcAlphaSaturate BlendFactor = C.NVG_SRC_ALPHA_SATURATE
)

// CompositeOperation specify the operation for composition.
type CompositeOperation int

// Composite operations.
const (
	SourceOver      CompositeOperation = C.NVG_SOURCE_OVER
	SourceIn        CompositeOperation = C.NVG_SOURCE_IN
	SourceOut       CompositeOperation = C.NVG_SOURCE_OUT
	Atop            CompositeOperation = C.NVG_ATOP
	DestinationOver CompositeOperation = C.NVG_DESTINATION_OVER
	DestinationIn   CompositeOperation = C.NVG_DESTINATION_IN
	DestinationOut  CompositeOperation = C.NVG_DESTINATION_OUT
	DestinationAtop CompositeOperation = C.NVG_DESTINATION_ATOP
	Lighter         CompositeOperation = C.NVG_LIGHTER
	Copy            CompositeOperation = C.NVG_COPY
	Xor             CompositeOperation = C.NVG_XOR
)

// CompositeOperationState records the state of a composition operation.
type CompositeOperationState C.NVGcompositeOperationState

func (state CompositeOperationState) c() C.NVGcompositeOperationState {
	return C.NVGcompositeOperationState(state)
}

// GlyphPosition is the position of a glyph from an input string.
type GlyphPosition C.NVGglyphPosition

func (pos GlyphPosition) c() C.NVGglyphPosition {
	return C.NVGglyphPosition(pos)
}

// TextRow stores the range, width and position of a row of text.
type TextRow C.NVGtextRow

func (row TextRow) c() C.NVGtextRow {
	return C.NVGtextRow(row)
}

// ImageFlag indicates how images should be processed.
type ImageFlag int

// Image flags.
const (
	ImageGenerateMipmaps ImageFlag = C.NVG_IMAGE_GENERATE_MIPMAPS
	ImageRepeatX         ImageFlag = C.NVG_IMAGE_REPEATX
	ImageRepeatY         ImageFlag = C.NVG_IMAGE_REPEATY
	ImageFlipy           ImageFlag = C.NVG_IMAGE_FLIPY
	ImagePremultiplied   ImageFlag = C.NVG_IMAGE_PREMULTIPLIED
	ImageNearest         ImageFlag = C.NVG_IMAGE_NEAREST
)

// Context is a NanoVGo context for vector graphics rendering.
type Context C.NVGcontext

func (ctx *Context) c() *C.NVGcontext {
	return (*C.NVGcontext)(ctx)
}

// BeginFrame begins drawing a new frame.
//
// Calls to NanoVGo drawing API should be wrapped in Context.BeginFrame() and
// Context.EndFrame(). Context.BeginFrame() defines the size of the window to
// render to in relation to currently set viewport (i.e. glViewport on GL
// backends). Device pixel ratio allows to control the rendering on Hi-DPI
// devices.
//
// For example, GLFW returns two dimensions for an opened window: window size
// and framebuffer size. In that case you would set windowWidth/Height to the
// window size, deivcePixelRatio to frameBufferWidth / windowWidth.
func (ctx *Context) BeginFrame(windowWidth, windowHeight, devicePixelRatio float32) {
	C.nvgBeginFrame(ctx.c(), C.float(windowWidth), C.float(windowHeight), C.float(devicePixelRatio))
}

// CancelFrame cancels drawing the current frame.
func (ctx *Context) CancelFrame() {
	C.nvgCancelFrame(ctx.c())
}

// EndFrame ends drawing and flushes remaining render state.
func (ctx *Context) EndFrame() {
	C.nvgEndFrame(ctx.c())
}

// Composite operations.
//
// The composite operations in NanoVGo are modeled after HTML Canvas API, and
// the blend func is based on OpenGL (see corresponding manuals for more info).
// The colors in the blending state have premultiplied alpha.

// GlobalCompositeOperation sets the composite operation. op should be one of
// CompositeOperation.
func (ctx *Context) GlobalCompositeOperation(op CompositeOperation) {
	C.nvgGlobalCompositeOperation(ctx.c(), C.int(op))
}

// GlobalCompositeBlendFunc sets the composite operation with custom pixel
// arithmetic. sFactor and dFactor should be one of BlendFactor.
func (ctx *Context) GlobalCompositeBlendFunc(sFactor, dFactor BlendFactor) {
	C.nvgGlobalCompositeBlendFunc(ctx.c(), C.int(sFactor), C.int(dFactor))
}

// GlobalCompositeBlendFuncSeparate sets the composite operation with custom
// pixel arithmetic for RGB and alpha components separately. The parameters
// should be one of BlendFactor.
func (ctx *Context) GlobalCompositeBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha BlendFactor) {
	C.nvgGlobalCompositeBlendFuncSeparate(ctx.c(), C.int(srcRGB), C.int(dstRGB), C.int(srcAlpha), C.int(dstAlpha))
}

// Color utils.
//
// Colors in NanoVGo are stored as unsigned ints in ABGR format.

// RGB returns a color value from red, green and blue values. Alpha will be set
// to 255 (1.0f).
func RGB(r, g, b uint8) Color {
	return Color(C.nvgRGB(C.uchar(r), C.uchar(g), C.uchar(b)))
}

// RGBf returns a color value from red, green and blue values. Alpha will be set
// to 1.0f.
func RGBf(r, g, b float32) Color {
	return Color(C.nvgRGBf(C.float(r), C.float(g), C.float(b)))
}

// RGBA returns a color value from red, green, blue and alpha values.
func RGBA(r, g, b, a uint8) Color {
	return Color(C.nvgRGBA(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}

// RGBAf returns a color value from red, green, blue and alpha values.
func RGBAf(r, g, b, a float32) Color {
	return Color(C.nvgRGBAf(C.float(r), C.float(g), C.float(b), C.float(a)))
}

// LerpRGBA linearly interpolates from c0 to c1, and returns resulting color
// value.
func LerpRGBA(c0, c1 Color, u float32) Color {
	return Color(C.nvgLerpRGBA(c0.c(), c1.c(), C.float(u)))
}

// TransRGBA sets transparency of c, and returns the resulting color value.
func TransRGBA(c Color, a uint8) Color {
	return Color(C.nvgTransRGBA(c.c(), C.uchar(a)))
}

// TransRGBAf sets transparency of c, and returns the resulting color value.
func TransRGBAf(c Color, a float32) Color {
	return Color(C.nvgTransRGBAf(c.c(), C.float(a)))
}

// HSL returns a color value specified by hue, saturation and lightness.
//
// HSL values are all in range [0..1], alpha will be set to 255.
func HSL(h, s, l float32) Color {
	return Color(C.nvgHSL(C.float(h), C.float(s), C.float(l)))
}

// HSLA returns a color value specified by hue, saturation, lightness and alpha.
//
// HSL values are all in range [0..1], alpah in range [0..255].
func HSLA(h, s, l float32, a uint8) Color {
	return Color(C.nvgHSLA(C.float(h), C.float(s), C.float(l), C.uchar(a)))
}
