package main

// Main Battle Tank

type MainBattleTank struct {
	BaseTank
}

func NewMainBattleTank() *MainBattleTank {
	return &MainBattleTank{
		BaseTank{
			fireBehavior:     NewCannonArmament(),
			movementBehavior: NewTracksMovement(),
		},
	}
}

// Infantry Armored Vehicle

type InfantryArmoredVehicle struct {
	BaseTank
}

func NewInfantryArmoredVehicle() *InfantryArmoredVehicle {
	return &InfantryArmoredVehicle{
		BaseTank{
			fireBehavior:     NewMachineGunArmament(),
			movementBehavior: NewTracksMovement(),
		},
	}
}

// Armored Personnel Carrier

type ArmoredPersonnelCarrier struct {
	BaseTank
}

func NewArmoredPersonnelCarrier() *ArmoredPersonnelCarrier {
	return &ArmoredPersonnelCarrier{
		BaseTank{
			fireBehavior:     NewMachineGunArmament(),
			movementBehavior: NewWheelsMovement(),
		},
	}
}

// Uaz

type Uaz struct {
	BaseTank
}

func NewUaz() *Uaz {
	return &Uaz{
		BaseTank{
			fireBehavior:     NewNoArmament(),
			movementBehavior: NewNoMovement(),
		},
	}
}
