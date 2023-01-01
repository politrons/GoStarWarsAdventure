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
	"You wake up confuse, you're not entirely sure what happens. You have a very bad headache and you have some blood in your head." +
		"You still not sure how, but you end up in the front door of this house in Tatooine." +
		" Maybe uou can find more info inside the house....",
	"Once you're inside your see several terminals. One is configure with a Pod racer. It would be great get out of here as soon as possible.",
	"You escape the zone in the pod. Then you notice that two storm troopers start shooting to you. You dont have any weapons. But you know the pod can go faster",
}

//Collection of actions of the whole game
var imagesActions = [][]string{
	{"inside", "house"},
	{"activate", "po"},
	{"accelerate"},
}
