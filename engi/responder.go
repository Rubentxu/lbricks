// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engi

type Responder interface {
	Render()
	Resize(width, height int)
	Preload()
	Setup()
	Close()
	Update(dt float32)
	Mouse(x, y float32, action MouseAction)
	Scroll(amount float32)
	Key(key Key, modifier Modifier, action KeyAction)
	Type(char rune)
}
