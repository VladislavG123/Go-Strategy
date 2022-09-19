package main

type IMovementBehavior interface {
	move()
}

type IFireBehavior interface {
	fire()
}

type ITank interface {
	move()
	fire()

	setMovementBehavior(behavior IMovementBehavior)
	setFireBehavior(behavior IFireBehavior)
}

// Base Tank

type BaseTank struct {
	fireBehavior     IFireBehavior
	movementBehavior IMovementBehavior
}

func (tank *BaseTank) move() {
	tank.movementBehavior.move()
}

func (tank *BaseTank) fire() {
	tank.fireBehavior.fire()
}

func (tank *BaseTank) setMovementBehavior(behavior IMovementBehavior) {
	tank.movementBehavior = behavior
}

func (tank *BaseTank) setFireBehavior(behavior IFireBehavior) {
	tank.fireBehavior = behavior
}
