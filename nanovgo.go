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

// Context is a NanoVGo context for vector graphics rendering.
type Context C.NVGcontext

func (ctx *Context) c() *C.NVGcontext {
	return (*C.NVGcontext)(ctx)
}

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
