package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)
// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	var ParticlesTable []Particle

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < config.General.InitNumParticles; i++ {
		ParticlesTable = append(ParticlesTable, 
			Particle{
				PositionX: float64(rand.Intn(config.General.WindowSizeX)),
				PositionY: float64(rand.Intn(config.General.WindowSizeY)),
				ScaleX: 1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
			})
		}

	return System{Content: ParticlesTable}
}

// return System{Content: []Particle{
	// 	Particle{
	// 		PositionX: float64(rand.Intn(config.General.WindowSizeX),
	// 		PositionY: float64(rand.Intn(config.General.WindowSizeY)),
	// 		ScaleX: 1, ScaleY: 1,
	// 		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
	// 		Opacity: 1,
	// 	},
	// }}