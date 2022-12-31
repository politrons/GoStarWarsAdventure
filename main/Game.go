package main

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
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

//Index of background image
var imageIndex = 0

//Collection of image of the whole game
var images = []string{
	"./assets/background.png",
	"./assets/background2.png",
}

//Collection of texts of the whole game
var gameTexts = []string{
	"aaaaaaaaaaaaaaaa",
	"bbbbbbbbbbbbbbbb",
}

//Collection of actions of the whole game
var imagesActions = [][]string{
	{"attack"},
	{"fly"},
}

type Actions struct {
	values []string
}

type ArrayFeature interface {
	contains(element string) bool
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Game")

	background := loadImage(images[imageIndex])
	backgroundImg := canvas.NewImageFromImage(background)
	backgroundImg.FillMode = canvas.ImageFillOriginal

	var sprite = image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)

	addLabel(sprite, 300, 300, gameTexts[imageIndex])

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
				actions := Actions{values: imagesActions[imageIndex]}
				if actions.contains(text) {
					imageIndex = imageIndex + 1
					background := loadImage(images[imageIndex])

					backgroundImg = canvas.NewImageFromImage(background)
					backgroundImg.FillMode = canvas.ImageFillOriginal

					sprite := image.NewRGBA(background.Bounds())
					addLabel(sprite, 300, 300, gameTexts[imageIndex])

					appendImage(sprite, backgroundImg, window)
				} else {
					sprite := image.NewRGBA(background.Bounds())
					addLabel(sprite, 400, 400, "Action not valid.")
					appendImage(sprite, backgroundImg, window)
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
Function to append the image into the container once we we create a [Raster] instance from
the original image.
*/
func appendImage(sprite *image.RGBA, backgroundImg *canvas.Image, window fyne.Window) {
	playerImg := canvas.NewRasterFromImage(sprite)
	c := container.New(layout.NewMaxLayout(), backgroundImg, playerImg)
	window.SetContent(c)
}

/**
Util function to check if an action is part of the array passed
*/
func (actions Actions) contains(element string) bool {
	for _, action := range actions.values {
		if action == element {
			return true
		}
	}
	return false
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
