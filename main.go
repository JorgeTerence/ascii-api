package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

const (
	KERNEL string = " .-+*o9SYM#$"
)

func main() {

	img, err := openImage("./nois.jpg")
	if err != nil {
		os.Exit(1)
	}

	pixels, err := getPixels(img)
	if err != nil {
		os.Exit(1)
	}

	downscale := make([][]float64, 0)
	// factor := 2
	scaleY, scaleX := 3, 2

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	// if height%3 == 0 {
	// 	height++
	// }
	// if width%3 == 0 {
	// 	height--
	// }
	if height%2 == 0 {
		height++
	}
	if width%2 == 0 {
		height--
	}

	// TODO: avoid downscale sample overflowing the matrix
	// TODO: using a scale factor (x + y) instead of hardcoding values

	for y := 0; y < height; y += scaleY {
		row := make([]float64, 0)
		for x := 0; x < width; x += scaleX {
			row = append(row, (pixels[y][x].luminance()+pixels[y][x+1].luminance()+pixels[y+1][x].luminance()+pixels[y+1][x+1].luminance()+pixels[y+2][x].luminance()+pixels[y+2][x+1].luminance())/ float64(scaleY * scaleX))
		}
		downscale = append(downscale, row)
	}

	for _, row := range downscale {
		for _, p := range row {
			fmt.Print(string(KERNEL[int(p*10)]))
		}
		fmt.Print("\n")
	}
}

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}

	return img, nil
}

func getPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}

func (p Pixel) luminance() float64 {
	return float64(p.R+p.G+p.B) / 765
}

// func avg(m [][]Pixel, x, y int) Pixel {

// }
