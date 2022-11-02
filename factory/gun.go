package main

import (
	"errors"
	"fmt"
)

type GunInterface interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() GunInterface {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 GUn",
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() GunInterface {
	return &Musket{
		Gun: Gun{
			name:  "Musket Gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (GunInterface, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, errors.New("wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetail(ak47)
	printDetail(musket)
}

func printDetail(gun GunInterface) {
	fmt.Printf("Gun : %s", gun.getName())
	fmt.Println()
	fmt.Printf("Power : %d", gun.getPower())
	fmt.Println()
}
