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
	"math"

	"github.com/beta/nanovgo"
)

func loadDemoData(vg *nanovgo.Context) (*demoData, error) {
	if vg == nil {
		return nil, fmt.Errorf("NanoVGo context is nil")
	}

	var data = new(demoData)

	for i := 0; i < 12; i++ {
		var file = fmt.Sprintf("../nanovg/example/images/image%d.jpg", i+1)
		data.images[i] = vg.CreateImage(file, 0)
		if data.images[i] == nil {
			return nil, fmt.Errorf("failed to load %s", file)
		}
	}

	data.fontIcons = vg.CreateFont("icons", "../nanovg/example/entypo.ttf")
	if data.fontIcons == nil {
		return nil, fmt.Errorf("failed to add font icons")
	}
	data.fontNormal = vg.CreateFont("sans", "../nanovg/example/Roboto-Regular.ttf")
	if data.fontNormal == nil {
		return nil, fmt.Errorf("failed to add font italic")
	}
	data.fontBold = vg.CreateFont("sans-bold", "../nanovg/example/Roboto-Bold.ttf")
	if data.fontBold == nil {
		return nil, fmt.Errorf("failed to add font bold")
	}
	data.fontEmoji = vg.CreateFont("emoji", "../nanovg/example/NotoEmoji-Regular.ttf")
	if data.fontEmoji == nil {
		return nil, fmt.Errorf("failed to add font emoji")
	}
	vg.AddFallbackFontID(data.fontNormal, data.fontEmoji)
	vg.AddFallbackFontID(data.fontBold, data.fontEmoji)

	return data, nil
}

type demoData struct {
	fontNormal, fontBold, fontIcons, fontEmoji *nanovgo.Font
	images                                     [12]*nanovgo.Image
}

func (data *demoData) free(vg *nanovgo.Context) {
	if vg == nil {
		return
	}

	for i := 0; i < 12; i++ {
		data.images[i].Delete()
	}
}

func (data *demoData) render(vg *nanovgo.Context, mx, my, width, height float32, t float64, blowup bool) {
	drawEyes(vg, width-250, 50, 150, 100, mx, my, t)
	drawParagraph(vg, width-450, 50, 150, 100, mx, my)
	drawGraph(vg, 0, height/2, width, height/2, t)
	drawColorwheel(vg, width-300, height-300, 250.0, 250.0, t)

	// Line joints.
	drawLines(vg, 120, height-50, 600, 50, t)

	// Line caps.
	drawWidths(vg, 10, 50, 30)

	// Line caps.
	drawCaps(vg, 10, 300, 30)

	drawScissor(vg, 50, height-80, t)

	vg.Save()
	if blowup {
		vg.Rotate(float32(math.Sin(t*0.3)) * 5.0 / 180.0 * float32(math.Pi))
		vg.Scale(2.0, 2.0)
	}

	// Widgets
	drawWindow(vg, "Widgets `n Stuff", 50, 50, 300, 400)
	var x float32 = 60
	var y float32 = 95
	drawSearchBox(vg, "Search", x, y, 280, 25)
	y += 40
	drawDropDown(vg, "Effects", x, y, 280, 28)
	var popy = y + 14
	y += 45

	// Form
	drawLabel(vg, "Login", x, y, 280, 20)
	y += 25
	drawEditBox(vg, "Email", x, y, 280, 28)
	y += 35
	drawEditBox(vg, "Password", x, y, 280, 28)
	y += 38
	drawCheckBox(vg, "Remember me", x, y, 140, 28)
	drawButton(vg, IconLogin, "Sign in", x+138, y, 140, 28, vg.RGBA(0, 96, 128, 255))
	y += 45

	// Slider
	drawLabel(vg, "Diameter", x, y, 280, 20)
	y += 25
	drawEditBoxNum(vg, "123.00", "px", x+180, y, 100, 28)
	drawSlider(vg, 0.4, x, y, 170, 28)
	y += 55

	drawButton(vg, IconTrash, "Delete", x, y, 160, 28, vg.RGBA(128, 16, 8, 255))
	drawButton(vg, 0, "Cancel", x+170, y, 110, 28, vg.RGBA(0, 0, 0, 0))

	// Thumbnails box
	drawThumbnails(vg, 365, popy-30, 160, 300, data.images, t)

	vg.Restore()
}

func saveScreenShot(width, height int, premult bool, name string) {
	// TODO: implement saveScreenShot.
}

// Icon constants.
const (
	IconSearch       = 0x1F50D
	IconCircledCross = 0x2716
	IconChevronRight = 0xE75E
	IconCheck        = 0x2713
	IconLogin        = 0xE740
	IconTrash        = 0xE729
)

func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func clampf(a, min, max float32) float32 {
	if a < min {
		return min
	} else if a > max {
		return max
	} else {
		return a
	}
}

func isBlack(color nanovgo.Color) bool {
	return color.R() == 0 && color.G() == 0 && color.B() == 0 && color.A() == 0
}

func drawWindow(vg *nanovgo.Context, title string, x, y, width, height float32) {
	var cornerRadius float32 = 3.0

	vg.Save()

	// Window.
	vg.BeginPath()
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.FillColor(vg.RGBA(28, 30, 34, 192))
	vg.Fill()

	// Drop shadow.
	var shadowPaint = vg.BoxGradient(x, y+2, width, height, cornerRadius*2, 10, vg.RGBA(0, 0, 0, 128), vg.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(x-10, y-10, width+20, height+30)
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(shadowPaint)
	vg.Fill()

	// Header.
	var headerPaint = vg.LinearGradient(x, y, x, y+15, vg.RGBA(255, 255, 255, 0), vg.RGBA(0, 0, 0, 16))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+1, width-2, 30, cornerRadius-1)
	vg.FillPaint(headerPaint)
	vg.Fill()
	vg.BeginPath()
	vg.MoveTo(x+0.5, y+0.5+30)
	vg.LineTo(x+0.5+width-1, y+0.5+30)
	vg.StrokeColor(vg.RGBA(0, 0, 0, 32))
	vg.Stroke()

	vg.FontSize(18.0)
	vg.FontFace("sans-bold")
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)

	vg.FontBlur(2)
	vg.FillColor(vg.RGBA(0, 0, 0, 128))
	vg.Text(x+width/2, y+16+1, title)

	vg.FontBlur(0)
	vg.FillColor(vg.RGBA(220, 220, 220, 160))
	vg.Text(x+width/2, y+16, title)

	vg.Restore()
}

func drawSearchBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	// TODO: implement drawSearchBox.
}

func drawDropDown(vg *nanovgo.Context, text string, x, y, width, height float32) {
	// TODO: implement drawDropDown.
}

func drawLabel(vg *nanovgo.Context, text string, x, y, width, height float32) {
	// TODO: implement drawLabel.
}

func drawEditBoxBase(vg *nanovgo.Context, x, y, width, height float32) {
	// TODO: implement drawEditBoxBase.
}

func drawEditBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	// TODO: implement drawEditBox.
}

func drawEditBoxNum(vg *nanovgo.Context, text string, units string, x, y, width, height float32) {
	// TODO: implement drawEditBoxNum.
}

func drawCheckBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	// TODO: implement drawCheckBox.
}

func drawButton(vg *nanovgo.Context, preIcon int, text string, x, y, width, height float32, color nanovgo.Color) {
	// TODO: implement drawButton.
}

func drawSlider(vg *nanovgo.Context, pos, x, y, width, height float32) {
	// TODO: implement drawSlider.
}

func drawEyes(vg *nanovgo.Context, x, y, width, height, mx, my float32, t float64) {
	// TODO: implement drawEyes.
}

func drawGraph(vg *nanovgo.Context, x, y, width, height float32, t float64) {
	// TODO: implement drawGraph.
}

func drawSpinner(vg *nanovgo.Context, cx, cy, r float32, t float64) {
	// TODO: implement drawSpinner.
}

func drawThumbnails(vg *nanovgo.Context, x, y, width, height float32, images [12]*nanovgo.Image, t float64) {
	// TODO: implement drawThumbnails.
}

func drawColorwheel(vg *nanovgo.Context, x, y, w, h float32, t float64) {
	// TODO: implement drawColorwheel.
}

func drawLines(vg *nanovgo.Context, x, y, width, height float32, t float64) {
	// TODO: implement drawLines.
}

func drawParagraph(vg *nanovgo.Context, x, y, width, height, mx, my float32) {
	// TODO: implement drawParagraph.
}

func drawWidths(vg *nanovgo.Context, x, y, width float32) {
	// TODO: implement drawWidths.
}

func drawCaps(vg *nanovgo.Context, x, y, width float32) {
	// TODO: implement drawCaps.
}

func drawScissor(vg *nanovgo.Context, x, y float32, t float64) {
	// TODO: implement drawScissor.
}
