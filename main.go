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
			"fort.jpg":             "https://plus.unsplash.com/premium_photo-1661930618375-aafabc2bf3e7?q=80&w=2599&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"park.jpg":             "https://images.unsplash.com/photo-1519331379826-f10be5486c6f?q=80&w=2670&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"mall.jpg":             "https://images.unsplash.com/photo-1580793241553-e9f1cce181af?q=80&w=2664&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"temple.jpg":           "https://images.unsplash.com/photo-1721532867177-edaf8cfe686b?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"church.jpg":           "https://images.unsplash.com/photo-1465848059293-208e11dfea17?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"mosque.jpg":           "https://images.unsplash.com/photo-1575682631529-7d47334f022f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"museum.jpg":           "https://images.unsplash.com/photo-1562754193-ba39a22c110b?q=80&w=2040&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"monument.jpg":         "https://images.unsplash.com/photo-1710822334460-32dbfd4d5d5f?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"cave.jpg":             "https://images.unsplash.com/photo-1550992402-9b1fc58fd76d?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"palace.jpg":           "https://images.unsplash.com/photo-1524229321985-1e1989075d9b?q=80&w=2107&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"hill.jpg":             "https://images.unsplash.com/photo-1476988186444-a7189cf07b3f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"lake.jpg":             "https://images.unsplash.com/photo-1501785888041-af3ef285b470?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"river.jpg":            "https://images.unsplash.com/photo-1519852476561-ec618b0183ba?q=80&w=2056&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"beach.jpg":            "https://images.unsplash.com/photo-1424581342241-2b1aba4d3462?q=80&w=2000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"waterfall.jpg":        "https://images.unsplash.com/photo-1610805177214-885738d255f1?q=80&w=2128&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"dam.jpg":              "https://images.unsplash.com/photo-1570106230673-3bab9f2f3c63?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"forest.jpg":           "https://images.unsplash.com/photo-1686890363911-635fb164ab96?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"jungle-safari.jpg":    "https://images.unsplash.com/photo-1656828059237-add66db82a2b?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"botanical-garden.jpg": "https://images.unsplash.com/photo-1598002582975-6ea1f15c7027?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"zoo.jpg":              "https://images.unsplash.com/photo-1603039529403-6ec390efcf4e?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"aquarium.jpg":         "https://images.unsplash.com/photo-1580140404772-decde5a9cc9f?q=80&w=2071&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"mountain.jpg":         "https://images.unsplash.com/photo-1454496522488-7a8e488e8606?q=80&w=2076&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"desert.jpg":           "https://images.unsplash.com/photo-1527519135413-1e146b552e10?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"valley.jpg":           "https://images.unsplash.com/photo-1468901184895-0cec1ee34ff5?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"cable-car.jpg":        "https://images.unsplash.com/photo-1615639394567-1a6f222651b1?q=80&w=2106&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"trek.jpg":             "https://images.unsplash.com/uploads/141148589884100082977/a816dbd7?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"adventure.jpg":        "https://images.unsplash.com/photo-1618083707368-b3823daa2726?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"heritage.jpg":         "https://images.unsplash.com/photo-1616606484004-5ef3cc46e39d?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"cityscape.jpg":        "https://cdn.pixabay.com/photo/2021/05/10/05/27/mumbai-6242623_960_720.jpg",
			"village.jpg":          "https://images.unsplash.com/photo-1647184223407-ef8273a6822c?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"island.jpg":           "https://images.unsplash.com/photo-1559128010-7c1ad6e1b6a5?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"sunrise.jpg":          "https://images.unsplash.com/photo-1484766280341-87861644c80d?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"sunset.jpg":           "https://images.unsplash.com/photo-1496614932623-0a3a9743552e?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"roadtrip.jpg":         "https://images.unsplash.com/photo-1469854523086-cc02fe5d8800?q=80&w=2021&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"festival.jpg":         "https://images.unsplash.com/photo-1468234847176-28606331216a?q=80&w=2077&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			"market.jpg":           "https://images.unsplash.com/photo-1572402123736-c79526db405a?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
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
