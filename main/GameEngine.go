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
	"time"
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
	titleContent.appendImage(createImage("./assets/logo.jpg"))
	imageContent.appendImage(createImage(images[gameLevel]))
	appendLabel(gameTexts[gameLevel])
	appendInput()
	errorContent.Hide()

	boardContent.Add(titleContent.container)
	boardContent.Add(imageContent.container)
	boardContent.Add(labelContent)
	boardContent.Add(errorContent)
	boardContent.Add(inputContent)

	go renderMiddleLevelGame(gameLevel)
	go renderGameOver(gameLevel)

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
High order Function to obtain the user action over the level and determine if he pass to the next level or not.
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
				levelImg := createImage(images[gameLevel])
				imageContent.appendImage(levelImg)
				appendLabel(gameTexts[gameLevel])
				go renderMiddleLevelGame(gameLevel)
				go renderGameOver(gameLevel)
			} else {
				appendErrorLabel(fmt.Sprintf("Wrong Action. You're missing %d actions", remains))
			}
		}
	}
}

func createImage(imagePath string) *canvas.Image {
	gameImage := loadImage(imagePath)
	gameImageRender := canvas.NewImageFromImage(gameImage)
	gameImageRender.FillMode = canvas.ImageFillOriginal
	return gameImageRender
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
Function to be run [async] during the middle of the level to give a hint.
We use [recover] operator to handle possible panic runtime errors.
*/
func renderMiddleLevelGame(asyncGameLevel int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from Panic:", err)
		}
	}()
	time.Sleep(60 * time.Second)
	if gameLevel == asyncGameLevel {
		appendLabel(middleGameText[gameLevel])
		middleImage := createImage(middleGame[gameLevel])
		imageContent.appendImage(middleImage)
	}
}

/**
Function to be run [async] the game over of the level.
We use [recover] operator to handle possible panic runtime errors.
*/
func renderGameOver(asyncGameLevel int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from Panic:", err)
		}
	}()
	time.Sleep(120 * time.Second)
	if gameLevel == asyncGameLevel {
		appendLabel(gameOverText[gameLevel])
		gameOverImage := createImage(gameOver[gameLevel])
		imageContent.appendImage(gameOverImage)
		restartGame()
	}
}

/**
Clear all content and invoke again the main method.
*/
func restartGame() {
	time.Sleep(5 * time.Second)
	imageContent.container.RemoveAll()
	titleContent.container.RemoveAll()
	labelContent.RemoveAll()
	inputContent.RemoveAll()
	boardContent.RemoveAll()
	gameLevel = 0
	main()
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
Util function to check if an action is part of the LevelAction
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
