package main

import (
	"fmt"

	"github.com/telroshan/go-sfml/v2/graphics"
)

func setCardSprite(sprite *graphics.Struct_SS_sfSprite, texture *graphics.Struct_SS_sfTexture, card Card) {
	graphics.SfSprite_setTexture(*sprite, *texture, 1)

	rect := graphics.NewSfIntRect()

	textureDims := graphics.SfTexture_getSize(*texture)

	valuePosition := int(int(textureDims.GetX()) / 13 * (card.Value - 1))
	suitPosition := int(int(textureDims.GetY()) / 5 * (card.Suit - 1))

	fmt.Println(valuePosition, ":", suitPosition)

	rect.SetLeft(valuePosition)
	rect.SetWidth(int(textureDims.GetX() / 13))
	rect.SetTop(suitPosition)
	rect.SetHeight(int(textureDims.GetY() / 5))

	graphics.SfSprite_setTextureRect(*sprite, rect)
}
