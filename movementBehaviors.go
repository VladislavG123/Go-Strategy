package main

import "fmt"

// WheelsMovement

type WheelsMovement struct {
}

func NewWheelsMovement() *WheelsMovement {
	return &WheelsMovement{}
}

func (m *WheelsMovement) Move() {
	fmt.Println("Moving be wheels")
}

// TracksMovement

type TracksMovement struct {
}

func NewTracksMovement() *TracksMovement {
	return &TracksMovement{}
}

func (m *TracksMovement) Move() {
	fmt.Println("Moving be tracks")
}

// No Movement

type NoMovement struct {
}

func NewNoMovement() *NoMovement {
	return &NoMovement{}
}

func (m *NoMovement) Move() {
	fmt.Println("Can't Move")
}
