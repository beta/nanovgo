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
#include <stdlib.h>
#include <OpenGL/gl3.h>
#include "nanovg/src/nanovg.h"
#include "nanovg/src/nanovg_gl.h"
#include "nanovg/src/nanovg_gl_utils.h"
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
	cXform := make([]C.float, 0, 6)
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
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformIdentity(&cDst[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformTranslate sets the transform to a translation matrix.
func TransformTranslate(dst *[6]float32, tx, ty float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformTranslate(&cDst[0], C.float(tx), C.float(ty))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformScale sets the transform to a scale matrix.
func TransformScale(dst *[6]float32, sx, sy float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformScale(&cDst[0], C.float(sx), C.float(sy))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformRotate sets the transform to a rotate matrix. angle is specified in
// radians.
func TransformRotate(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformRotate(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformSkewX sets the transform to a skew-x matrix. angle is specified in
// radians.
func TransformSkewX(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformSkewX(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformSkewY sets the transform to a skew-y matrix. angle is specified in
// radians.
func TransformSkewY(dst *[6]float32, angle float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}

	C.nvgTransformSkewY(&cDst[0], C.float(angle))

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformMultiply sets the transform to the result of multiplication of the
// two transforms, of A = A*B.
func TransformMultiply(dst *[6]float32, src [6]float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}
	cSrc := make([]C.float, 0, 6)
	for _, num := range src {
		cSrc = append(cSrc, C.float(num))
	}

	C.nvgTransformMultiply(&cDst[0], &cSrc[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformPremultiply sets the transform to the result of multiplication of
// the two transforms, of A = B*A.
func TransformPremultiply(dst *[6]float32, src [6]float32) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}
	cSrc := make([]C.float, 0, 6)
	for _, num := range src {
		cSrc = append(cSrc, C.float(num))
	}

	C.nvgTransformPremultiply(&cDst[0], &cSrc[0])

	for i, num := range cDst {
		(*dst)[i] = float32(num)
	}
}

// TransformInverse sets dst to the inverse of src. Returns true if the inverse
// could be calculated, else false.
func TransformInverse(dst *[6]float32, src [6]float32) (succeeded bool) {
	cDst := make([]C.float, 0, 6)
	for _, num := range *dst {
		cDst = append(cDst, C.float(num))
	}
	cSrc := make([]C.float, 0, 6)
	for _, num := range src {
		cSrc = append(cSrc, C.float(num))
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
	cXform := make([]C.float, 0, 6)
	for _, num := range xform {
		cXform = append(cXform, C.float(num))
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

// CreateImage creates an image by loading it from the disk from filename.
// Returns a handle to the image.
func (ctx *Context) CreateImage(filename string, imageFlags ImageFlag) Image {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return Image{
		cImage: C.nvgCreateImage(ctx.c(), cFilename, C.int(imageFlags)),
		ctx:    ctx,
	}
}

// CreateImageMem creates an image by loading it from data, a chunk of memory.
// Returns a handle to the image.
func (ctx *Context) CreateImageMem(imageFlags ImageFlag, data []uint8) Image {
	var dataLen = len(data)
	var cData = make([]C.uchar, 0, dataLen)
	for _, d := range data {
		cData = append(cData, C.uchar(d))
	}
	return Image{
		cImage: C.nvgCreateImageMem(ctx.c(), C.int(imageFlags), &cData[0], C.int(dataLen)),
		ctx:    ctx,
	}
}

// CreateImageRGBA creates an image from data. Returns a handle to the image.
func (ctx *Context) CreateImageRGBA(width, height int, imageFlags ImageFlag, data []uint8) Image {
	var cData = make([]C.uchar, 0, len(data))
	for _, d := range data {
		cData = append(cData, C.uchar(d))
	}
	return Image{
		cImage: C.nvgCreateImageRGBA(ctx.c(), C.int(width), C.int(height), C.int(imageFlags), &cData[0]),
		ctx:    ctx,
	}
}

// UpdateImage updates image data.
func (image Image) UpdateImage(data []uint8) {
	var cData = make([]C.uchar, 0, len(data))
	for _, d := range data {
		cData = append(cData, C.uchar(d))
	}
	C.nvgUpdateImage(image.ctx.c(), image.cImage, &cData[0])
}

// Size returns the dimensions of image.
func (image Image) Size() (width, height int) {
	var cWidth, cHeight C.int
	C.nvgImageSize(image.ctx.c(), image.cImage, &cWidth, &cHeight)
	width, height = int(cWidth), int(cHeight)
	return
}

// Delete deletes image.
func (image Image) Delete() {
	C.nvgDeleteImage(image.ctx.c(), image.cImage)
}
