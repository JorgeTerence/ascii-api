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
	img, err := openImage("./selfie-piracicaba.jpg")
	if err != nil {
		os.Exit(1)
	}

	pixels, err := getPixels(img)
	if err != nil {
		os.Exit(1)
	}

	downscale := make([][]float64, 0)
	scaleY, scaleX := 6, 4

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	if height%scaleY != 0 {
		height -= height % scaleY
	}

	if width%scaleX != 0 {
		width -= width % scaleX
	}

	fmt.Printf("Image scale: %d x %d\n", bounds.Max.X, bounds.Max.Y)
	fmt.Printf("Matrix scale %d x %d\n", width, height)

	for y := 0; y < height; y += scaleY {
		row := make([]float64, 0)
		for x := 0; x < width; x += scaleX {
			newLuminance := 0.0

			for sY := y; sY < y+scaleY; sY++ {
				for sX := x; sX < x+scaleX; sX++ {
					newLuminance += pixels[sY][sX].luminance()
				}
			}

			row = append(row, newLuminance/float64(scaleY*scaleX))
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
