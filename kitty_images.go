package main

import (
	"encoding/base64"
	"fmt"
	"image"
)

const (
	KITTY_CHUNK int = 4096
)

func DisplayImageSubcell(id, X, Y, Z int) {
	fmt.Printf("\x1b_Ga=p,i=%d,q=2,X=%d,Y=%d,C=1,z=%d;\x1b\\", id, X, Y, Z)
}

func DisplayImage(id int) {
	// Sane defaults: X,Y = 0 ; Z = 10
	fmt.Printf("\x1b_Ga=p,i=%d,q=2,X=0,Y=0,C=1,z=%d;\x1b\\", id, 10)
}

func ProvisionImage(img *image.RGBA, id int) {
	// Announce data type and size
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	imgBase64Str := base64.StdEncoding.EncodeToString(img.Pix)

	n := len(imgBase64Str)
	fmt.Printf("\x1B_Gt=d,f=32,q=2,") // data, RGBA, dont reply
	fmt.Printf("i=%d,s=%d,v=%d,m=1;", id, width, height)

	// Send data in chunks
	for sent := 0; sent < n; {
		size := n - sent
		if size > KITTY_CHUNK {
			size = KITTY_CHUNK
		}

		cont := 0
		if sent+size < n {
			cont = 1
		}

		if sent > 0 {
			fmt.Printf("\x1B_Gm=%d;", cont)
		}

		fmt.Printf(imgBase64Str[sent : sent+size])
		fmt.Printf("\x1B\\")
		sent += size
	}
}

func ClearImage(id int) {
	fmt.Printf("\x1b_Ga=d,d=i,i=%d;\x1b\\", id)
}

func DeleteImage(id int) {
	fmt.Printf("\x1b_Ga=d,d=I,i=%d;\x1b\\", id)
}
