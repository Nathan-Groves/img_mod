package Colors

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func PrintColors() {

	 var imageUrl string = ""

	fmt.Println("Enter name of image to printColor of")

	fmt.Scanln(&imageUrl)
	
  
	reader, err := os.Open(imageUrl)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

  
	outfile, err := os.Create("printColors.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer outfile.Close()


	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
      
			colors := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
			outfile.WriteString(colors)

			colors = ""
		}
	}
	
	fmt.Println("Processing Pixel Colors. Saving pixel output to printColors.txt" + ...."done")
}
