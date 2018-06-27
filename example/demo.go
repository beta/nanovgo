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

import "github.com/beta/nanovgo"

func loadDemoData(vg *nanovgo.Context) (*demoData, error) {
	// TODO: implement loadDemoData.
	return nil, nil
}

type demoData struct {
	fontNormal, fontBold, fontIcons, fontEmoji *nanovgo.Font
	images                                     [12]nanovgo.Image
}

func (data *demoData) free() {
	// TODO: implement free.
}

func (data *demoData) render(vg *nanovgo.Context, mx, my, width, height, t float64, blowup bool) {
	// TODO: implement render.
}

func saveScreenShot(width, height int, premult bool, name string) {
	// TODO: implement saveScreenShot.
}