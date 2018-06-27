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

	"github.com/beta/glfw"
	"github.com/beta/nanovgo"
	"github.com/go-gl/gl/v3.2-core/gl"
)

// DemoMSAA specifies whether to enable MSAA.
const DemoMSAA = false

func errorCallback(err glfw.Error, desc string) {
	fmt.Printf("GLFW error: %d: %s\n", int(err), desc)
}

var (
	blowup     bool
	screenshot bool
	premult    bool
)

func key(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierFlag) {
	if key == glfw.KeyEscape && action == glfw.Press {
		win.SetShouldClose(true)
	}
	if key == glfw.KeySpace && action == glfw.Press {
		blowup = !blowup
	}
	if key == glfw.KeyS && action == glfw.Press {
		screenshot = true
	}
	if key == glfw.KeyP && action == glfw.Press {
		premult = !premult
	}
}

func main() {
	var ctx = glfw.Init()
	if ctx == nil {
		panic("failed to init GLFW")
	}
	defer ctx.Terminate()

	var fps = NewPerfGraph(GraphRenderFPS, "Frame Time")
	var cpuGraph = NewPerfGraph(GraphRenderMS, "CPU Time")

	ctx.SetErrorCallback(errorCallback)

	ctx.WindowHint(glfw.ContextVersionMajor, 3)
	ctx.WindowHint(glfw.ContextVersionMinor, 2)
	ctx.WindowHintBool(glfw.OpenGLForwardCompat, true)
	ctx.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	ctx.WindowHintBool(glfw.OpenGLDebugContext, true)

	if DemoMSAA {
		ctx.WindowHint(glfw.Samples, 4)
	}

	var win = ctx.CreateWindow(1000, 600, "NanoVGo", nil, nil)
	if win == nil {
		panic("failed to create window")
	}
	defer win.Destroy()

	win.SetKeyCallback(key)
	ctx.MakeContextCurrent(win)

	var vg *nanovgo.Context
	if DemoMSAA {
		vg = nanovgo.CreateContext(nanovgo.StencilStrokes | nanovgo.Debug)
	} else {
		vg = nanovgo.CreateContext(nanovgo.Antialias | nanovgo.StencilStrokes | nanovgo.Debug)
	}
	if vg == nil {
		panic("failed to init NanoVGo")
	}
	defer vg.Delete()

	demo, err := loadDemoData(vg)
	if err != nil {
		panic("failed to load demo data")
	}

	ctx.SwapInterval(0)
	ctx.SetTime(0)
	var prevt = ctx.GetTime()
	var cpuTime float64

	// Use github.com/go-gl/gl for OpenGL APIs.
	err = gl.Init()
	if err != nil {
		panic("failed to init OpenGL")
	}

	for !win.ShouldClose() {
		var t = ctx.GetTime()
		var dt = t - prevt
		prevt = t

		var mx, my = win.GetCursorPos()
		var winWidth, winHeight = win.GetSize()
		var fbWidth, fbHeight = win.GetFramebufferSize()
		// Calculate pixel ratio for Hi-DPI devices.
		var pxRatio = float64(fbWidth) / float64(winWidth)

		gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
		if premult {
			gl.ClearColor(0, 0, 0, 0)
		} else {
			gl.ClearColor(0.3, 0.3, 0.32, 1.0)
		}
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

		vg.BeginFrame(float32(winWidth), float32(winHeight), float32(pxRatio))

		demo.render(vg, float32(mx), float32(my), float32(winWidth), float32(winHeight), t, blowup)

		fps.Render(vg, 5, 5)
		cpuGraph.Render(vg, 5+200+5, 5)

		vg.EndFrame()

		// Measure the CPU time taken excluding swap buffers (as the swap may
		// wait for GPU).
		cpuTime = ctx.GetTime() - t

		fps.Update(dt)
		cpuGraph.Update(cpuTime)

		if screenshot {
			screenshot = false
			saveScreenShot(fbWidth, fbHeight, premult, "dump.png")
		}

		win.SwapBuffers()
		ctx.PollEvents()
	}

	demo.free(vg)

	fmt.Printf("Average Frame Time: %.2f ms\n", fps.Average()*1000.0)
	fmt.Printf("          CPU Time: %.2f ms\n", cpuGraph.Average()*1000.0)
}
