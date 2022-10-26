package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

var deck []Card
var hand []Card

var gameState string
var quit bool

var vm window.SfVideoMode
var cs window.SfContextSettings
var w graphics.Struct_SS_sfRenderWindow

var font graphics.Struct_SS_sfFont

var vmOptions VMOptions

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Starting...")
	font = graphics.SfFont_createFromFile("ttf/Russo.ttf")
	cs = window.NewSfContextSettings()
	if font == nil || font.Swigcptr() == 0 {
		panic("couldn't load font")
	}

	quit = false
	gameState = "game"
	deck = makeShuffledDeck()
	vmOptions = unmarhsalVMOptions()
	setupVM(&vm)
	setupWindow()
	defer graphics.SfFont_destroy(font)
	defer window.SfWindow_destroy(w)
	defer window.DeleteSfContextSettings(cs)
	for !quit {
		switch gameState {
		case "game":
			gameLoop()
			quit = true
		}
	}
	fmt.Println("Finished")
}
