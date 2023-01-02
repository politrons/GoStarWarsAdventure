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
	"strings"
)

//Index of background image
var gameLevel = 0

type ActionFeature interface {
	containsAll(element string) bool
}

var window = app.New().NewWindow("StarWars Adventure")

//All containers used in the game and append in the
var boardContent = container.NewVBox()
var titleContent = ContentExtension{container.NewVBox()}
var labelContent = container.NewVBox()
var errorContent = container.NewVBox()
var inputContent = container.NewVBox()
var imageContent = ContentExtension{container.NewVBox()}

func main() {
	titleContent.appendImage(createLevelImage("./assets/logo.jpg"))
	imageContent.appendImage(createLevelImage(images[gameLevel]))
	appendLabel(gameTexts[gameLevel])
	appendInput()
	errorContent.Hide()

	boardContent.Add(titleContent.container)
	boardContent.Add(imageContent.container)
	boardContent.Add(labelContent)
	boardContent.Add(errorContent)
	boardContent.Add(inputContent)

	window.SetContent(boardContent)
	window.CenterOnScreen()
	window.ShowAndRun()

}

type ContentFeature interface {
	appendImage(image *canvas.Image)
}

type ContentExtension struct {
	container *fyne.Container
}

func (containerExt ContentExtension) appendImage(image *canvas.Image) {
	containerExt.container.RemoveAll()
	containerExt.container.Add(image)
}

func appendLabel(text string) {
	labelContent.RemoveAll()
	label := widget.NewLabel(text)
	labelContent.Add(label)
}

func appendErrorLabel(text string) {
	errorContent.Show()
	errorContent.RemoveAll()
	label := widget.NewLabel(text)
	errorContent.Add(label)
}

func appendInput() {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter action...")
	saveButton := widget.NewButton("Enter", processAction(input))
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
			cleanErrorLabel()
			levelActions := imagesActions[gameLevel]
			allActions, remains := levelActions.containsAll(userActions)
			if allActions {
				gameLevel = gameLevel + 1
				cleanInput()
				levelImg := createLevelImage(images[gameLevel])
				imageContent.appendImage(levelImg)
				appendLabel(gameTexts[gameLevel])
			} else {
				appendErrorLabel(fmt.Sprintf("Wrong Action. You're missing %d actions", remains))
			}
		}
	}
}

func createLevelImage(imagePath string) *canvas.Image {
	background := loadImage(imagePath)
	backgroundImg := canvas.NewImageFromImage(background)
	backgroundImg.FillMode = canvas.ImageFillOriginal
	return backgroundImg
}

func cleanErrorLabel() {
	errorContent.RemoveAll()
	errorContent.Hide()
}

func cleanInput() {
	inputContent.RemoveAll()
	appendInput()
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
func (actions LevelAction) containsAll(action string) (bool, int) {
	var actionCount = 0
	for _, userAction := range strings.Split(action, " ") {
		for _, action := range actions.actions {
			if action == userAction {
				actionCount += 1
			}
		}
	}
	log.Println("Action words found ", actionCount)
	return actions.minActions <= actionCount, actions.minActions - actionCount
}
