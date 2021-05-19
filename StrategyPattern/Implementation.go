package main

import "fmt"

type Weapon interface {
	doAttack()
}

type GameCharacter struct {
	weapon Weapon
}

func (c *GameCharacter) attack() {
	c.weapon.doAttack()
}

func (c *GameCharacter) setWeapon(newWeapon Weapon) {
	c.weapon = newWeapon
}

type Knife struct {
}

func (knife *Knife) doAttack() {
	fmt.Println("Attack with Knife")
}

type Gun struct {
}

func (gun *Gun) doAttack() {
	fmt.Println("Attack with Gun")
}
