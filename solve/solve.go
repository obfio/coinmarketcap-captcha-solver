package solve

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func SolveImage(img image.Image) int {
	// Split the image into puzzle piece image and target image, and save them
	puzzlePieceImage, targetImage, puzzlePieceMiddle := splitImages(img)
	saveImage(puzzlePieceImage, "./examples/piece.png")
	saveImage(targetImage, "./examples/target.png")
	fmt.Println(puzzlePieceMiddle)

	// Find the dominant row in the puzzle piece image
	dominantRow := findDominantRow(puzzlePieceImage, puzzlePieceMiddle)
	fmt.Println(dominantRow)

	// add 20 to dominantRow
	dominantRow += 20

	// make red pixel at `x,y` of puzzlePieceImage
	bounds := puzzlePieceImage.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	rgba.Set(puzzlePieceMiddle, dominantRow, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	saveImage(rgba, "./examples/piece.png")

	// get the place on the image where the shading is most defined
	shadedX := findShadedArea(targetImage, dominantRow)
	bounds = targetImage.Bounds()
	rgba = image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	rgba.Set(shadedX, dominantRow, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	saveImage(rgba, "./examples/target.png")
	return shadedX
}

func findShadedArea(img image.Image, row int) int {
	bounds := img.Bounds()
	xStart := -1
	minIntensity := math.MaxFloat64
	for x := bounds.Min.X + 60; x < bounds.Max.X; x++ {
		c := img.At(x, row)
		r, g, b, _ := c.RGBA()
		r8 := float64(r >> 8)
		g8 := float64(g >> 8)
		b8 := float64(b >> 8)
		intensity := r8 + g8 + b8
		if intensity < minIntensity {
			if intensity+60 < minIntensity {
				xStart = x
			}
			minIntensity = intensity
		}
	}
	return xStart
}

func findDominantRow(img image.Image, middle int) int {
	bounds := img.Bounds()
	rowPixelCount := make(map[int]int)
	maxPixels := 0
	dominantRow := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		r, g, b, _ := img.At(middle, y).RGBA()
		if r != 0 && g != 0 && b != 0 {
			rowPixelCount[y]++
			if rowPixelCount[y] > maxPixels {
				maxPixels = rowPixelCount[y]
				dominantRow = y
			}
		}
	}
	return dominantRow
}

func splitImages(img image.Image) (image.Image, image.Image, int) {
	bounds := img.Bounds()

	var splitX int
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		r, g, b, _ := img.At(x, 0).RGBA()
		if r != 0 && g != 0 && b != 0 {
			splitX = x
			break
		}
	}
	puzzleRect := image.Rect(bounds.Min.X, bounds.Min.Y, splitX, bounds.Max.Y)
	targetRect := image.Rect(splitX, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)

	puzzleImg := image.NewRGBA(puzzleRect)
	draw.Draw(puzzleImg, puzzleRect, img, bounds.Min, draw.Src)

	targetImg := image.NewRGBA(targetRect)
	draw.Draw(targetImg, targetRect, img, image.Point{X: splitX, Y: 0}, draw.Src)

	return puzzleImg, targetImg, splitX / 2
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
