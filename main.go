package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/manyids2/govips-manyids2/vips"
)

func main() {
	width, height := 256, 512
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{128, 128, 128, 128})
		}
	}

	fmt.Printf("img: %d, %d\n", img.Bounds().Size().X, img.Bounds().Size().Y)

	id := 150
	ProvisionImage(img, id)
	DisplayImage(id)
	time.Sleep(1000 * time.Millisecond)
	DeleteImage(id)

	vi, _ := vips.MemoryRGBA(img, width, height)
	fmt.Printf("vi: %d, %d\n", vi.Width(), vi.Height())

	vi.Resize(0.5, vips.KernelNearest)
	fmt.Printf("vi: %d, %d\n", vi.Width(), vi.Height())

	rsz := image.NewRGBA(image.Rect(0, 0, vi.Width(), vi.Height()))
	rsz.Pix, _ = vi.ToBytes()
	fmt.Printf("rsz: %d, %d\n", rsz.Bounds().Size().X, rsz.Bounds().Size().Y)

	id = 300
	ProvisionImage(rsz, id)
	DisplayImage(id)
	time.Sleep(1000 * time.Millisecond)
	DeleteImage(id)
}
