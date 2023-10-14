package main

import (
	"fmt"
	"image"

	"github.com/manyids2/govips-manyids2/vips"
)

func main() {
	width, height := 256, 512
	img := image.NewRGBA(image.Rect(0, 0, height, width))
	fmt.Printf("img: %d, %d\n", img.Bounds().Size().X, img.Bounds().Size().Y)

	vi, _ := vips.MemoryRGBA(img, width, height)
	vi.Resize(0.5, vips.KernelNearest)
	fmt.Printf("vi: %d, %d\n", vi.Width(), vi.Height())

	rsz := image.NewRGBA(image.Rect(0, 0, vi.Width(), vi.Height()))
	rsz.Pix = vi.Buffer()
	fmt.Printf("rsz: %d, %d\n", rsz.Bounds().Size().X, rsz.Bounds().Size().Y)
}
