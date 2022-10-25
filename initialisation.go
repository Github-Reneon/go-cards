package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

func getDeck() []Card {
	if _, err := os.Stat("config/deck.json"); err == nil {
		bytes, _ := ioutil.ReadFile("deck.json")
		var deck []Card
		json.Unmarshal(bytes, &deck)
		return deck
	} else {
		return nil
	}
}

func shuffleDeck(deck []Card) []Card {
	ret := deck
	for i := range ret {
		j := rand.Intn(i + 1)
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func unmarhsalVMOptions() VMOptions {
	path := "config/vmoptions.json"
	var vmOptions VMOptions
	vmOptions.windowName = "Reverse Blackjack"
	if _, err := os.Stat(path); err == nil {
		bytes, _ := ioutil.ReadFile(path)
		json.Unmarshal(bytes, &vmOptions)
		return vmOptions
	} else {
		vmOptions.Width = 700
		vmOptions.Height = 600
		return vmOptions
	}
}

func setupWindow() {
	w = graphics.SfRenderWindow_create(vm, vmOptions.windowName, uint(window.SfResize|window.SfClose), cs)
}