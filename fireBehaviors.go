package main

import "fmt"

// CannonArmament

type CannonArmament struct {
}

func NewCannonArmament() *CannonArmament {
	return &CannonArmament{}
}

func (a *CannonArmament) fire() {
	fmt.Println("Boom!")
}

// MachineGun Armament

type MachineGunArmament struct {
}

func NewMachineGunArmament() *MachineGunArmament {
	return &MachineGunArmament{}
}

func (m *MachineGunArmament) fire() {
	fmt.Println("Trrrrrrrr")
}

// No Armament

type NoArmament struct {
}

func NewNoArmament() *NoArmament {
	return &NoArmament{}
}

func (m *NoArmament) fire() {
	fmt.Println("Can't fire")
}
