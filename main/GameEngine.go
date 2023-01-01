package main

import (
	"fmt"
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
	"strings"
)

//Index of background image
var gameLevel = 0

type Actions struct {
	values []string
}

type ActionFeature interface {
	containsAll(element string) bool
}

var window = app.New().NewWindow("StarWars Adventure")

//All containers used in the game and append in the
var boardContent = container.NewVBox()
var labelContent = container.NewVBox()
var inputContent = container.NewVBox()
var imageContent = container.NewVBox()

func main() {
	appendImage(createLevelImage())
	appendLabel(gameTexts[gameLevel])
	appendInput()

	boardContent.Add(imageContent)
	boardContent.Add(labelContent)
	boardContent.Add(inputContent)

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

/**
Function to obtain the user action over the level and determine if he pass to the next level or not.
*/
func processAction(input *widget.Entry) func() {
	return func() {
		userActions := input.Text
		log.Println("Action was:", userActions)
		if len(userActions) != 0 {
			levelActions := Actions{values: imagesActions[gameLevel]}
			if levelActions.containsAll(userActions) {
				gameLevel = gameLevel + 1
				levelImg := createLevelImage()
				appendImage(levelImg)
				appendLabel(gameTexts[gameLevel])
			} else {
				appendLabel("Wrong Action")
			}
		}
	}
}

func createLevelImage() *canvas.Image {
	background := loadImage(images[gameLevel])
	backgroundImg := canvas.NewImageFromImage(background)
	backgroundImg.FillMode = canvas.ImageFillOriginal
	return backgroundImg
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
func (actions Actions) containsAll(action string) bool {
	var actionCount = 0
	for _, userAction := range strings.Split(action, " ") {
		for _, action := range actions.values {
			if action == userAction {
				actionCount += 1
			}
		}
	}
	log.Println("Action words found ", actionCount)
	return len(actions.values) == actionCount
}
