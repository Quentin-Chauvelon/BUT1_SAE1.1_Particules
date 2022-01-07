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

	rand.Seed(time.Now().UnixNano())

	var ParticlesTable []Particle
	
	for i := 0; i < config.General.InitNumParticles; i++ {
		var particle Particle
		ParticlesTable = append(ParticlesTable, particle.Add())
	}

	return System{Content: ParticlesTable}
}