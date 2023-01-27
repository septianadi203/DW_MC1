package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(character string) Profile {

	if character == "Sasuke" {
		return Profile{
			Name:   character,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	}
	if character == "Goku" {
		return Profile{

			Name:   character,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	}
	if character == "Naruto" {
		return Profile{

			Name:   character,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}

	return Profile{}

}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
	return profile
}

func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name: ", profile.Name)
	fmt.Println("Health: ", profile.Health)
	fmt.Println("Power: ", profile.Power)
	fmt.Println("Exp: ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health: ", powerUp.Health)
	fmt.Println("Power: ", powerUp.Power)
	fmt.Println("Exp: ", powerUp.Exp)
}
