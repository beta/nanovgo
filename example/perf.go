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

package main

import (
	"fmt"

	"github.com/beta/nanovgo"
)

// GraphHistoryCount is the length of historical values to be stored in a
// performance graph.
const GraphHistoryCount = 100

// NewPerfGraph creates and returns a new performance graph.
func NewPerfGraph(style PerfGraphStyle, name string) *PerfGraph {
	return &PerfGraph{
		style: style,
		name:  name,
	}
}

// PerfGraph stores and analyzes the performance values.
type PerfGraph struct {
	style  PerfGraphStyle
	name   string
	values [GraphHistoryCount]float64
	head   int
}

// PerfGraphStyle specifies the rendering style of a performance graph.
type PerfGraphStyle int

// PerfGraph styles.
const (
	GraphRenderFPS PerfGraphStyle = iota
	GraphRenderMS
	GraphRenderPercent
)

// Update records the frameTime into fps.
func (fps *PerfGraph) Update(frameTime float64) {
	fps.head = (fps.head + 1) % GraphHistoryCount
	fps.values[fps.head] = frameTime
}

// Average returns the average FPS.
func (fps *PerfGraph) Average() float64 {
	var avg float64
	for i := 0; i < GraphHistoryCount; i++ {
		avg += fps.values[i]
	}
	return avg / float64(GraphHistoryCount)
}

// Render renders the performance graph onto location (x,y) of ctx.
func (fps *PerfGraph) Render(ctx *nanovgo.Context, x, y float32) {
	var avg = fps.Average()
	var width, height float32 = 200, 35

	ctx.BeginPath()
	ctx.Rect(x, y, width, height)
	ctx.FillColor(ctx.RGBA(0, 0, 0, 128))
	ctx.Fill()

	ctx.BeginPath()
	ctx.MoveTo(x, y+height)
	if fps.style == GraphRenderFPS {
		for i := 0; i < GraphHistoryCount; i++ {
			var v = 1.0 / (0.00001 + fps.values[(fps.head+i)/GraphHistoryCount])
			if v > 80.0 {
				v = 80.0
			}
			var vx = x + (float32(i)/(GraphHistoryCount-1))*width
			var vy = y + height - (float32(v/80.0) * height)
			ctx.LineTo(vx, vy)
		}
	} else if fps.style == GraphRenderPercent {
		for i := 0; i < GraphHistoryCount; i++ {
			var v = fps.values[(fps.head+i)%GraphHistoryCount] * 1.0
			if v > 100.0 {
				v = 100.0
			}
			var vx = x + (float32(i)/(GraphHistoryCount-1))*width
			var vy = y + height - (float32(v/100.0) * height)
			ctx.LineTo(vx, vy)
		}
	} else {
		for i := 0; i < GraphHistoryCount; i++ {
			var v = fps.values[(fps.head+i)%GraphHistoryCount] * 1000.0
			if v > 20.0 {
				v = 20.0
			}
			var vx = x + (float32(i)/(GraphHistoryCount-1))*width
			var vy = y + height - (float32(v/20.0) * height)
			ctx.LineTo(vx, vy)
		}
	}
	ctx.LineTo(x+width, y+height)
	ctx.FillColor(ctx.RGBA(255, 192, 0, 128))
	ctx.Fill()

	ctx.FontFace("sans")

	if len(fps.name) > 0 {
		ctx.FontSize(14.0)
		ctx.TextAlign(nanovgo.AlignLeft | nanovgo.AlignTop)
		ctx.FillColor(ctx.RGBA(240, 240, 240, 192))
		ctx.Text(x+3, y+1, fps.name)
	}

	if fps.style == GraphRenderFPS {
		ctx.FontSize(18.0)
		ctx.TextAlign(nanovgo.AlignRight | nanovgo.AlignTop)
		ctx.FillColor(ctx.RGBA(240, 240, 240, 255))
		ctx.Text(x+width-3, y+1, fmt.Sprintf("%.2f FPS", 1.0/avg))

		ctx.FontSize(15.0)
		ctx.TextAlign(nanovgo.AlignRight | nanovgo.AlignBottom)
		ctx.FillColor(ctx.RGBA(240, 240, 240, 160))
		ctx.Text(x+width-3, y+height-1, fmt.Sprintf("%.2f ms", avg*1000.0))
	} else if fps.style == GraphRenderPercent {
		ctx.FontSize(18.0)
		ctx.TextAlign(nanovgo.AlignRight | nanovgo.AlignTop)
		ctx.FillColor(ctx.RGBA(240, 240, 240, 255))
		ctx.Text(x+width-3, y+1, fmt.Sprintf("%.1f %%", avg*1.0))
	} else {
		ctx.FontSize(18.0)
		ctx.TextAlign(nanovgo.AlignRight | nanovgo.AlignTop)
		ctx.FillColor(ctx.RGBA(240, 240, 240, 255))
		ctx.Text(x+width-3, y+1, fmt.Sprintf("%.2f ms", avg*1000.0))
	}
}
