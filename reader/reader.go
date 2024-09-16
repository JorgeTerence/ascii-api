package reader

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([][]float64, int, int, error) {
	if !filepath.IsAbs(path) {
		var err error
		path, err = filepath.Abs(filepath.Clean(path))

		if err != nil {
			return nil, 0, 0, err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var luminanceData [][]float64
	for y := 0; y < height; y++ {
		var row []float64
		for x := 0; x < width; x++ {
			row = append(row, luminance(img.At(x, y).RGBA()))
		}
		luminanceData = append(luminanceData, row)
	}

	return luminanceData, width, height, nil
}

func luminance(r, g, b, a uint32) float64 {
	return (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / (float64(a))
}
