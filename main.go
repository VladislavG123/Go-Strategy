package main

func main() {
	division := []ITank{NewMainBattleTank(), NewInfantryArmoredVehicle(), NewArmoredPersonnelCarrier(), NewUaz()}

	for _, tank := range division[:] {
		tank.move()
		tank.fire()
	}
}
