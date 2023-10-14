package main

import (
	"fmt"
	"image"
)

func main() {
	height, width := 256, 512
	image := image.NewRGBA(image.Rect(0, 0, height, width))
	fmt.Println(image)
}
