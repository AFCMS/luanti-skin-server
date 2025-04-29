package skins

import "image"

// Luanti Head 1, opaque

var LuantiHead1Front = image.Rect(8, 8, 16, 16)
var LuantiHead1Top = image.Rect(8, 0, 16, 8)
var LuantiHead1Left = image.Rect(16, 8, 24, 16)
var LuantiHead1Right = image.Rect(0, 8, 8, 16)
var LuantiHead1Back = image.Rect(24, 8, 32, 16)
var LuantiHead1Bottom = image.Rect(16, 0, 24, 8)

var LuantiHead1 = []image.Rectangle{
	LuantiHead1Front,
	LuantiHead1Top,
	LuantiHead1Left,
	LuantiHead1Right,
	LuantiHead1Back,
	LuantiHead1Bottom,
}

// Luanti Head 2, transparent

var LuantiHead2Front = image.Rect(40, 8, 48, 16)
var LuantiHead2Top = image.Rect(40, 0, 48, 8)
var LuantiHead2Left = image.Rect(48, 8, 56, 16)
var LuantiHead2Right = image.Rect(32, 8, 40, 16)
var LuantiHead2Back = image.Rect(56, 8, 64, 16)
var LuantiHead2Bottom = image.Rect(48, 0, 56, 8)

var LuantiHead2 = []image.Rectangle{
	LuantiHead2Front,
	LuantiHead2Top,
	LuantiHead2Left,
	LuantiHead2Right,
	LuantiHead2Back,
	LuantiHead2Bottom,
}
