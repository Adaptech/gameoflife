package game

import (
	"testing"

	"github.com/Adaptech/gameoflife/commands"
)

func TestPlayWithNoGridShouldFail(t *testing.T) {
	var sut = Game{}
	_, err := sut.Execute(&commands.Play{GameId: "game-1", Grid: ""})
	if err == nil {
		t.Error("expected error")
	}
}

func TestPlayWithRSizeGridShouldSucceed(t *testing.T) {
	var sut = Game{}
	// must be 8x8=64 grid.
	grid := "dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
	events, _ := sut.Execute(&commands.Play{GameId: "game-1", Grid: grid})
	if len(events) != 1 {
		t.Error("expected an event")
	}
}
func TestPlayWithWrongGridSizeShouldFail(t *testing.T) {
	var sut = Game{}
	// must be 8x8=64 grid.
	grid := "ddddddddddddddddddddddddddddddddddd"
	_, err := sut.Execute(&commands.Play{GameId: "game-1", Grid: grid})
	if err == nil {
		t.Error("expected an error")
	}
}
