package main

import "math"

type Rectangle struct {
	x, y, w, h float64
}

func (r Rectangle) IsOverlapping(r2 Rectangle) bool {
	if r.x > r2.x+r2.w || r.x+r.w < r2.x {
		return false
	}
	if r.y > r2.y+r2.h || r.y+r.h < r2.y {
		return false
	}
	return true
}

func (r Rectangle) IsOverlappingWithScreen() bool {
	return r.IsOverlapping(Rectangle{0, 0, screenW, screenH})
}

type CollisionType int

const (
	NotColliding CollisionType = iota
	CollidingFromLeft
	CollidingFromRight
	CollidingFromTop
	CollidingFromBottom
)

func (r Rectangle) CollisionCase(r2 Rectangle) CollisionType {
	if !r.IsOverlapping(r2) {
		return NotColliding
	}

	leftMargin := r2.x + r2.w - r.x
	rightMargin := r.x + r.w - r2.x
	topMargin := r2.y + r2.h - r.y
	bottomMargin := r.y + r.h - r2.y

	smallestMargin := math.MaxFloat64
	chosenCase := NotColliding
	if leftMargin > 0. && leftMargin < smallestMargin {
		smallestMargin = leftMargin
		chosenCase = CollidingFromLeft
	}
	if rightMargin > 0. && rightMargin < smallestMargin {
		smallestMargin = rightMargin
		chosenCase = CollidingFromRight
	}
	if bottomMargin > 0. && bottomMargin < smallestMargin {
		smallestMargin = bottomMargin
		chosenCase = CollidingFromBottom
	}
	if topMargin > 0. && topMargin < smallestMargin {
		smallestMargin = topMargin
		chosenCase = CollidingFromTop
	}

	return chosenCase
}
