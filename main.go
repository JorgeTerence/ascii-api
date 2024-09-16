package main

import (
	"fmt"
	"os"
	// "path/filepath"
	"reader"
)

const (
	KERNEL string = " .;coPO?S#"
)

// TODO: Download images from URL
// TODO: Output to file
// TODO: Add flags for scale and output file

func main() {
	inputFile := os.Args[1]
	// filepath.Ext(inputFile)
	luminanceData, w, h, err := reader.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err.Error())
		os.Exit(1)
	}

	var buffer []byte
	scaleY, scaleX := 8, 4

	h -= h % scaleY
	w -= w % scaleX

	for y := 0; y < h; y += scaleY {
		for x := 0; x < w; x += scaleX {
			newLuminance := sample(luminanceData, y, x, scaleY, scaleX)
			buffer = append(buffer, KERNEL[int(newLuminance*10-1)])
		}
		buffer = append(buffer, '\n')
	}

	os.WriteFile("output.txt", buffer, 0644)
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
