package main

import (
	"fmt"
	"strings"
)

// | Meta  | |  Set  | |  Point Score    |
// 0000-0000 0000-0000 0000-0000 0000-0000

type LiveSet uint32
type WonSet uint8

const (
	// Hard coded, given 4-bit encoding, and the need for a winning tiebreak to be 1 above the target
	// Could allow this limit to be 14, but 12 seems reasonable
	GAMES_LIMIT   = 12
	GAMES_DEFAULT = 6
	GAMES_MINIMUM = 1

	POINT_LEFT  = LiveSet(1 << 8)
	POINT_RIGHT = LiveSet(1)
	POINT_RESET = 0xFFFF0000

	GAME_LEFT  = LiveSet(1 << 20)
	GAME_RIGHT = LiveSet(1 << 16)

	DUECE = uint8(3)

	META_DONE               = LiveSet(1 << 31)
	META_GAMES_TARGET_SHIFT = 24
)

var Scores = [5]string{
	" 0",
	"15",
	"30",
	"40",
	"Ad",
}

func NewLiveSet(gamesTarget uint) LiveSet {

	if gamesTarget < GAMES_MINIMUM || gamesTarget > GAMES_LIMIT {
		gamesTarget = GAMES_DEFAULT
	}

	ls := LiveSet(gamesTarget << META_GAMES_TARGET_SHIFT)

	return ls

}

func (ls LiveSet) PointLeft() LiveSet {

	scoreLeft, scoreRight := ls.PointScores()

	if scoreLeft == DUECE {
		if scoreRight < DUECE {
			// 40-15, 40-30, etc
			return ls.GameLeft()
		} else if scoreRight > DUECE {
			// Right was in advantage, return to duece
			ls -= POINT_RIGHT
			return ls
		}
		// Otherwise fallthrough to simple advantage
	} else if scoreLeft > DUECE {
		// From advantage
		return ls.GameLeft()
	}

	// Base case, simple point won
	ls += POINT_LEFT
	return ls

}

func (ls LiveSet) PointRight() LiveSet {

	scoreLeft, scoreRight := ls.PointScores()

	if scoreRight == DUECE {
		if scoreLeft < DUECE {
			// 15-40, 30-40, etc
			return ls.GameRight()
		} else if scoreLeft > DUECE {
			// Left was in advantage, return to duece
			ls -= POINT_LEFT
			return ls
		}
		// Otherwise fallthrough to simple advantage
	} else if scoreRight > DUECE {
		// From advantage
		return ls.GameRight()
	}

	// Base case, simple point won
	ls += POINT_RIGHT
	return ls

}

func (ls LiveSet) GameLeft() LiveSet {
	return (ls & POINT_RESET) + GAME_LEFT
}

func (ls LiveSet) GameRight() LiveSet {
	return (ls & POINT_RESET) + GAME_RIGHT
}

func (ls LiveSet) PointScores() (uint8, uint8) {
	return uint8(ls >> 8), uint8(ls)
}

func (ls LiveSet) Score() string {

	scoreLeft, scoreRight := ls.PointScores()

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

	scoreLeft, scoreRight := ls.PointScores()

	var b strings.Builder

	fmt.Fprintf(&b, "%08b %08b %08b %08b\n", metaData, setScore, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "% 8d  % 3d-% 3d % 8d % 8d\n", metaData, setScore>>4, setScore&0xF, scoreLeft, scoreRight)
	fmt.Fprintf(&b, "Meta     Games    Score L  Score R\n\n")

	fmt.Fprintln(&b, ls.Score())

	return b.String()

}
