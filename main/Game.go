package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Board struct {
	canvasWidth  float32
	canvasHeight float32
	fps          int
	then         int64
	margin       int
}

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

	now := time.Now().UnixMilli()
	game := &Board{
		800,
		500,
		10,
		now,
		4,
	}

	fpsInterval := int64(1000 / game.fps)

	backgroundImg := canvas.NewImageFromImage(background)
	backgroundImg.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)

	c := container.New(layout.NewMaxLayout(), backgroundImg, playerImg)
	window.SetContent(c)

	go func() {

		for {
			time.Sleep(time.Millisecond)

			now := time.Now().UnixMilli()
			elapsed := now - game.then

			if elapsed > fpsInterval {
				game.then = now
				draw.Draw(sprite, sprite.Bounds(), image.Transparent, image.ZP, draw.Src)
				playerImg = canvas.NewRasterFromImage(sprite)
				c.Refresh()

			}
		}

	}()

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
				backgroundImg = canvas.NewImageFromImage(background2)
				backgroundImg.FillMode = canvas.ImageFillOriginal

				sprite = image.NewRGBA(background.Bounds())

				playerImg = canvas.NewRasterFromImage(sprite)

				c = container.New(layout.NewMaxLayout(), backgroundImg, playerImg)
				window.SetContent(c)
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
