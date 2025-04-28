package utils

import (
	"image"
	"image/draw"
	"log"
)

// Luanti Head 1, opaque

var LuantiHead1Front = image.Rect(8, 8, 16, 16)
var LuantiHead1Top = image.Rect(8, 0, 16, 8)

// Luanti Head 2, transparent

var LuantiHead2Front = image.Rect(40, 8, 48, 16)
var LuantiHead2Top = image.Rect(40, 0, 48, 8)

func ImageToRGBA(img image.Image) (image.RGBA, error) {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return *rgba, nil
}

// AssertAreaOpaque Check if the area of a RGBA image is fully opaque
func AssertAreaOpaque(img *image.RGBA, area image.Rectangle) bool {
	bounds := img.Bounds()
	if area.Min.X < bounds.Min.X || area.Min.Y < bounds.Min.Y ||
		area.Max.X > bounds.Max.X || area.Max.Y > bounds.Max.Y {
		return false
	}

	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a == 0 {
				return false
			}
		}
	}

	return true
}

// AssertAreaOpacity Check if in the area of a image pixels are either fully opaque or fully transparent
func AssertAreaOpacity(img *image.RGBA, area image.Rectangle) bool {
	bounds := img.Bounds()
	if area.Min.X < bounds.Min.X || area.Min.Y < bounds.Min.Y ||
		area.Max.X > bounds.Max.X || area.Max.Y > bounds.Max.Y {
		return false
	}

	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0 && a != 0xffff {
				return false
			}
		}
	}

	return true
}

func CopyArea(dst *image.RGBA, src *image.RGBA, area image.Rectangle, destPoint image.Point) {
	destRect := area.Sub(area.Min).Add(destPoint)
	draw.Draw(dst, destRect, src, area.Min, draw.Over)
}

// SkinExtractHead Return the 8x8 head of a 64x32 skin
func SkinExtractHead(img *image.RGBA) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 8, 16))

	r := AssertAreaOpaque(img, LuantiHead1Front)
	if !r {
		log.Fatalln("Head 1 front area is not opaque")
	}
	r2 := AssertAreaOpacity(img, LuantiHead2Front)
	if !r2 {
		log.Fatalln("Head 2 front area")
	}

	CopyArea(dst, img, LuantiHead1Front, image.Point{})
	CopyArea(dst, img, LuantiHead2Front, image.Point{})

	return dst
}
