package skins

import (
	"image"
	"log"
)

// SkinExtractHead Return the 8x8 head of a 64x32 skin
func SkinExtractHead(img *image.RGBA) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 8, 16))

	if !AssertImageOpacity(img) {
		log.Fatalln("Image contains non-opaque pixels")
	}

	if !AssertAreaOpaque(img, LuantiHead1Front) {
		log.Fatalln("Head 1 front area is not opaque")
	}

	CopyArea(dst, img, LuantiHead1Front, image.Point{})
	CopyArea(dst, img, LuantiHead2Front, image.Point{})

	return dst
}
