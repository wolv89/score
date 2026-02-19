package main

import (
	"fmt"
	"strings"
)

// | Meta  | |  Set  | |  Point Score    |
// 0000-0000 0000-0000 0000-0000 0000-0000

type LiveSet uint32

const (
	POINT_LEFT  = LiveSet(1 << 8)
	POINT_RIGHT = LiveSet(1)
	POINT_RESET = 0xFFFF0000

	GAME_LEFT  = LiveSet(1 << 20)
	GAME_RIGHT = LiveSet(1 << 16)
)

var Scores = [5]string{
	" 0",
	"15",
	"30",
	"40",
	"Ad",
}

func (ls LiveSet) PointLeft() LiveSet {

	ls += POINT_LEFT
	return ls

}

func (ls LiveSet) PointRight() LiveSet {

	ls += POINT_RIGHT
	return ls

}

func (ls LiveSet) GameLeft() LiveSet {
	return (ls & POINT_RESET) + GAME_LEFT
}

func (ls LiveSet) GameRight() LiveSet {
	return (ls & POINT_RESET) + GAME_RIGHT
}

func (ls LiveSet) Score() string {

	scoreLeft := uint8(ls >> 8)
	scoreRight := uint8(ls)

	var b strings.Builder
	b.Grow(5)

	if scoreLeft >= 4 {
		b.WriteString("Ad-  ")
	} else if scoreRight >= 4 {
		b.WriteString("  -Ad")
	} else {
		fmt.Fprintf(&b, "%s-%s", Scores[scoreLeft], Scores[scoreRight])
	}

	return b.String()

}

func (ls LiveSet) Render() string {

	metaData := uint8(ls >> 24)
	setScore := uint8(ls >> 16)

	scoreLeft := uint8(ls >> 8)
	scoreRight := uint8(ls)

	var b strings.Builder

	fmt.Fprintf(&b, "%08b %08b %08b %08b\n", metaData, setScore, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "% 8d  % 3d-% 3d % 8d % 8d\n", metaData, setScore>>4, setScore&0xF, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "Meta     Games    Score L  Score R\n\n")

	fmt.Fprintln(&b, ls.Score())

	return b.String()

}
