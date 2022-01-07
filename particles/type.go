package particles

import (
	"project-particles/config"
	"math/rand"
)

// System définit un système de particules.
// Pour le moment il ne contient qu'un tableau de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content []Particle
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	SpeedX, SpeedY 					float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
}


// Add crée une particule.
// Elle définit la particule avec ses différentes caractéristiques.
// # Sortie : la particule créée
func (particle Particle) Add() Particle {

	var PositionX, PositionY float64

	// Définir la position en X et Y en fonction de random spawn
	if config.General.RandomSpawn {
		PositionX, PositionY = float64(rand.Intn(config.General.WindowSizeX)), float64(rand.Intn(config.General.WindowSizeY))
	} else {
		PositionX, PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
	}

	particle = Particle {
		PositionX: PositionX,
		PositionY: PositionY,
		ScaleX: 1, ScaleY: 1,
		SpeedX : config.General.MinXSpeed + rand.Float64() * (config.General.MaxXSpeed - config.General.MinXSpeed),
		SpeedY : config.General.MinYSpeed + rand.Float64() * (config.General.MaxYSpeed - config.General.MinYSpeed),
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
	}

	return particle
}