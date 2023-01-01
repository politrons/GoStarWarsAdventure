package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "golang.org/x/image/font"
	_ "golang.org/x/image/font/basicfont"
	_ "golang.org/x/image/math/fixed"
	"image"
	"image/jpeg"
	"log"
	"os"
)

//Index of background image
var gameLevel = 0

type Actions struct {
	values []string
}

type ArrayFeature interface {
	contains(element string) bool
}

var window = app.New().NewWindow("StarWars Adventure")

var boardContent = container.NewVBox()
var labelContent = container.NewVBox()
var inputContent = container.NewVBox()
var imageContent = container.NewVBox()

func main() {

	background := loadImage(images[gameLevel])
	backgroundImg := canvas.NewImageFromImage(background)

	backgroundImg.SetMinSize(fyne.Size{800, 800})
	backgroundImg.FillMode = canvas.ImageFillOriginal

	//var sprite = image.NewRGBA(background.Bounds())
	//appendLogic(sprite, backgroundImg, gameTexts[gameLevel], window)

	appendImage(backgroundImg)
	appendLabel(gameTexts[gameLevel])
	appendInput()

	boardContent.Add(imageContent)
	boardContent.Add(labelContent)
	boardContent.Add(inputContent)

	//c := container.NewGridWithRows(2, backgroundImg, open, close, insert, playerImg)

	window.SetContent(boardContent)

	window.CenterOnScreen()
	window.ShowAndRun()

}

func appendImage(backgroundImg *canvas.Image) {
	imageContent.RemoveAll()
	imageContent.Add(backgroundImg)
}

func appendLabel(text string) {
	labelContent.RemoveAll()
	labelContent.Add(widget.NewLabel(text))
}

func appendInput() {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter action...")
	saveButton := widget.NewButton("Save", processAction(input))
	inputContent.RemoveAll()
	inputContent.Add(input)
	inputContent.Add(saveButton)
}

//
///**
//Function to append the image into the container once we we create a [Raster] instance from
//the original image.
//*/
//func appendLogic(sprite *image.RGBA, backgroundImg *canvas.Image, text string, window fyne.Window) {
//	//playerImg := canvas.NewRasterFromImage(sprite)
//
//	label := widget.NewLabel(text)
//
//	labelContent := container.NewVBox(label)
//
//	input := widget.NewEntry()
//	input.SetPlaceHolder("Enter action...")
//
//	saveButton := widget.NewButton("Save", processAction(input))
//
//	imageContent := container.NewVBox(backgroundImg)
//
//	inputContent := container.NewVBox(input, saveButton)
//
//	imageContent.Add(labelContent)
//	imageContent.Add(inputContent)
//
//	//c := container.NewGridWithRows(2, backgroundImg, open, close, insert, playerImg)
//
//	window.SetContent(imageContent)
//}

func processAction(input *widget.Entry) func() {
	return func() {
		text := input.Text
		log.Println("Action was:", text)
		if len(text) != 0 {
			actions := Actions{values: imagesActions[gameLevel]}
			if actions.contains(text) {
				gameLevel = gameLevel + 1
				background := loadImage(images[gameLevel])
				backgroundImg := canvas.NewImageFromImage(background)
				backgroundImg.FillMode = canvas.ImageFillOriginal

				appendImage(backgroundImg)
				appendLabel(gameTexts[gameLevel])
				//appendLogic(sprite, backgroundImg, gameTexts[gameLevel], window)
			} else {
				appendLabel("Wrong Action")
			}
		}
	}
}

/**
Function to load a new image from a path using [os] package to read an create [File],
and [jpeg] to Decode from a file into an [Image] instance .
Using [defer] we guarantee that after we load the file image is close,
before we end the function.
*/
func loadImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(imgFile)
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}

	imgData, err := jpeg.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return imgData.(image.Image)
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
