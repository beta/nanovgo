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
	drawButton(vg, IconLogin, "Sign in", x+138, y, 140, 28, nanovgo.RGBA(0, 96, 128, 255))
	y += 45

	// Slider
	drawLabel(vg, "Diameter", x, y, 280, 20)
	y += 25
	drawEditBoxNum(vg, "123.00", "px", x+180, y, 100, 28)
	drawSlider(vg, 0.4, x, y, 170, 28)
	y += 55

	drawButton(vg, IconTrash, "Delete", x, y, 160, 28, nanovgo.RGBA(128, 16, 8, 255))
	drawButton(vg, 0, "Cancel", x+170, y, 110, 28, nanovgo.RGBA(0, 0, 0, 0))

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

func cpToUTF8(cp int) string {
	var n int
	switch {
	case cp < 0x80:
		n = 1
	case cp < 0x800:
		n = 2
	case cp < 0x10000:
		n = 3
	case cp < 0x200000:
		n = 4
	case cp < 0x4000000:
		n = 5
	case cp <= 0x7fffffff:
		n = 6
	}
	var str = make([]byte, 6)
	switch n {
	case 6:
		str[5] = byte(0x80 | (cp & 0x3f))
		cp = cp >> 6
		cp |= 0x4000000
		fallthrough
	case 5:
		str[4] = byte(0x80 | (cp & 0x3f))
		cp = cp >> 6
		cp |= 0x200000
		fallthrough
	case 4:
		str[3] = byte(0x80 | (cp & 0x3f))
		cp = cp >> 6
		cp |= 0x10000
		fallthrough
	case 3:
		str[2] = byte(0x80 | (cp & 0x3f))
		cp = cp >> 6
		cp |= 0x800
		fallthrough
	case 2:
		str[1] = byte(0x80 | (cp & 0x3f))
		cp = cp >> 6
		cp |= 0xc0
		fallthrough
	case 1:
		str[0] = byte(cp)
	}
	return string(str)
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func drawWindow(vg *nanovgo.Context, title string, x, y, width, height float32) {
	var cornerRadius float32 = 3.0

	vg.Save()

	// Window.
	vg.BeginPath()
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.FillColor(nanovgo.RGBA(28, 30, 34, 192))
	vg.Fill()

	// Drop shadow.
	var shadowPaint = vg.BoxGradient(x, y+2, width, height, cornerRadius*2, 10, nanovgo.RGBA(0, 0, 0, 128), nanovgo.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(x-10, y-10, width+20, height+30)
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(shadowPaint)
	vg.Fill()

	// Header.
	var headerPaint = vg.LinearGradient(x, y, x, y+15, nanovgo.RGBA(255, 255, 255, 0), nanovgo.RGBA(0, 0, 0, 16))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+1, width-2, 30, cornerRadius-1)
	vg.FillPaint(headerPaint)
	vg.Fill()
	vg.BeginPath()
	vg.MoveTo(x+0.5, y+0.5+30)
	vg.LineTo(x+0.5+width-1, y+0.5+30)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 32))
	vg.Stroke()

	vg.FontSize(18.0)
	vg.FontFace("sans-bold")
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)

	vg.FontBlur(2)
	vg.FillColor(nanovgo.RGBA(0, 0, 0, 128))
	vg.Text(x+width/2, y+16+1, title)

	vg.FontBlur(0)
	vg.FillColor(nanovgo.RGBA(220, 220, 220, 160))
	vg.Text(x+width/2, y+16, title)

	vg.Restore()
}

func drawSearchBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	var cornerRadius = height/2 - 1

	// Edit.
	var bg = vg.BoxGradient(x, y+1.5, width, height, height/2, 5, nanovgo.RGBA(0, 0, 0, 16), nanovgo.RGBA(0, 0, 0, 92))
	vg.BeginPath()
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.FillPaint(bg)
	vg.Fill()

	vg.FontSize(height * 1.3)
	vg.FontFace("icons")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 64))
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)
	vg.Text(x+height*0.55, y+height*0.55, cpToUTF8(IconSearch))

	vg.FontSize(20.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 32))

	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
	vg.Text(x+height*1.05, y+height*0.5, text)

	vg.FontSize(height * 1.3)
	vg.FontFace("icons")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 32))
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)
	vg.Text(x+width-height*0.55, y+height*0.55, cpToUTF8(IconCircledCross))
}

func drawDropDown(vg *nanovgo.Context, text string, x, y, width, height float32) {
	var cornerRadius float32 = 4.0

	var bg = vg.LinearGradient(x, y, x, y+height, nanovgo.RGBA(255, 255, 255, 16), nanovgo.RGBA(0, 0, 0, 16))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+1, width-2, height-2, cornerRadius-1)
	vg.FillPaint(bg)
	vg.Fill()

	vg.BeginPath()
	vg.RoundedRect(x+0.5, y+0.5, width-1, height-1, cornerRadius-0.5)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 48))
	vg.Stroke()

	vg.FontSize(20.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 160))
	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
	vg.Text(x+height*0.3, y+height*0.5, text)

	vg.FontSize(height * 1.3)
	vg.FontFace("icons")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 64))
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)
	vg.Text(x+width-height*0.5, y+height*0.5, cpToUTF8(IconChevronRight))
}

func drawLabel(vg *nanovgo.Context, text string, x, y, width, height float32) {
	vg.FontSize(18.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 128))

	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
	vg.Text(x, y+height*0.5, text)
}

func drawEditBoxBase(vg *nanovgo.Context, x, y, width, height float32) {
	var bg = vg.BoxGradient(x+1, y+1+1.5, width-2, height-2, 3, 4, nanovgo.RGBA(255, 255, 255, 32), nanovgo.RGBA(32, 32, 32, 32))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+1, width-2, height-2, 4-1)
	vg.FillPaint(bg)
	vg.Fill()

	vg.BeginPath()
	vg.RoundedRect(x+0.5, y+0.5, width-1, height-1, 4-0.5)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 48))
	vg.Stroke()
}

func drawEditBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	drawEditBoxBase(vg, x, y, width, height)

	vg.FontSize(20.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 64))
	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignLeft)
	vg.Text(x+height*0.3, y+height*0.5, text)
}

func drawEditBoxNum(vg *nanovgo.Context, text string, units string, x, y, width, height float32) {
	drawEditBoxBase(vg, x, y, width, height)

	var uw, _ = vg.TextBounds(0, 0, units)

	vg.FontSize(18.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 64))
	vg.TextAlign(nanovgo.AlignRight | nanovgo.AlignMiddle)
	vg.Text(x+width-height*0.3, y+height*0.5, units)

	vg.FontSize(20.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 128))
	vg.TextAlign(nanovgo.AlignRight | nanovgo.AlignMiddle)
	vg.Text(x+width-uw-height*0.5, y+height*0.5, text)
}

func drawCheckBox(vg *nanovgo.Context, text string, x, y, width, height float32) {
	vg.FontSize(18.0)
	vg.FontFace("sans")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 160))

	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
	vg.Text(x+28, y+height*0.5, text)

	var bg = vg.BoxGradient(x+1, y+float32(int(height*0.5))-9+1, 18, 18, 3, 3, nanovgo.RGBA(0, 0, 0, 32), nanovgo.RGBA(0, 0, 0, 92))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+float32(int(height*0.5))-9, 18, 18, 3)
	vg.FillPaint(bg)
	vg.Fill()

	vg.FontSize(40)
	vg.FontFace("icons")
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 128))
	vg.TextAlign(nanovgo.AlignCenter | nanovgo.AlignMiddle)
	vg.Text(x+9+2, y+height*0.5, cpToUTF8(IconCheck))
}

func drawButton(vg *nanovgo.Context, preIcon int, text string, x, y, width, height float32, color nanovgo.Color) {
	var cornerRadius float32 = 4.0
	var alpha uint8
	if isBlack(color) {
		alpha = 16
	} else {
		alpha = 32
	}
	var bg = vg.LinearGradient(x, y, x, y+height, nanovgo.RGBA(255, 255, 255, alpha), nanovgo.RGBA(0, 0, 0, alpha))
	vg.BeginPath()
	vg.RoundedRect(x+1, y+1, width-2, height-2, cornerRadius-1)
	if !isBlack(color) {
		vg.FillColor(color)
		vg.Fill()
	}
	vg.FillPaint(bg)
	vg.Fill()

	vg.BeginPath()
	vg.RoundedRect(x+0.5, y+0.5, width-1, height-1, cornerRadius-0.5)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 48))
	vg.Stroke()

	vg.FontSize(20.0)
	vg.FontFace("sans-bold")
	var tw, iw float32
	tw, _ = vg.TextBounds(0, 0, text)
	if preIcon != 0 {
		vg.FontSize(height * 1.3)
		vg.FontFace("icons")
		iw, _ = vg.TextBounds(0, 0, cpToUTF8(preIcon))
		iw += height * 0.15
	}

	if preIcon != 0 {
		vg.FontSize(height * 1.3)
		vg.FontFace("icons")
		vg.FillColor(nanovgo.RGBA(255, 255, 255, 96))
		vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
		vg.Text(x+width*0.5-tw*0.5-iw*0.75, y+height*0.5, cpToUTF8(preIcon))
	}

	vg.FontSize(20.0)
	vg.FontFace("sans-bold")
	vg.TextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)
	vg.FillColor(nanovgo.RGBA(0, 0, 0, 160))
	vg.Text(x+width*0.5-tw*0.5+iw*0.25, y+height*0.5-1, text)
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 160))
	vg.Text(x+width*0.5-tw*0.5+iw*0.25, y+height*0.5, text)
}

func drawSlider(vg *nanovgo.Context, pos, x, y, width, height float32) {
	var cy = y + float32(int(height*0.5))
	var kr = float32(int(height * 0.25))

	vg.Save()

	// Slot.
	var bg = vg.BoxGradient(x, cy-2+1, width, 4, 2, 2, nanovgo.RGBA(0, 0, 0, 32), nanovgo.RGBA(0, 0, 0, 128))
	vg.BeginPath()
	vg.RoundedRect(x, cy-2, width, 4, 2)
	vg.FillPaint(bg)
	vg.Fill()

	// Knob Shadow.
	bg = vg.RadialGradient(x+float32(int(pos*width)), cy+1, kr-3, kr+3, nanovgo.RGBA(0, 0, 0, 64), nanovgo.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(x+float32(int(pos*width))-kr-5, cy-kr-5, kr*2+5+5, kr*2+5+5+3)
	vg.Circle(x+float32(int(pos*width)), cy, kr)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(bg)
	vg.Fill()

	// Knob.
	var knob = vg.LinearGradient(x, cy-kr, x, cy+kr, nanovgo.RGBA(255, 255, 255, 16), nanovgo.RGBA(0, 0, 0, 16))
	vg.BeginPath()
	vg.Circle(x+float32(int(pos*width)), cy, kr-1)
	vg.FillColor(nanovgo.RGBA(40, 43, 48, 255))
	vg.Fill()
	vg.FillPaint(knob)
	vg.Fill()

	vg.BeginPath()
	vg.Circle(x+float32(int(pos*width)), cy, kr-0.5)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 92))
	vg.Stroke()

	vg.Restore()
}

func drawEyes(vg *nanovgo.Context, x, y, width, height, mx, my float32, t float64) {
	var ex = width * 0.23
	var ey = height * 0.5
	var lx = x + ex
	var ly = y + ey
	var rx = x + width - ex
	var ry = y + ey
	var dx, dy, d float32
	var br float32
	if ex < ey {
		br = ex * 0.5
	} else {
		br = ey * 0.5
	}
	var blink = 1 - float32(math.Pow(math.Sin(t*0.5), 200)*0.8)

	var bg = vg.LinearGradient(x, y+height*0.5, x+width*0.1, y+height, nanovgo.RGBA(0, 0, 0, 32), nanovgo.RGBA(0, 0, 0, 16))
	vg.BeginPath()
	vg.Ellipse(lx+3.0, ly+16.0, ex, ey)
	vg.Ellipse(rx+3.0, ry+16.0, ex, ey)
	vg.FillPaint(bg)
	vg.Fill()

	bg = vg.LinearGradient(x, y+height*0.25, x+width*0.1, y+height, nanovgo.RGBA(220, 220, 220, 255), nanovgo.RGBA(128, 128, 128, 255))
	vg.BeginPath()
	vg.Ellipse(lx, ly, ex, ey)
	vg.Ellipse(rx, ry, ex, ey)
	vg.FillPaint(bg)
	vg.Fill()

	dx = (mx - rx) / (ex * 10)
	dy = (my - ry) / (ey * 10)
	d = float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if d > 1.0 {
		dx /= d
		dy /= d
	}
	dx *= ex * 0.4
	dy *= ey * 0.5
	vg.BeginPath()
	vg.Ellipse(lx+dx, ly+dy+ey*0.25*(1-blink), br, br*blink)
	vg.FillColor(nanovgo.RGBA(32, 32, 32, 255))
	vg.Fill()

	dx = (mx - rx) / (ex * 10)
	dy = (my - ry) / (ey * 10)
	d = float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if d > 1.0 {
		dx /= d
		dy /= d
	}
	dx *= ex * 0.4
	dy *= ey * 0.5
	vg.BeginPath()
	vg.Ellipse(rx+dx, ry+dy+ey*0.25*(1-blink), br, br*blink)
	vg.FillColor(nanovgo.RGBA(32, 32, 32, 255))
	vg.Fill()

	var gloss = vg.RadialGradient(lx-ex*0.25, ly-ey*0.5, ex*0.1, ex*0.75, nanovgo.RGBA(255, 255, 255, 128), nanovgo.RGBA(255, 255, 255, 0))
	vg.BeginPath()
	vg.Ellipse(lx, ly, ex, ey)
	vg.FillPaint(gloss)
	vg.Fill()

	gloss = vg.RadialGradient(rx-ex*0.25, ry-ey*0.5, ex*0.1, ex*0.75, nanovgo.RGBA(255, 255, 255, 128), nanovgo.RGBA(255, 255, 255, 0))
	vg.BeginPath()
	vg.Ellipse(rx, ry, ex, ey)
	vg.FillPaint(gloss)
	vg.Fill()
}

func drawGraph(vg *nanovgo.Context, x, y, width, height float32, t float64) {
	var samples [6]float32
	var sx, sy [6]float32
	var dx = width / 5.0

	samples[0] = (1 + float32(math.Sin(t*1.2345+math.Cos(t*0.33457)*0.44))) * 0.5
	samples[1] = (1 + float32(math.Sin(t*0.68363+math.Cos(t*1.3)*1.55))) * 0.5
	samples[2] = (1 + float32(math.Sin(t*1.1642+math.Cos(t*0.33457)*1.24))) * 0.5
	samples[3] = (1 + float32(math.Sin(t*0.56345+math.Cos(t*1.63)*0.14))) * 0.5
	samples[4] = (1 + float32(math.Sin(t*1.6245+math.Cos(t*0.254)*0.3))) * 0.5
	samples[5] = (1 + float32(math.Sin(t*0.345+math.Cos(t*0.03)*0.6))) * 0.5

	for i := 0; i < 6; i++ {
		sx[i] = x + float32(i)*dx
		sy[i] = y + height*samples[i]*0.8
	}

	// Graph background.
	var bg = vg.LinearGradient(x, y, x, y+height, nanovgo.RGBA(0, 160, 192, 0), nanovgo.RGBA(0, 160, 192, 64))
	vg.BeginPath()
	vg.MoveTo(sx[0], sy[0])
	for i := 1; i < 6; i++ {
		vg.BezierTo(sx[i-1]+dx*0.5, sy[i-1], sx[i]-dx*0.5, sy[i], sx[i], sy[i])
	}
	vg.LineTo(x+width, y+height)
	vg.LineTo(x, y+height)
	vg.FillPaint(bg)
	vg.Fill()

	// Graph line.
	vg.BeginPath()
	vg.MoveTo(sx[0], sy[0]+2)
	for i := 1; i < 6; i++ {
		vg.BezierTo(sx[i-1]+dx*0.5, sy[i-1]+2, sx[i]-dx*0.5, sy[i]+2, sx[i], sy[i]+2)
	}
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 32))
	vg.StrokeWidth(3.0)
	vg.Stroke()

	vg.BeginPath()
	vg.MoveTo(sx[0], sy[0])
	for i := 1; i < 6; i++ {
		vg.BezierTo(sx[i-1]+dx*0.5, sy[i-1], sx[i]-dx*0.5, sy[i], sx[i], sy[i])
	}
	vg.StrokeColor(nanovgo.RGBA(0, 160, 192, 255))
	vg.StrokeWidth(3.0)
	vg.Stroke()

	// Graph sample pos.
	for i := 0; i < 6; i++ {
		bg = vg.RadialGradient(sx[i], sy[i]+2, 3.0, 8.0, nanovgo.RGBA(0, 0, 0, 32), nanovgo.RGBA(0, 0, 0, 0))
		vg.BeginPath()
		vg.Rect(sx[i]-10, sy[i]-10+2, 20, 20)
		vg.FillPaint(bg)
		vg.Fill()
	}

	vg.BeginPath()
	for i := 0; i < 6; i++ {
		vg.Circle(sx[i], sy[i], 4.0)
	}
	vg.FillColor(nanovgo.RGBA(0, 160, 192, 255))
	vg.Fill()
	vg.BeginPath()
	for i := 0; i < 6; i++ {
		vg.Circle(sx[i], sy[i], 2.0)
	}
	vg.FillColor(nanovgo.RGBA(220, 220, 220, 255))
	vg.Fill()

	vg.StrokeWidth(1.0)
}

func drawSpinner(vg *nanovgo.Context, cx, cy, r float32, t float64) {
	var a0 = 0.0 + float32(t*6)
	var a1 = float32(math.Pi + t*6)
	var r0 = r
	var r1 = r * 0.75
	var ax, ay, bx, by float32

	vg.Save()

	vg.BeginPath()
	vg.Arc(cx, cy, r0, a0, a1, nanovgo.CW)
	vg.Arc(cx, cy, r1, a1, a0, nanovgo.CCW)
	vg.ClosePath()
	ax = cx + float32(math.Cos(float64(a0))*float64(r0+r1)*0.5)
	ay = cy + float32(math.Sin(float64(a0))*float64(r0+r1)*0.5)
	bx = cx + float32(math.Cos(float64(a1))*float64(r0+r1)*0.5)
	by = cy + float32(math.Sin(float64(a1))*float64(r0+r1)*0.5)
	var paint = vg.LinearGradient(ax, ay, bx, by, nanovgo.RGBA(0, 0, 0, 0), nanovgo.RGBA(0, 0, 0, 128))
	vg.FillPaint(paint)
	vg.Fill()

	vg.Restore()
}

func drawThumbnails(vg *nanovgo.Context, x, y, width, height float32, images [12]*nanovgo.Image, t float64) {
	var cornerRadius float32 = 3.0
	var ix, iy, iw, ih float32
	var thumb float32 = 60.0
	var arry float32 = 30.5
	var imgw, imgh int
	var stackh = float32((len(images)/2))*(thumb+10) + 10
	var u = (1 + float32(math.Cos(t*0.5))) * 0.5
	var u2 = (1 - float32(math.Cos(t*0.2))) * 0.5
	var scrollh, dv float32

	vg.Save()

	// Drop shadow.
	var shadowPaint = vg.BoxGradient(x, y+4, width, height, cornerRadius*2, 20, nanovgo.RGBA(0, 0, 0, 128), nanovgo.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(x-10, y-10, width+20, height+30)
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(shadowPaint)
	vg.Fill()

	// Window.
	vg.BeginPath()
	vg.RoundedRect(x, y, width, height, cornerRadius)
	vg.MoveTo(x-10, y+arry)
	vg.LineTo(x+1, y+arry-11)
	vg.LineTo(x+1, y+arry+11)
	vg.FillColor(nanovgo.RGBA(200, 200, 200, 255))
	vg.Fill()

	vg.Save()
	vg.Scissor(x, y, width, height)
	vg.Translate(0, -(stackh-height)*u)

	dv = 1.0 / float32(len(images)-1)

	for i := 0; i < len(images); i++ {
		var tx, ty, v, a float32
		tx = x + 10
		ty = y + 10
		tx += float32(i%2) * (thumb + 10)
		ty += float32(i/2) * (thumb + 10)
		imgw, imgh = images[i].Size()
		if imgw < imgh {
			iw = thumb
			ih = iw * float32(imgh) / float32(imgw)
			ix = 0
			iy = -(ih - thumb) * 0.5
		} else {
			ih = thumb
			iw = ih * float32(imgw) / float32(imgh)
			ix = -(iw - thumb) * 0.5
			iy = 0
		}

		v = float32(i) * dv
		a = clampf((u2-v)/dv, 0, 1)

		if a < 1.0 {
			drawSpinner(vg, tx+thumb/2, ty+thumb/2, thumb*0.25, t)
		}

		var imgPaint = vg.ImagePattern(tx+ix, ty+iy, iw, ih, 0.0/180.0*math.Pi, images[i], a)
		vg.BeginPath()
		vg.RoundedRect(tx, ty, thumb, thumb, 5)
		vg.FillPaint(imgPaint)
		vg.Fill()

		shadowPaint = vg.BoxGradient(tx-1, ty, thumb+2, thumb+2, 5, 3, nanovgo.RGBA(0, 0, 0, 128), nanovgo.RGBA(0, 0, 0, 0))
		vg.BeginPath()
		vg.Rect(tx-5, ty-5, thumb+10, thumb+10)
		vg.RoundedRect(tx, ty, thumb, thumb, 6)
		vg.PathWinding(nanovgo.Hole)
		vg.FillPaint(shadowPaint)
		vg.Fill()

		vg.BeginPath()
		vg.RoundedRect(tx+0.5, ty+0.5, thumb-1, thumb-1, 4-0.5)
		vg.StrokeWidth(1.0)
		vg.StrokeColor(nanovgo.RGBA(255, 255, 255, 192))
		vg.Stroke()
	}
	vg.Restore()

	// Hide fades.
	var fadePaint = vg.LinearGradient(x, y, x, y+6, nanovgo.RGBA(200, 200, 200, 255), nanovgo.RGBA(200, 200, 200, 0))
	vg.BeginPath()
	vg.Rect(x+4, y, width-8, 6)
	vg.FillPaint(fadePaint)
	vg.Fill()

	fadePaint = vg.LinearGradient(x, y+height, x, y+height-6, nanovgo.RGBA(200, 200, 200, 255), nanovgo.RGBA(200, 200, 200, 0))
	vg.BeginPath()
	vg.Rect(x+4, y+height-6, width-8, 6)
	vg.FillPaint(fadePaint)
	vg.Fill()

	// Scroll bar.
	shadowPaint = vg.BoxGradient(x+width-12+1, y+4+1, 8, height-8, 3, 4, nanovgo.RGBA(0, 0, 0, 32), nanovgo.RGBA(0, 0, 0, 92))
	vg.BeginPath()
	vg.RoundedRect(x+width-12, y+4, 8, height-8, 3)
	vg.FillPaint(shadowPaint)
	vg.Fill()

	scrollh = (height / stackh) * (height - 8)
	shadowPaint = vg.BoxGradient(x+width-12-1, y+4+(height-8-scrollh)*u-1, 8, scrollh, 3, 4, nanovgo.RGBA(220, 220, 220, 255), nanovgo.RGBA(128, 128, 128, 255))
	vg.BeginPath()
	vg.RoundedRect(x+width-12+1, y+4+1+(height-8-scrollh)*u, 8-2, scrollh-2, 2)
	vg.FillPaint(shadowPaint)
	vg.Fill()

	vg.Restore()
}

func drawColorwheel(vg *nanovgo.Context, x, y, width, height float32, t float64) {
	var r0, r1, ax, ay, bx, by, cx, cy, aeps, r float32
	var hue = float32(math.Sin(t * 0.12))

	vg.Save()

	cx = x + width*0.5
	cy = y + height*0.5
	if width < height {
		r1 = width*0.5 - 5.0
	} else {
		r1 = height*0.5 - 5.0
	}
	r0 = r1 - 20.0
	aeps = 0.5 / r1 // Half a pixel arc length in radians (2pi cancels out).

	for i := 0; i < 6; i++ {
		var a0 = float32(i)/6.0*float32(math.Pi)*2.0 - aeps
		var a1 = float32(i+1.0)/6.0*float32(math.Pi)*2.0 + aeps
		vg.BeginPath()
		vg.Arc(cx, cy, r0, a0, a1, nanovgo.CW)
		vg.Arc(cx, cy, r1, a1, a0, nanovgo.CCW)
		vg.ClosePath()
		ax = cx + float32(math.Cos(float64(a0)))*(r0+r1)*0.5
		ay = cy + float32(math.Sin(float64(a0)))*(r0+r1)*0.5
		bx = cx + float32(math.Cos(float64(a1)))*(r0+r1)*0.5
		by = cy + float32(math.Sin(float64(a1)))*(r0+r1)*0.5
		var paint = vg.LinearGradient(ax, ay, bx, by, nanovgo.HSLA(a0/(math.Pi*2), 1.0, 0.55, 255), nanovgo.HSLA(a1/(math.Pi*2), 1.0, 0.55, 255))
		vg.FillPaint(paint)
		vg.Fill()
	}

	vg.BeginPath()
	vg.Circle(cx, cy, r0-0.5)
	vg.Circle(cx, cy, r1+0.5)
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 64))
	vg.StrokeWidth(1.0)
	vg.Stroke()

	// Selector.
	vg.Save()
	vg.Translate(cx, cy)
	vg.Rotate(hue * math.Pi * 2)

	// Marker on.
	vg.StrokeWidth(2.0)
	vg.BeginPath()
	vg.Rect(r0-1, -3, r1-r0+2, 6)
	vg.StrokeColor(nanovgo.RGBA(255, 255, 255, 192))
	vg.Stroke()

	var paint = vg.BoxGradient(r0-3, -5, r1-r0+6, 10, 2, 4, nanovgo.RGBA(0, 0, 0, 128), nanovgo.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(r0-2-10, -4-10, r1-r0+4+20, 8+20)
	vg.Rect(r0-2, -4, r1-r0+4, 8)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(paint)
	vg.Fill()

	// Center triangle.
	r = r0 - 6
	ax = float32(math.Cos(120.0/180.0*math.Pi)) * r
	ay = float32(math.Sin(120.0/180.0*math.Pi)) * r
	bx = float32(math.Cos(-120.0/180.0*math.Pi)) * r
	by = float32(math.Sin(-120.0/180.0*math.Pi)) * r
	vg.BeginPath()
	vg.MoveTo(r, 0)
	vg.LineTo(ax, ay)
	vg.LineTo(bx, by)
	vg.ClosePath()
	paint = vg.LinearGradient(r, 0, ax, ay, nanovgo.HSLA(hue, 1.0, 0.5, 255), nanovgo.RGBA(255, 255, 255, 255))
	vg.FillPaint(paint)
	vg.Fill()
	paint = vg.LinearGradient((r+ax)*0.5, (0+ay)*0.5, bx, by, nanovgo.RGBA(0, 0, 0, 0), nanovgo.RGBA(0, 0, 0, 255))
	vg.FillPaint(paint)
	vg.Fill()
	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 64))
	vg.Stroke()

	// Select circle on triangle.
	ax = float32(math.Cos(120.0/180.0*math.Pi)) * r * 0.3
	ay = float32(math.Sin(120.0/180.0*math.Pi)) * r * 0.4
	vg.StrokeWidth(2.0)
	vg.BeginPath()
	vg.Circle(ax, ay, 5)
	vg.StrokeColor(nanovgo.RGBA(255, 255, 255, 192))
	vg.Stroke()

	paint = vg.RadialGradient(ax, ay, 7, 9, nanovgo.RGBA(0, 0, 0, 64), nanovgo.RGBA(0, 0, 0, 0))
	vg.BeginPath()
	vg.Rect(ax-20, ay-20, 40, 40)
	vg.Circle(ax, ay, 7)
	vg.PathWinding(nanovgo.Hole)
	vg.FillPaint(paint)
	vg.Fill()

	vg.Restore()

	vg.Restore()
}

func drawLines(vg *nanovgo.Context, x, y, width, height float32, t float64) {
	var pad float32 = 5.0
	var s = width/9.0 - pad*2
	var pts [4 * 2]float32
	var fx, fy float32
	var joins = [3]nanovgo.LineJoin{nanovgo.Miter, nanovgo.RoundJoin, nanovgo.Bevel}
	var caps = [3]nanovgo.LineCap{nanovgo.Butt, nanovgo.RoundCap, nanovgo.Square}

	vg.Save()
	pts[0] = -s*0.25 + float32(math.Cos(t*0.3))*s*0.5
	pts[1] = float32(math.Sin(t*0.3)) * s * 0.5
	pts[2] = -s * 0.25
	pts[3] = 0
	pts[4] = s * 0.25
	pts[5] = 0
	pts[6] = s*0.25 + float32(math.Cos(-t*0.3))*s*0.5
	pts[7] = float32(math.Sin(-t*0.3)) * s * 0.5

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fx = x + s*0.5 + float32(i*3+j)/9.0*width + pad
			fy = y - s*0.5 + pad

			vg.LineCap(caps[i])
			vg.LineJoin(joins[j])

			vg.StrokeWidth(s * 0.3)
			vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 160))
			vg.BeginPath()
			vg.MoveTo(fx+pts[0], fy+pts[1])
			vg.LineTo(fx+pts[2], fy+pts[3])
			vg.LineTo(fx+pts[4], fy+pts[5])
			vg.LineTo(fx+pts[6], fy+pts[7])
			vg.Stroke()

			vg.LineCap(nanovgo.Butt)
			vg.LineJoin(nanovgo.Bevel)

			vg.StrokeWidth(1.0)
			vg.StrokeColor(nanovgo.RGBA(0, 192, 255, 255))
			vg.BeginPath()
			vg.MoveTo(fx+pts[0], fy+pts[1])
			vg.LineTo(fx+pts[2], fy+pts[3])
			vg.LineTo(fx+pts[4], fy+pts[5])
			vg.LineTo(fx+pts[6], fy+pts[7])
			vg.Stroke()
		}
	}

	vg.Restore()
}

func drawParagraph(vg *nanovgo.Context, x, y, width, height, mx, my float32) {
	// TODO: implement drawParagraph.
}

func drawWidths(vg *nanovgo.Context, x, y, width float32) {
	vg.Save()

	vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 255))

	for i := 0; i < 20; i++ {
		var w = (float32(i) + 0.5) * 0.1
		vg.StrokeWidth(w)
		vg.BeginPath()
		vg.MoveTo(x, y)
		vg.LineTo(x+width, y+width*0.3)
		vg.Stroke()
		y += 10
	}

	vg.Restore()
}

func drawCaps(vg *nanovgo.Context, x, y, width float32) {
	var caps = [3]nanovgo.LineCap{nanovgo.Butt, nanovgo.RoundCap, nanovgo.Square}
	var lineWidth float32 = 8.0

	vg.Save()

	vg.BeginPath()
	vg.Rect(x-lineWidth/2, y, width+lineWidth, 40)
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 32))
	vg.Fill()

	vg.BeginPath()
	vg.Rect(x, y, width, 40)
	vg.FillColor(nanovgo.RGBA(255, 255, 255, 32))
	vg.Fill()

	vg.StrokeWidth(lineWidth)
	for i := 0; i < 3; i++ {
		vg.LineCap(caps[i])
		vg.StrokeColor(nanovgo.RGBA(0, 0, 0, 255))
		vg.BeginPath()
		vg.MoveTo(x, y+float32(i)*10+5)
		vg.LineTo(x+width, y+float32(i)*10+5)
		vg.Stroke()
	}

	vg.Restore()
}

func drawScissor(vg *nanovgo.Context, x, y float32, t float64) {
	vg.Save()

	// Draw first rect and set scissor to it's area.
	vg.Translate(x, y)
	vg.Rotate(nanovgo.DegToRad(5))
	vg.BeginPath()
	vg.Rect(-20, -20, 60, 40)
	vg.FillColor(nanovgo.RGBA(255, 0, 0, 255))
	vg.Fill()
	vg.Scissor(-20, -20, 60, 40)

	// Draw second rectangle with offset and rotation.
	vg.Translate(40, 0)
	vg.Rotate(float32(t))

	// Draw the intended second rectangle without any scissoring.
	vg.Save()
	vg.ResetScissor()
	vg.BeginPath()
	vg.Rect(-20, -10, 60, 30)
	vg.FillColor(nanovgo.RGBA(255, 128, 0, 64))
	vg.Fill()
	vg.Restore()

	// Draw second rectangle with combined scissoring.
	vg.IntersectScissor(-20, -10, 60, 30)
	vg.BeginPath()
	vg.Rect(-20, -10, 60, 30)
	vg.FillColor(nanovgo.RGBA(255, 128, 0, 255))
	vg.Fill()

	vg.Restore()
}
