package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/manyids2/govips-manyids2/vips"
)

// Print to terminal using kitty protocol
func kittyPrint(img *image.RGBA, id int) {
	ProvisionImage(img, id)
	DisplayImage(id)
	time.Sleep(1000 * time.Millisecond)
	DeleteImage(id)
}

// Create random patterned image.RGBA
func createImageRGBA(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// First 1/4th dark blue
	for x := 0; x < width/4; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{0, 0, 128, 128})
		}
	}
	// Rest 3/4th dark red
	for x := width / 4; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{128, 0, 0, 128})
		}
	}
	return img
}

func RescaleImage(img *image.RGBA, scale float64, kernel vips.Kernel, do_copy bool) *image.RGBA {
	// Convert to *C.VipsImage
	vi, _ := vips.MemoryRGBA(img, do_copy)

	// Do the transforms
	vi.Resize(scale, kernel)

	// Convert back to image.RGBA
	rsz := image.NewRGBA(image.Rect(0, 0, vi.Width(), vi.Height()))
	rsz.Pix, _ = vi.ToBytes()

	return rsz
}

func main() {
	// Create random image
	width, height := 256, 512
	img := createImageRGBA(width, height)

	// Print to check
	fmt.Printf("img: %d, %d\n", img.Bounds().Size().X, img.Bounds().Size().Y)
	kittyPrint(img, 150)

	// Resize the image, with copy
	scale := 3.0
	do_copy := true
	rsz := RescaleImage(img, scale, vips.KernelNearest, do_copy)

	// Print to check
	fmt.Printf("img: %d, %d\n", img.Bounds().Size().X, img.Bounds().Size().Y)
	fmt.Printf("rsz: %d, %d\n", rsz.Bounds().Size().X, rsz.Bounds().Size().Y)
	kittyPrint(rsz, 150)
}
