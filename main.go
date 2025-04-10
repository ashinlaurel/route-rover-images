package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const (
	width   = 400
	height  = 250
	quality = 75
)

type ImageCategory struct {
	Name   string
	Images map[string]string // filename -> Unsplash URL
}

var categories = []ImageCategory{
	{
		Name: "place",
		Images: map[string]string{
			"fort.jpg": "https://plus.unsplash.com/premium_photo-1661930618375-aafabc2bf3e7?q=80&w=2599&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"park.jpg": "https://images.unsplash.com/photo-1519331379826-f10be5486c6f?q=80&w=2670&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"mall.jpg": "https://images.unsplash.com/photo-1580793241553-e9f1cce181af?q=80&w=2664&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			// ... add others here
		},
	},
	{
		Name: "food",
		Images: map[string]string{
			"street-food.jpg": "https://images.unsplash.com/photo-abc1",
			"dhaba.jpg":       "https://images.unsplash.com/photo-abc2",
			// ... add others here
		},
	},
}

func downloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	img, err := imaging.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return img, nil
}

func resizeAndSaveImage(img image.Image, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	resized := imaging.Resize(img, width, height, imaging.Lanczos)

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer out.Close()

	return jpeg.Encode(out, resized, &jpeg.Options{Quality: quality})
}

func processImage(category, name, url string) error {
	img, err := downloadImage(url)
	if err != nil {
		return fmt.Errorf("error downloading %s: %v", name, err)
	}

	outputPath := filepath.Join(category, name)
	if err := resizeAndSaveImage(img, outputPath); err != nil {
		return fmt.Errorf("error saving %s: %v", name, err)
	}

	fmt.Printf("✔ Processed: %s/%s\n", category, name)
	return nil
}

func main() {
	fmt.Println("Starting image download and resize...")

	for _, category := range categories {
		fmt.Printf("\n🗂 Category: %s\n", category.Name)
		for name, url := range category.Images {
			if err := processImage(category.Name, name, url); err != nil {
				fmt.Printf("⚠ %v\n", err)
			}
		}
	}

	fmt.Println("\n✅ All done!")
}
