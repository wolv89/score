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

	GAME_LEFT  = LiveSet(1 << 20)
	GAME_RIGHT = LiveSet(1 << 16)
)

func (ls LiveSet) PointLeft() LiveSet {

	ls += POINT_LEFT
	return ls

}

func (ls LiveSet) PointRight() LiveSet {

	ls += POINT_RIGHT
	return ls

}

func (ls LiveSet) GameLeft() LiveSet {

	ls += GAME_LEFT
	return ls

}

func (ls LiveSet) GameRight() LiveSet {

	ls += GAME_RIGHT
	return ls

}

func (ls LiveSet) Render() string {

	metaData := uint8(ls >> 24)
	setScore := uint8(ls >> 16)

	scoreLeft := uint8(ls >> 8)
	scoreRight := uint8(ls)

	var b strings.Builder

	fmt.Fprintf(&b, "%08b %08b %08b %08b\n", metaData, setScore, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "% 8d % 8d % 8d % 8d\n", metaData, setScore, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "Meta     Games    Score L  Score R\n\n")

	return b.String()

}
