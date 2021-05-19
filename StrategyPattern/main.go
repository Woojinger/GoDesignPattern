package main

func main() {
	character := new(GameCharacter)
	knife := new(Knife)
	gun := new(Gun)
	character.setWeapon(knife)
	character.attack()
	// "Attack with Knife"
	character.setWeapon(gun)
	character.attack()
	// "Attack with Gun"
}
