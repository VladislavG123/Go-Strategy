package main

type IMovementBehavior interface {
	Move()
}

type IFireBehavior interface {
	Fire()
}

type ITank interface {
	SetMovementBehavior(behavior IMovementBehavior)
	SetFireBehavior(behavior IFireBehavior)

	IMovementBehavior
	IFireBehavior
}

// Base Tank

type BaseTank struct {
	fireBehavior     IFireBehavior
	movementBehavior IMovementBehavior
}

func (tank *BaseTank) Move() {
	tank.movementBehavior.Move()
}

func (tank *BaseTank) Fire() {
	tank.fireBehavior.Fire()
}

func (tank *BaseTank) SetMovementBehavior(behavior IMovementBehavior) {
	tank.movementBehavior = behavior
}

func (tank *BaseTank) SetFireBehavior(behavior IFireBehavior) {
	tank.fireBehavior = behavior
}
