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
	"./assets/pod.jpeg",
	"./assets/cantina.jpeg",
	"./assets/han_solo.jpeg",
}

//Collection of texts of the whole game
var gameTexts = []string{
	"You wake up confuse, you're not entirely sure what happens.\n " +
		"You have a very bad headache and you have some blood in your head." +
		"You still not sure how, but you end up in the front door of this house in Tatooine.\n" +
		" Maybe you can find more info inside the house....",
	"Once you're inside your see several terminals. One is configure with a Pod racer. It would be great get out of here as soon as possible.\n" +
		"Probably you need to interact with the system to allow you to use the pod",
	"You escape the zone in the pod. Then you notice that two storm troopers start shooting to you. You dont have any weapons.\n" +
		" But you know the pod can go faster....",
	"You manage to scape, then you get inside the cantina. You need to find a pilot to help you out to escape. \n" +
		"Then you see a strange couple of a man and beast sit in the corner of the cantina.\n" +
		"Maybe you should approach....",
	"Hi, I'm Han Solo, best pilot of the Galaxy. Where do you need to go?. The price wont be cheap by the way.\n" +
		"Then you remember the old Kenobi told you about a remote planet where you can find Yoda. The Jedi master",
}

type LevelAction struct {
	minActions int
	actions    []string
}

//Collection of actions of the whole game
var imagesActions = []LevelAction{
	{2, []string{"into", "inside", "house", "cave"}},
	{2, []string{"start", "on", "activate", "pod", "racer"}},
	{1, []string{"accelerate", "full-throttler"}},
	{1, []string{"approach", "go", "walk", "shit"}},
	{1, []string{"Dagobah", "dagobah"}},
}
