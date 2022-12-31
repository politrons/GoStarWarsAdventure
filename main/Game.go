package main

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"golang.org/x/image/font"
	_ "golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	_ "golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	_ "golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
)

func loadImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}

	imgData, err := png.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return imgData.(image.Image)
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Game")

	background := loadImage("./assets/background.png")
	background2 := loadImage("./assets/background2.png")

	backgroundImg := canvas.NewImageFromImage(background)
	backgroundImg.FillMode = canvas.ImageFillOriginal

	var sprite = image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)

	addLabel(sprite, 300, 300, "hello world bla bla bla blA,bla,")

	c := container.New(layout.NewMaxLayout(), backgroundImg, playerImg)

	window.SetContent(c)

	go func() {

		// To create dynamic array
		arr := make([]string, 0)
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter Text: ")
			// Scans a line from Stdin(Console)
			scanner.Scan()
			// Holds the string that scanned
			text := scanner.Text()
			if len(text) != 0 {

				if text == "attack" {
					backgroundImg = canvas.NewImageFromImage(background2)
					backgroundImg.FillMode = canvas.ImageFillOriginal

					sprite = image.NewRGBA(background.Bounds())
					addLabel(sprite, 300, 300, "second text bla bla bla bla bla")

					playerImg = canvas.NewRasterFromImage(sprite)

					c = container.New(layout.NewMaxLayout(), backgroundImg, playerImg)
					window.SetContent(c)
				}

				fmt.Println(text)
				arr = append(arr, text)
			} else {
				break
			}

		}
		// Use collected inputs
		fmt.Println(arr)

	}()

	window.CenterOnScreen()
	window.ShowAndRun()

}

/**
Function to add label into image.
We use [Drawer] type to use [DrawString] implementation to add label text
into the current image.
*/
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
