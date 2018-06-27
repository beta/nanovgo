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

// Package nanovgo is a Go binding for NanoVG
// (https://github.com/memononen/nanovg), a small antialiased vector graphics
// rendering library for OpenGL.
package nanovgo

/*
#cgo darwin LDFLAGS: -framework OpenGL
#define NANOVG_GL3_IMPLEMENTATION
#include <stdlib.h>
#include <OpenGL/gl3.h>
#include "nanovg/src/nanovg.h"
#include "nanovg/src/nanovg_gl.h"
#include "nanovg/src/nanovg_gl_utils.h"

static float color_r(NVGcolor color) {
	return color.r;
}

static float color_g(NVGcolor color) {
	return color.g;
}

static float color_b(NVGcolor color) {
	return color.b;
}

static float color_a(NVGcolor color) {
	return color.a;
}
*/
import "C"
import (
	"unsafe"
)

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

// R returns the red value of color.
func (color Color) R() float32 {
	return float32(C.color_r(color.c()))
}

// G returns the green value of color.
func (color Color) G() float32 {
	return float32(C.color_g(color.c()))
}

// B returns the blue value of color.
func (color Color) B() float32 {
	return float32(C.color_b(color.c()))
}

// A returns the alpha value of color.
func (color Color) A() float32 {
	return float32(C.color_a(color.c()))
}

// Paint is a paint style used for painting.
type Paint C.NVGpaint

func (paint Paint) c() C.NVGpaint {
	return C.NVGpaint(paint)
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

// Solidities.
const (
	// Solid is for CCW.
	Solid Winding = C.NVG_SOLID
	// Hole is for CW.
	Hole Winding = C.NVG_HOLE
)

// LineCap specifies the style of line caps.
type LineCap int

// Line caps.
const (
	Butt     LineCap = C.NVG_BUTT
	RoundCap LineCap = C.NVG_ROUND
	Square   LineCap = C.NVG_SQUARE
)

// LineJoin specifies the style of sharp path corners.
type LineJoin int

// Line joins.
const (
	Miter     LineJoin = C.NVG_MITER
	RoundJoin LineJoin = C.NVG_ROUND
	Bevel     LineJoin = C.NVG_BEVEL
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

// X returns the x-coordinate of the logical glyph position.
func (pos GlyphPosition) X() float32 {
	return float32(pos.c().x)
}

// MinX returns the bounds of the glyph shape.
func (pos GlyphPosition) MinX() float32 {
	return float32(pos.c().minx)
}

// MaxX returns the bounds of the glyph shape.
func (pos GlyphPosition) MaxX() float32 {
	return float32(pos.c().maxx)
}

// TextRow stores the range, width and position of a row of text.
type TextRow C.NVGtextRow

func (row TextRow) c() C.NVGtextRow {
	return C.NVGtextRow(row)
}

// Width returns the logical width of row.
func (row TextRow) Width() float32 {
	return float32(row.c().width)
}

// MinX returns the actual bounds of row. Logical width and bounds can differ
// because of kerning and some parts over extending.
func (row TextRow) MinX() float32 {
	return float32(row.c().minx)
}

// MaxX returns the actual bounds of row. Logical width and bounds can differ
// because of kerning and some parts over extending.
func (row TextRow) MaxX() float32 {
	return float32(row.c().maxx)
}

// Text returns the text of row.
func (row TextRow) Text() string {
	return C.GoStringN(row.c().start, C.int(uintptr(unsafe.Pointer(row.c().end))-uintptr(unsafe.Pointer(row.c().start))))
}

// Next returns the remaining text.
func (row TextRow) Next() string {
	return C.GoString(row.c().next)
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

// State handling.
//
// NanoVGo contains states which represent how paths will be rendered. The state
// contains transform, fill and stroke styles, text and font styles, and scissor
// clipping.

// Save pushes and saves the current render state into a state stack.
//
// A matching Context.Restore() must be used to restore the state.
func (ctx *Context) Save() {
	C.nvgSave(ctx.c())
}

// Restore pops and restores the current render state.
func (ctx *Context) Restore() {
	C.nvgRestore(ctx.c())
}

// Reset resets current render state to default values. This does not affect the
// render state stack.
func (ctx *Context) Reset() {
	C.nvgReset(ctx.c())
}

// Render styles.
//
// Fill and stroke render style can be either a solid color or a paint which is
// a gradient or a pattern. Solid color is simply defined as a color value,
// different kinds of paints can be created using Context.LinearGradient(),
// Context.BoxGradient(), Context.RadialGradient() and Context.ImagePattern().
//
// Current render style can be saved and restored using Context.Save() and
// Context.Restore().

// ShapeAntialias sets whether to draw antialias for Context.Stroke() and
// Context.Fill(). It's enabled by default.
func (ctx *Context) ShapeAntialias(enabled bool) {
	var cEnabled int
	if enabled {
		cEnabled = 1
	}
	C.nvgShapeAntiAlias(ctx.c(), C.int(cEnabled))
}

// StrokeColor sets the current stroke style to a solid color.
func (ctx *Context) StrokeColor(color Color) {
	C.nvgStrokeColor(ctx.c(), color.c())
}

// StrokePaint sets the current stroke style to a paint, which can be one of the
// gradients or a pattern.
func (ctx *Context) StrokePaint(paint Paint) {
	C.nvgStrokePaint(ctx.c(), paint.c())
}

// FillColor sets the current fill style to a solid color.
func (ctx *Context) FillColor(color Color) {
	C.nvgFillColor(ctx.c(), color.c())
}

// FillPaint sets the current fill style to a pint, which can be one of the
// gradients or a pattern.
func (ctx *Context) FillPaint(paint Paint) {
	C.nvgFillPaint(ctx.c(), paint.c())
}

// MiterLimit sets the miter limit of the stroke style.
//
// Miter limit controls when a sharp corner is beveled.
func (ctx *Context) MiterLimit(limit float32) {
	C.nvgMiterLimit(ctx.c(), C.float(limit))
}

// StrokeWidth sets the stroke width of the stroke style.
func (ctx *Context) StrokeWidth(size float32) {
	C.nvgStrokeWidth(ctx.c(), C.float(size))
}

// LineCap sets how the end of the line (cap) is drawn. cap can be one of Butt
// (default), RoundCap and Square.
func (ctx *Context) LineCap(cap LineCap) {
	C.nvgLineCap(ctx.c(), C.int(cap))
}

// LineJoin sets how sharp path corners are drawn. join can be one of Miter
// (default), RoundJoin and Bevel.
func (ctx *Context) LineJoin(join LineJoin) {
	C.nvgLineJoin(ctx.c(), C.int(join))
}

// GlobalAlpha sets the transparency applied to all rendered shapes.
//
// Already transparent paths will get proportionally more transparent as well.
func (ctx *Context) GlobalAlpha(alpha float32) {
	C.nvgGlobalAlpha(ctx.c(), C.float(alpha))
}

// Transforms.
//
// The paths, gradients, patterns and scissor regions are transformed by an
// transformation matrix at the time when they are passed to the API. The
// current transformation matrix is an affine matrix:
//
//     [sx kx tx]
//     [ky sy ty]
//     [ 0  0  1]
//
// Where: sx,sy define scaling, kx,ky skewing, and tx,ty translation. The last
// row is assumed to be 0,0,1 and is not stored.
//
// Apart from Context.ResetTransform(), each transformation function first
// creates specific transformation matrix and pre-multiplies the current
// transformation by it.
//
// Current coordinate system (transformation) can be saved and restored using
// Context.Save() and Context.Restore().

// ResetTransform resets the current transform to an indentity matrix.
func (ctx *Context) ResetTransform() {
	C.nvgResetTransform(ctx.c())
}

// Transform premultiplies the current coordinate system by a matrix.
//
// The parameters are interpreted as a matrix as follows:
//
//     [a c e]
//     [b d f]
//     [0 0 1]
func (ctx *Context) Transform(a, b, c, d, e, f float32) {
	C.nvgTransform(ctx.c(), C.float(a), C.float(b), C.float(c), C.float(d), C.float(e), C.float(f))
}

// Translate translates the current coordinate system.
func (ctx *Context) Translate(x, y float32) {
	C.nvgTranslate(ctx.c(), C.float(x), C.float(y))
}

// Rotate rotates the current coordinate system. angle is specified in radians.
func (ctx *Context) Rotate(angle float32) {
	C.nvgRotate(ctx.c(), C.float(angle))
}

// SkewX skews the current coordinate system along the X axis. angle is
// specified in radians.
func (ctx *Context) SkewX(angle float32) {
	C.nvgSkewX(ctx.c(), C.float(angle))
}

// SkewY skews the current coordinate system along the Y axis. angle is
// specified in radians.
func (ctx *Context) SkewY(angle float32) {
	C.nvgSkewY(ctx.c(), C.float(angle))
}

// Scale scales the current coordinate system.
func (ctx *Context) Scale(x, y float32) {
	C.nvgScale(ctx.c(), C.float(x), C.float(y))
}

// CurrentTransform returns the top part (a-f) of the current transformation
// matrix.
//
//     [a c e]
//     [b d f]
//     [0 0 1]
func (ctx *Context) CurrentTransform() [6]float32 {
	cXform := make([]C.float, 6)
	C.nvgCurrentTransform(ctx.c(), &cXform[0])
	var xform [6]float32
	for i, num := range cXform {
		xform[i] = float32(num)
	}
	return xform
}

// The following functions can be used to make calculations on 2x3
// transformation matrices. A 2x3 matrix is represented as [6]float32.

// TransformIdentity sets the transform to an identity matrix.
func TransformIdentity(dst *[6]float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformIdentity(&cDst[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformTranslate sets the transform to a translation matrix.
func TransformTranslate(dst *[6]float32, tx, ty float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformTranslate(&cDst[0], C.float(tx), C.float(ty))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformScale sets the transform to a scale matrix.
func TransformScale(dst *[6]float32, sx, sy float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformScale(&cDst[0], C.float(sx), C.float(sy))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformRotate sets the transform to a rotate matrix. angle is specified in
// radians.
func TransformRotate(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformRotate(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformSkewX sets the transform to a skew-x matrix. angle is specified in
// radians.
func TransformSkewX(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformSkewX(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformSkewY sets the transform to a skew-y matrix. angle is specified in
// radians.
func TransformSkewY(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}

	C.nvgTransformSkewY(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformMultiply sets the transform to the result of multiplication of the
// two transforms, of A = A*B.
func TransformMultiply(dst *[6]float32, src [6]float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}
	cSrc := make([]C.float, 6)
	for i, num := range src {
		cSrc[i] = C.float(num)
	}

	C.nvgTransformMultiply(&cDst[0], &cSrc[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformPremultiply sets the transform to the result of multiplication of
// the two transforms, of A = B*A.
func TransformPremultiply(dst *[6]float32, src [6]float32) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}
	cSrc := make([]C.float, 6)
	for i, num := range src {
		cSrc[i] = C.float(num)
	}

	C.nvgTransformPremultiply(&cDst[0], &cSrc[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformInverse sets dst to the inverse of src. Returns true if the inverse
// could be calculated, else false.
func TransformInverse(dst *[6]float32, src [6]float32) (succeeded bool) {
	cDst := make([]C.float, 6)
	for i, num := range *dst {
		cDst[i] = C.float(num)
	}
	cSrc := make([]C.float, 6)
	for i, num := range src {
		cSrc[i] = C.float(num)
	}

	succeeded = int(C.nvgTransformInverse(&cDst[0], &cSrc[0])) == 1
	if succeeded {
		for i, num := range cDst {
			(*dst)[i] = float32(num)
		}
	}
	return
}

// TransformPoint transforms a point (srcX,srcY) by xform.
func TransformPoint(xform [6]float32, srcX, srcY float32) (dstX, dstY float32) {
	cXform := make([]C.float, 6)
	for i, num := range xform {
		cXform[i] = C.float(num)
	}
	var cDstX, cDstY C.float

	C.nvgTransformPoint(&cDstX, &cDstY, &cXform[0], C.float(srcX), C.float(srcY))

	dstX, dstY = float32(cDstX), float32(cDstY)
	return
}

// DegToRad converts degrees to radians.
func DegToRad(deg float32) float32 {
	return float32(C.nvgDegToRad(C.float(deg)))
}

// RadToDeg converts radians to degrees.
func RadToDeg(rad float32) float32 {
	return float32(C.nvgRadToDeg(C.float(rad)))
}

// Images
//
// NanoVGo allows you to load JPG, PNG, PSD, TGA, PIC and GIF files to be used
// for rendering. In addition you can upload your own image. The image loading
// is provided by stb_image.
// The parameter imageFlags is combination of flags defined in ImageFlag.

// Image is a handle to an loaded image.
type Image struct {
	cImage C.int
	ctx    *Context
}

func (image *Image) c() C.int {
	return image.cImage
}

// CreateImage creates an image by loading it from the disk from filename.
// Returns a handle to the image.
func (ctx *Context) CreateImage(filename string, imageFlags ImageFlag) *Image {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return &Image{
		cImage: C.nvgCreateImage(ctx.c(), cFilename, C.int(imageFlags)),
		ctx:    ctx,
	}
}

// CreateImageMem creates an image by loading it from data, a chunk of memory.
// Returns a handle to the image.
func (ctx *Context) CreateImageMem(imageFlags ImageFlag, data []uint8) *Image {
	var dataLen = len(data)
	var cData = make([]C.uchar, dataLen)
	for i, d := range data {
		cData[i] = C.uchar(d)
	}
	return &Image{
		cImage: C.nvgCreateImageMem(ctx.c(), C.int(imageFlags), &cData[0], C.int(dataLen)),
		ctx:    ctx,
	}
}

// CreateImageRGBA creates an image from data. Returns a handle to the image.
func (ctx *Context) CreateImageRGBA(width, height int, imageFlags ImageFlag, data []uint8) *Image {
	var cData = make([]C.uchar, len(data))
	for i, d := range data {
		cData[i] = C.uchar(d)
	}
	return &Image{
		cImage: C.nvgCreateImageRGBA(ctx.c(), C.int(width), C.int(height), C.int(imageFlags), &cData[0]),
		ctx:    ctx,
	}
}

// UpdateImage updates image data.
func (image *Image) UpdateImage(data []uint8) {
	var cData = make([]C.uchar, len(data))
	for i, d := range data {
		cData[i] = C.uchar(d)
	}
	C.nvgUpdateImage(image.ctx.c(), image.c(), &cData[0])
}

// Size returns the dimensions of image.
func (image *Image) Size() (width, height int) {
	var cWidth, cHeight C.int
	C.nvgImageSize(image.ctx.c(), image.c(), &cWidth, &cHeight)
	width, height = int(cWidth), int(cHeight)
	return
}

// Delete deletes image.
func (image *Image) Delete() {
	C.nvgDeleteImage(image.ctx.c(), image.c())
}

// Paints.
//
// NanoVGo supports four types of paints: linear gradient, box gradient, radial
// gradient and image pattern. These can be used as paints for strokes and
// fills.

// LinearGradient creates and returns a linear gradient. Parameters
// (startX,startY)-(endX,endY) specify the start and end coordinates of the
// linear gradient, startColor specifies the start color and endColor the end
// color.
//
// The gradient is transformed by the current transform when it is passed to
// Context.FillPaint() or Context.StrokePaint().
func (ctx *Context) LinearGradient(startX, startY, endX, endY float32, startColor, endColor Color) Paint {
	return Paint(C.nvgLinearGradient(ctx.c(), C.float(startX), C.float(startY), C.float(endX), C.float(endY), startColor.c(), endColor.c()))
}

// BoxGradient creates and returns a box gradient. A box gradient is a feathered
// rounded rectangle, it is useful for rendering drop shadows or highlights for
// boxes. Parameters (x,y) define the top-left corner of the rectangle,
// (width,height) define the size of the rectangle, radius defines the corner
// radius, and feather defines how blurry the border of the rectangle is.
// innerColor specifies the inner color and outerColor the outer color of the
// gradient.
//
// The gradient is transformed by the current transform when it is passed to
// Context.FillPaint() or Context.StrokePaint().
func (ctx *Context) BoxGradient(x, y, width, height, radius, feather float32, innerColor, outerColor Color) Paint {
	return Paint(C.nvgBoxGradient(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height), C.float(radius), C.float(feather), innerColor.c(), outerColor.c()))
}

// RadialGradient creates and returns a radian gradient. Parameters
// (centerX,centerY) specify the center, innerRadius and outerRadius specify the
// inner and outer radius of the gradient, startColor specifies the start color
// and endColor the end color.
//
// The gradient is transformed by the current transform when it is passed to
// Context.FillPaint() or Context.StrokePaint().
func (ctx *Context) RadialGradient(centerX, centerY, innerRadius, outerRadius float32, startColor, endColor Color) Paint {
	return Paint(C.nvgRadialGradient(ctx.c(), C.float(centerX), C.float(centerY), C.float(innerRadius), C.float(outerRadius), startColor.c(), endColor.c()))
}

// ImagePattern creates and returns an image pattern. Parameters (x,y) specify
// the left-top location of the image pattern, (imageWidth,imageHeight) the size
// of one image, angle rotation around the top-left corner, image is handle to
// the image to render.
//
// The gradient is transformed by the current transform when it is passed to
// Context.FillPaint() or Context.StrokePaint().
func (ctx *Context) ImagePattern(x, y, imageWidth, imageHeight, angle float32, image *Image, alpha float32) Paint {
	return Paint(C.nvgImagePattern(ctx.c(), C.float(x), C.float(y), C.float(imageWidth), C.float(imageHeight), C.float(angle), image.cImage, C.float(alpha)))
}

// Scissoring.
//
// Scissoring allows you to clip the rendering into a rectangle. This is useful
// for various user interface cases like rendering a text edit or a timeline.

// Scissor sets the current scissor rectangle. The scissor rectangle is
// transformed by the current transform.
func (ctx *Context) Scissor(x, y, width, height float32) {
	C.nvgScissor(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height))
}

// IntersectScissor intersects the current scissor rectangle with the specified
// rectangle. The scissor rectangle is transformed by the current transform.
//
// Note: in case the rotation of previous scissor rect differs from the current
// one, the intersection will be done between the specified rectangle and the
// previous scissor rectangle transformed in the current transform space. The
// resulting shape is always rectangle.
func (ctx *Context) IntersectScissor(x, y, width, height float32) {
	C.nvgIntersectScissor(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height))
}

// ResetScissor resets and disables scissoring.
func (ctx *Context) ResetScissor() {
	C.nvgResetScissor(ctx.c())
}

// Paths.
//
// Drawing a new shape starts with Context.BeginPath(), it clears all the
// currently defined paths. Then you define one or more paths and sub-paths
// which describe the shape. There are functions to draw common shapes like
// rectangles and circles, and lower level step-by-step functions, which allows
// to define a path curve by curve.
//
// NanoVGo uses oven-odd fill rule to draw the shapes. Solid shapes should have
// counter clockwise winding and holes should have counter clockwise order. To
// specify winding of a path you can call Context.PathWinding(). This is useful
// especially for the common shapes, which are drawn CCW.
//
// Finally you can fill the path using current fill style by calling
// Context.Fill(), and stroke it with current stroke style by calling
// Context.Stroke().
//
// The curve segments and sub-paths are transformed by the current transform.

// BeginPath clears the current path and sub-paths.
func (ctx *Context) BeginPath() {
	C.nvgBeginPath(ctx.c())
}

// MoveTo starts a new sub-path with point (x,y) as the first point.
func (ctx *Context) MoveTo(x, y float32) {
	C.nvgMoveTo(ctx.c(), C.float(x), C.float(y))
}

// LineTo adds a line segment from the last point in the path to point (x,y).
func (ctx *Context) LineTo(x, y float32) {
	C.nvgLineTo(ctx.c(), C.float(x), C.float(y))
}

// BezierTo adds a cubic bezier segment from the last point in the path via two
// control points ((c1X,c1Y) and (c2X,c2Y)) to point (x,y).
func (ctx *Context) BezierTo(c1X, c1Y, c2X, c2Y, x, y float32) {
	C.nvgBezierTo(ctx.c(), C.float(c1X), C.float(c1Y), C.float(c2X), C.float(c2Y), C.float(x), C.float(y))
}

// QuadTo adds a quadratic bezier segment from the last point in the path via a
// control point (cX,cY) to point (x,y).
func (ctx *Context) QuadTo(cX, cY, x, y float32) {
	C.nvgQuadTo(ctx.c(), C.float(cX), C.float(cY), C.float(x), C.float(y))
}

// ArcTo adds an arc segment at the corner defined by the last path point, and
// two points (x1,y1) and (x2,y2).
func (ctx *Context) ArcTo(x1, y1, x2, y2, radius float32) {
	C.nvgArcTo(ctx.c(), C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(radius))
}

// ClosePath closes current sub-path with a line segment.
func (ctx *Context) ClosePath() {
	C.nvgClosePath(ctx.c())
}

// PathWinding sets the current sub-path winding, see Winding.
func (ctx *Context) PathWinding(direction Winding) {
	C.nvgPathWinding(ctx.c(), C.int(direction))
}

// Arc creates a new circle arc shaped sub-path. The arc center is at (x,y), the
// arc radius is radius, and the arc is drawn from angle angle0 to angle1, and
// swept in direction direction (CCW or CW).
//
// Angles are specified in radians.
func (ctx *Context) Arc(x, y, radius, angle0, angle1 float32, direction Winding) {
	C.nvgArc(ctx.c(), C.float(x), C.float(y), C.float(radius), C.float(angle0), C.float(angle1), C.int(direction))
}

// Rect creates a new rectangle shaped sub-path.
func (ctx *Context) Rect(x, y, width, height float32) {
	C.nvgRect(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height))
}

// RoundedRect creates a new rounded rectangle shaped sub-path.
func (ctx *Context) RoundedRect(x, y, width, height, radius float32) {
	C.nvgRoundedRect(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height), C.float(radius))
}

// RoundedRectVarying creates a new rounded rectangle shaped sub-path with
// varying radii for each corner.
func (ctx *Context) RoundedRectVarying(x, y, width, height, radiusTopLeft, radiusTopRight, radiusBottomRight, radiusBottomLeft float32) {
	C.nvgRoundedRectVarying(ctx.c(), C.float(x), C.float(y), C.float(width), C.float(height), C.float(radiusTopLeft), C.float(radiusTopRight), C.float(radiusBottomRight), C.float(radiusBottomLeft))
}

// Ellipse creates a new ellipse shape sub-path. The center is at (x,y).
func (ctx *Context) Ellipse(x, y, radiusX, radiusY float32) {
	C.nvgEllipse(ctx.c(), C.float(x), C.float(y), C.float(radiusX), C.float(radiusY))
}

// Circle creates a new circle shaped sub-path. The center is at (x,y).
func (ctx *Context) Circle(x, y, radius float32) {
	C.nvgCircle(ctx.c(), C.float(x), C.float(y), C.float(radius))
}

// Fill fills the current path with the current fill style.
func (ctx *Context) Fill() {
	C.nvgFill(ctx.c())
}

// Stroke strokes the current path with the current stroke style.
func (ctx *Context) Stroke() {
	C.nvgStroke(ctx.c())
}

// Text.
//
// NanoVGo allows you to load .ttf files and use the font to render text.
//
// The appearance of the text can be defined by setting the current text style
// and by specifying the fill color. Common text and font settings such as font
// size, letter spacing and text align are supported. Font blur allows you to
// create simple text effects such as drop shadows.
//
// At render time the font face can be set based on the font handles or name.
//
// Font measure functions return values in local space, the calculations are
// carried in the same resolution as the final rendering. This is done because
// the text glyph positions are snapped to the nearest pixel sharp rendering.
//
// The local space means that values are not rotated or scale as per the current
// transformation. For example if you set font size to 12, which would mean that
// line height is 16, then regardless of the current scaling and rotatino, the
// returned line height is always 16. Some measures may vary because of the
// scaling since aforementioned pixel snapping.
//
// While this may sound a little odd, the setup allows you to always render the
// same way regardless of scaling. I.e. following works regardless of scaling:
//
//     text := "Text me up."
//     ctx.TextBounds(x, y, text, nil, bounds)
//     ctx.BeginPath()
//     ctx.RoundedRect(bounds[0], bounds[1], bounds[2]-bounds[0], bounds[3]-bounds[1])
//     ctx.Fill()
//
// Note: currently only solid color fill is supported for text.

// Font is a handle to a created font.
type Font struct {
	cFont C.int
}

func (font *Font) c() C.int {
	return font.cFont
}

// CreateFont creates a font by loading it from the disk from filename. Returns
// a handle to the font.
func (ctx *Context) CreateFont(name, filename string) *Font {
	var cName = C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cFilename = C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	return &Font{
		cFont: C.nvgCreateFont(ctx.c(), cName, cFilename),
	}
}

// CreateFontMem creates a font by loading it from data, a memory chunk. Returns
// a handle to the font.
func (ctx *Context) CreateFontMem(name string, data []uint8, freeData int) *Font {
	var cName = C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var dataLen = len(data)
	var cData = make([]C.uchar, dataLen)
	for i, d := range data {
		cData[i] = C.uchar(d)
	}

	return &Font{
		cFont: C.nvgCreateFontMem(ctx.c(), cName, &cData[0], C.int(dataLen), C.int(freeData)),
	}
}

// FindFont finds a loaded font with name, and returns a handle to it, or nil if
// the font is not found.
func (ctx *Context) FindFont(name string) *Font {
	var cName = C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cFont = C.nvgFindFont(ctx.c(), cName)
	if int(cFont) != -1 {
		return &Font{cFont: cFont}
	}
	return nil
}

// AddFallbackFontID adds a fallback font by its handle.
func (ctx *Context) AddFallbackFontID(baseFont, fallbackFont *Font) {
	C.nvgAddFallbackFontId(ctx.c(), baseFont.c(), fallbackFont.c())
}

// AddFallbackFont adds a fallback font by its name.
func (ctx *Context) AddFallbackFont(baseFont, fallbackFont string) {
	var cBaseFont = C.CString(baseFont)
	defer C.free(unsafe.Pointer(cBaseFont))
	var cFallbackFont = C.CString(fallbackFont)
	defer C.free(unsafe.Pointer(cFallbackFont))

	C.nvgAddFallbackFont(ctx.c(), cBaseFont, cFallbackFont)
}

// FontSize sets the font size of the current text style.
func (ctx *Context) FontSize(size float32) {
	C.nvgFontSize(ctx.c(), C.float(size))
}

// FontBlur sets the blur of the current text style.
func (ctx *Context) FontBlur(blur float32) {
	C.nvgFontBlur(ctx.c(), C.float(blur))
}

// TextLetterSpacing sets the letter spacing of the current text style.
func (ctx *Context) TextLetterSpacing(spacing float32) {
	C.nvgTextLetterSpacing(ctx.c(), C.float(spacing))
}

// TextLineHeight sets the proportional line height of the current text style.
// The line height is specified as multiple of font size.
func (ctx *Context) TextLineHeight(lineHeight float32) {
	C.nvgTextLineHeight(ctx.c(), C.float(lineHeight))
}

// TextAlign sets the text align of the current text style, see Align for
// options.
func (ctx *Context) TextAlign(align Align) {
	C.nvgTextAlign(ctx.c(), C.int(align))
}

// FontFaceID sets the font face of the current text style with a font handle.
func (ctx *Context) FontFaceID(font *Font) {
	C.nvgFontFaceId(ctx.c(), font.c())
}

// FontFace sets the font face of the current text style by the font name.
func (ctx *Context) FontFace(font string) {
	var cFont = C.CString(font)
	defer C.free(unsafe.Pointer(cFont))

	C.nvgFontFace(ctx.c(), cFont)
}

// Text draws text at location (x,y).
func (ctx *Context) Text(x, y float32, text string) {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.nvgText(ctx.c(), C.float(x), C.float(y), cText, (*C.char)(C.NULL))
}

// TextBox draws multi-line text at location (x,y) wrapped at the width
// breakRowWidth.
//
// White space is stripped at the beginning of the rows, the text is split at
// word boundaries or when new-line characters are encountered.
//
// Words longer than the max width are split at the nearest character (i.e. no
// hyphenation).
func (ctx *Context) TextBox(x, y, breakRowWidth float32, text string) {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.nvgTextBox(ctx.c(), C.float(x), C.float(y), C.float(breakRowWidth), cText, (*C.char)(C.NULL))
}

// TextBounds measures the specified text. Returns the horizontal advance of the
// measured text (i.e. where the next character should be drawn), and the
// bounding box of the text. The bounds values are [xmin, ymin, xmax, ymax].
//
// Measured values are returned in local coordinate space.
func (ctx *Context) TextBounds(x, y float32, text string) (float32, [4]float32) {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cBounds = make([]C.float, 4)
	var cAdvance = C.nvgTextBounds(ctx.c(), C.float(x), C.float(y), cText, (*C.char)(C.NULL), &cBounds[0])
	var bounds [4]float32
	for i := 0; i < 4; i++ {
		bounds[i] = float32(cBounds[i])
	}
	return float32(cAdvance), bounds
}

// TextBoxBounds measures the specified multi-line text. Returns the bounding
// box of the text. The bounds values are [xmin, ymin, xmax, ymax].
//
// Measured values are returned in local coordinate space.
func (ctx *Context) TextBoxBounds(x, y, breakRowWidth float32, text string) [4]float32 {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cBounds = make([]C.float, 4)
	C.nvgTextBoxBounds(ctx.c(), C.float(x), C.float(y), C.float(breakRowWidth), cText, (*C.char)(C.NULL), &cBounds[0])
	var bounds [4]float32
	for i := 0; i < 4; i++ {
		bounds[i] = float32(cBounds[i])
	}
	return bounds
}

// TextGlyphPositions calculates the glyph x position of text.
//
// Measured values are returned in local coordinate space.
func (ctx *Context) TextGlyphPositions(x, y float32, text string, maxPositions int) []GlyphPosition {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cPositions = make([]C.NVGglyphPosition, maxPositions)
	var count = int(C.nvgTextGlyphPositions(ctx.c(), C.float(x), C.float(y), cText, (*C.char)(C.NULL), &cPositions[0], C.int(maxPositions)))
	var positions = make([]GlyphPosition, 0, count)
	for i := 0; i < count; i++ {
		positions = append(positions, GlyphPosition(cPositions[i]))
	}
	return positions
}

// TextMetrics returns the vertical metrics based on the current text style.
//
// Measured values are returned in local coordinate space.
func (ctx *Context) TextMetrics() (ascender, descender, lineh float32) {
	var cAscender, cDescender, cLineh C.float
	C.nvgTextMetrics(ctx.c(), &cAscender, &cDescender, &cLineh)
	ascender, descender, lineh = float32(cAscender), float32(cDescender), float32(cLineh)
	return
}

// TextBreakLines breaks the text into lines.
//
// White space is stripped at the beginning of the rows, the text is split at
// word boundaries or when new-line characters are encountered.
//
// Words longer than the max width are split at the nearest character (i.e. no
// hyphenation).
func (ctx *Context) TextBreakLines(text string, breakRowWidth float32, maxRows int) []TextRow {
	var cText = C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cRows = make([]C.NVGtextRow, maxRows)
	var count = int(C.nvgTextBreakLines(ctx.c(), cText, (*C.char)(C.NULL), C.float(breakRowWidth), &cRows[0], C.int(maxRows)))
	var rows = make([]TextRow, 0, count)
	for i := 0; i < count; i++ {
		rows = append(rows, TextRow(cRows[i]))
	}
	return rows
}
