package dndcharacter

import (
	"math/rand"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	value := score - 10
	if value < 0 {
		return (value - 1) / 2
	}
	return value / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := make([]int, 4)
	for i := range rolls {
		rolls[i] = rand.Intn(6) + 1
	}
	for i := 0; i < len(rolls); i++ {
		for j := i + 1; j < len(rolls); j++ {
			if rolls[i] > rolls[j] {
				rolls[i], rolls[j] = rolls[j], rolls[i]
			}

		}
	}
	return rolls[1] + rolls[2] + rolls[3]
}

func calculateHitPoints(Constitution int) int {
	modifier := Modifier(Constitution)
	return modifier + 10
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	rand.Seed(time.Now().UnixNano())
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    Ability(),
	}
	c.Hitpoints = calculateHitPoints(c.Constitution)
	return c
}
