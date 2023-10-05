package Text

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func PrintImageText() {
	const W = 500
	const H = 300

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give me a number")

	fmt.Print("-> ")

	text, _ := reader.ReadString('\n')

	// Create a temporary file and write the byte slice to it
	tempFile, err := ioutil.TempFile("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored(text, W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	fmt.Println("Saving labeled image to coloredWord.png ... done")
	dc.SavePNG("coloredWord.png")

}
