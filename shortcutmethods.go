package main

import (
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

func getNullIntRect() graphics.SfIntRect {
	return (graphics.SfIntRect)(graphics.SwigcptrSfIntRect(0))
}

func setupVM(vm *window.SfVideoMode) {
	*vm = window.NewSfVideoMode()
	(*vm).SetWidth(vmOptions.Width)
	(*vm).SetHeight(vmOptions.Height)
	(*vm).SetBitsPerPixel(32)
}

func makeVector2(x float32, y float32) graphics.SfVector2f {
	v := graphics.NewSfVector2f()
	v.SetX(x)
	v.SetY(y)
	return v
}

func makeShuffledDeck() []Card {
	ret := getDeck()
	ret = shuffleDeck(ret)
	return ret
}

func getNullRenderState() graphics.SfRenderStates {
	return (graphics.SfRenderStates)(graphics.SwigcptrSfRenderStates(0))
}