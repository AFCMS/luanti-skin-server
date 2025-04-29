package skins

import (
	"image"
	"image/draw"
)

// ImageToRGBA converts an [image.Image] to [image.RGBA]
func ImageToRGBA(img image.Image) (image.RGBA, error) {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return *rgba, nil
}

// CopyArea copies a rectangular area from src to dst at the specified destination point
func CopyArea(dst *image.RGBA, src *image.RGBA, area image.Rectangle, destPoint image.Point) {
	destRect := area.Sub(area.Min).Add(destPoint)
	draw.Draw(dst, destRect, src, area.Min, draw.Over)
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

// AssertImageOpacity Check if a image pixels are either fully opaque or fully transparent
func AssertImageOpacity(img *image.RGBA) bool {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0 && a != 0xffff {
				return false
			}
		}
	}

	return true
}
