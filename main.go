package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"strconv"

	"github.com/Kianrad/videocropanalyzer/models"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// countHorizontalCropOccurrences counts how many times each HorizontalCrop struct appears in the array
func countHorizontalCropOccurrences(crops []models.HorizontalCrop) map[models.HorizontalCrop]int {
	counts := make(map[models.HorizontalCrop]int)
	for _, crop := range crops {
		counts[crop]++
	}
	return counts
}

// countVerticalCropOccurrences counts how many times each VerticalCrop struct appears in the array
func countVerticalCropOccurrences(crops []models.VerticalCrop) map[models.VerticalCrop]int {
	counts := make(map[models.VerticalCrop]int)
	for _, crop := range crops {
		counts[crop]++
	}
	return counts
}

// extractFrameAsJPEG extracts a specific frame from the video and returns it as a JPEG image
func extractFrameAsJPEG(inFileName string, frameNum int64) io.Reader {
	buf := bytes.NewBuffer(nil)
	ffmpeg.LogCompiledCommand = false
	err := ffmpeg.Input(inFileName, ffmpeg.KwArgs{"ss": frameNum}).
		Output("pipe:", ffmpeg.KwArgs{"frames:v": 1, "q:v": 1, "f": "image2"}).
		WithOutput(buf).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

// getTotalFrameCount returns the total duration (in frames) of the video file
func getTotalFrameCount(videoPath string) (float64, error) {
	data, err := ffmpeg.Probe(videoPath, nil)
	if err != nil {
		return 0, err
	}
	probeData := models.FFProbeOutput{}
	err = json.Unmarshal([]byte(data), &probeData)
	if err != nil {
		return 0, err
	}

	duration, err := strconv.ParseFloat(probeData.Format.Duration, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %v", err)
	}
	return duration, err
}

// calculateMiddleFrames calculates the frame indices to be analyzed
func calculateMiddleFrames(totalFrames int64, numFrames int64) ([]int64, error) {
	if totalFrames == 0 || numFrames > totalFrames {
		return nil, errors.New("invalid input parameters")
	}

	partSize := totalFrames / 3
	firstFrame := partSize
	gapSize := partSize / numFrames

	frames := make([]int64, numFrames)
	frames[0] = firstFrame
	for i := int64(1); i < numFrames-1; i++ {
		frames[i] = firstFrame + i*gapSize
	}
	frames[numFrames-1] = partSize * 2

	return frames, nil
}

// detectCropValues extracts the cropping values for the top/bottom and left/right borders
func detectCropValues(buf io.Reader) (models.HorizontalCrop, models.VerticalCrop) {
	horizontalCrop := models.HorizontalCrop{}
	verticalCrop := models.VerticalCrop{}

	img, err := jpeg.Decode(buf)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	threshold := uint8(10)

	hcrop := 0
	wcrop := 0

	for y := 0; y < height; y++ {
		allBlack := true
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if uint8(r>>8) > threshold || uint8(g>>8) > threshold || uint8(b>>8) > threshold {
				allBlack = false
			}
		}
		if allBlack {
			hcrop++
		} else {
			if hcrop > 0 {
				horizontalCrop.Top = hcrop - 1
			}
			hcrop = 0
		}

		if y == height-1 && hcrop > 0 {
			horizontalCrop.Bottom = hcrop
		}
	}

	for x := 0; x < width; x++ {
		allBlack := true
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if uint8(r>>8) > threshold || uint8(g>>8) > threshold || uint8(b>>8) > threshold {
				allBlack = false
			}
		}
		if allBlack {
			wcrop++
		} else {
			if wcrop > 0 {
				verticalCrop.Left = wcrop - 1
			}
			wcrop = 0
		}

		if x == width-1 && wcrop > 0 {
			verticalCrop.Right = wcrop
		}
	}
	return horizontalCrop, verticalCrop
}

// getMaxHorizontalCrop finds the maximum Top and Bottom values among HorizontalCrop that occur more than minCount times
func getMaxHorizontalCrop(crops []models.HorizontalCrop, minCount int) models.HorizontalCrop {
	counts := countHorizontalCropOccurrences(crops)
	maxTop := -1
	maxBottom := -1

	for crop, count := range counts {
		if count >= minCount {
			if crop.Top > maxTop {
				maxTop = crop.Top
			}
			if crop.Bottom > maxBottom {
				maxBottom = crop.Bottom
			}
		}
	}

	return models.HorizontalCrop{
		Top:    maxTop,
		Bottom: maxBottom,
	}
}

// getMaxVerticalCrop finds the maximum Left and Right values among VerticalCrop that occur more than minCount times
func getMaxVerticalCrop(crops []models.VerticalCrop, minCount int) models.VerticalCrop {
	counts := countVerticalCropOccurrences(crops)
	maxLeft := -1
	maxRight := -1

	for crop, count := range counts {
		if count >= minCount {
			if crop.Left > maxLeft {
				maxLeft = crop.Left
			}
			if crop.Right > maxRight {
				maxRight = crop.Right
			}
		}
	}

	return models.VerticalCrop{
		Left:  maxLeft,
		Right: maxRight,
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: videocropanalyzer <video_file_path>")
		return
	}

	file := os.Args[1]
	total, err := getTotalFrameCount(file)
	if err != nil {
		fmt.Printf("Error getting total frames: %v\n", err)
		return
	}

	frames, err := calculateMiddleFrames(int64(total), 15)
	if err != nil {
		fmt.Printf("Error calculating middle frames: %v\n", err)
		return
	}

	horizontalCrops := make([]models.HorizontalCrop, 15)
	verticalCrops := make([]models.VerticalCrop, 15)
	result := models.CropValues{}

	for i, frame := range frames {
		buf := extractFrameAsJPEG(file, frame)
		horizontalCrops[i], verticalCrops[i] = detectCropValues(buf)
	}

	// Handle horizontal crops
	horizontalCounts := countHorizontalCropOccurrences(horizontalCrops)
	hHasValue := false
	for crop, count := range horizontalCounts {
		if count > 5 {
			result.Top = crop.Top
			result.Bottom = crop.Bottom
			hHasValue = true
			break
		}
	}

	if !hHasValue {
		maxCrop := getMaxHorizontalCrop(horizontalCrops, 3)
		result.Top = maxCrop.Top
		result.Bottom = maxCrop.Bottom
	}

	// Handle vertical crops
	verticalCounts := countVerticalCropOccurrences(verticalCrops)
	vHasValue := false
	for crop, count := range verticalCounts {
		if count > 5 {
			result.Left = crop.Left
			result.Right = crop.Right
			vHasValue = true
			break
		}
	}

	if !vHasValue {
		maxCrop := getMaxVerticalCrop(verticalCrops, 3)
		result.Left = maxCrop.Left
		result.Right = maxCrop.Right
	}

	obj, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error marshalling result to JSON: %v\n", err)
		return
	}

	fmt.Print(string(obj))
}
