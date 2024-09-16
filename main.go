package main

import (
	"fmt"
	"os"
	"reader"
)

const (
	KERNEL string = " .;coPO?S#"
)

// TODO: Download images from URL
// TODO: Output to file
// TODO: Add flags for scale and output file

func main() {
	luminanceData, w, h, err := reader.ReadFile("./nois.jpg")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err.Error())
		os.Exit(1)
	}

	downscale := make([][]float64, 0)
	scaleY, scaleX := 8, 4

	h -= h % scaleY
	w -= w % scaleX

	for y := 0; y < h; y += scaleY {
		row := make([]float64, 0)
		for x := 0; x < w; x += scaleX {
			newLuminance := sample(luminanceData, y, x, scaleY, scaleX)
			row = append(row, newLuminance)
		}
		downscale = append(downscale, row)
	}

	for _, row := range downscale {
		for _, p := range row {
			fmt.Print(string(KERNEL[int(p*10-1)]))
		}
		fmt.Print("\n")
	}
}

func sample(data [][]float64, y, x, h, w int) float64 {
	var sum float64
	for _, row := range data[y : y+h] {
		for _, v := range row[x : x+w] {
			sum += v
		}
	}
	return sum / float64(h*w)
}
