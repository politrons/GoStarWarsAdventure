package main

import (
	_ "golang.org/x/image/font"
	_ "golang.org/x/image/font/basicfont"
	_ "golang.org/x/image/math/fixed"
)

//Collection of image of the whole game
var images = []string{
	"./assets/tatooin.jpeg",
	"./assets/tatooin_inside.jpeg",
	"./assets/vaina.jpeg",
}

//Collection of texts of the whole game
var gameTexts = []string{
	"You wake up confuse, you're not entirely sure what happens. You have a very bad headache and you have some blood in your mead." +
		"You still not sure how, but you end up in the front door of this house in Tatooine.",
	"bbbbbbbbbbbbbbbb",
	"ccccccccccccccc",
}

//Collection of actions of the whole game
var imagesActions = [][]string{
	{"inside"},
	{"activate"},
	{"drive"},
}
