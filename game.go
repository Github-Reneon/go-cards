package main

import (
	"fmt"
	"time"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

func gameLoop() {

	textures := map[string]graphics.Struct_SS_sfTexture{
		"cardsTexture": nil,
	}

	sprites := map[string][]graphics.Struct_SS_sfSprite{
		"cardSprites": nil,
	}

	// Handle textures
	loadTextures(&textures)

	for _, texture := range textures {
		defer graphics.SfTexture_destroy(texture)
	}

	//Handle sprites
	loadSprites(&sprites)

	for _, spriteGroup := range sprites {
		for _, sprite := range spriteGroup {
			defer graphics.SfSprite_destroy(sprite)
		}
	}

	//Setup event manager
	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	//Setup text
	scoreText := graphics.SfText_create()
	cashText := graphics.SfText_create()

	initText(&scoreText)
	initText(&cashText)

	defer graphics.SfText_destroy(scoreText)
	defer graphics.SfText_destroy(cashText)

	/*
		//set up the cards
		backCard := Card{
			Value: 3,
			Suit:  5,
		}
	*/

	x, y := 10.0, 10.0

	spd := 3.0

	down, right := 1.0, 1.0

	timestamp := time.Now()

	frames := 0

	frameRate := []float32{}

	mainMenu := false

	for !quit && !mainMenu {
		frameRate = append(frameRate, float32(time.Since(timestamp).Seconds()))
		deltatime := float32(time.Since(timestamp).Seconds())
		timestamp = time.Now()

		eventHandler(&ev, &spd)

		graphics.SfRenderWindow_clear(w, graphics.GetSfWhite())

		frames++

		average := getAverage(frameRate, frames)

		x += spd * float64(deltatime) * right

		y += spd * float64(deltatime) * down

		if x > float64(vmOptions.Width) || x < 0.0 {
			right *= -1.0
		}

		if y > float64(vmOptions.Height) || y < 0.0 {
			down *= -1.0
		}

		graphics.SfText_setPosition(scoreText, makeVector2(float32(x), float32(y)))

		if frames > 100 {
			frames = 0
			frameRate = []float32{}
		}

		graphics.SfText_setString(scoreText, fmt.Sprint(int(1/average)))

		graphics.SfRenderWindow_drawText(w, scoreText, getNullRenderState())

		graphics.SfRenderWindow_display(w)
	}
}

func eventHandler(ev *window.SfEvent, spd *float64) {
	for window.SfWindow_pollEvent(w, *ev) > 0 {
		if (*ev).GetEvType() == window.SfEventType(window.SfEvtClosed) {
			return
		}

		if (*ev).GetEvType() == window.SfEventType(window.SfEvtKeyReleased) {
			switch (*ev).GetKey().GetCode() {
			case window.SfKeyCode(window.SfKeyQ):
				quit = true
			case window.SfKeyCode(window.SfKeyEqual):
				*spd += 10.0
			case window.SfKeyCode(window.SfKeyHyphen):
				*spd -= 10.0
				if *spd < 0 {
					*spd = 0
				}
			}
		}
	}
}

func getAverage(framerate []float32, frames int) float32 {
	total := float32(0)

	for _, i := range framerate {
		total += i
	}

	return total / float32(frames)
}

func initText(text *graphics.Struct_SS_sfText) {
	graphics.SfText_setFont(*text, font)
	graphics.SfText_setCharacterSize(*text, 12)
	graphics.SfText_setFillColor(*text, graphics.GetSfBlack())
}

func loadSprites(sprites *map[string][]graphics.Struct_SS_sfSprite) {
	(*sprites)["cardsSprites"] = append((*sprites)["cardsSprites"], graphics.SfSprite_create())
}

func loadTextures(textures *map[string]graphics.Struct_SS_sfTexture) {
	cardTexturePath := "img/cards.jpg"
	(*textures)["cardsTexture"] = graphics.SfTexture_createFromFile(cardTexturePath, getNullIntRect())
}

func setScore(text *graphics.Struct_SS_sfText, hand *[]Card) {

	fmt.Println((*hand))

	var score []int
	var ones int

	score = append(score, 0)

	for i := 0; i < len(*hand); i++ {
		value := (*hand)[i].Value
		if value != 1 {
			if value > 10 {
				value = 10
			}
			score[0] += value
		} else {
			ones++
		}
	}

	if ones > 0 {
		firstElvn := false

		if score[0]+11 < 21 && !firstElvn {
			score = append(score, score[0])
			score[0] += 11
			score[1] += 1
		} else {
			score[0] += 1
		}

		if ones > 1 {
			for i := 1; i < ones; i++ {
				score[0] += 1
				if len(score) == 2 {
					score[1] += 1
				}
			}
		}
	}

	scoreText := "Score: "

	scoreText += fmt.Sprint(score[0])

	if len(score) == 2 {
		scoreText += "/"
		scoreText += fmt.Sprint(score[1])
	}

	for i := 0; i < len(score); i++ {
		if score[i] == 21 {
			scoreText += " Blackjack!"
			break
		}
		if score[i] > 21 {
			scoreText += " Bust!"
			break
		}
	}

	graphics.SfText_setString(*text, scoreText)

}
