package view

import (
	"github.com/veandco/go-sdl2/sdl"
)

type View interface {
	GetName() string

	HandleSDLEvent(*sdl.Renderer, sdl.Event) bool

	PollEvent() Event

	Draw(*sdl.Renderer)
}

type Event interface {
	GetType() EventType
}

type EventType uint8

const (
	CHANGE_VIEW EventType = iota
)

type ChangeViewEvent struct {
	ToChangeView View //fixme: 変数名何とかしろ
}

func (ev *ChangeViewEvent) GetType() EventType {
	return CHANGE_VIEW
}
