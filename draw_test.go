package main

import (
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

// Vérification que resetSystem réinitialise bien le système
func TestResetSystem (t *testing.T) {
	config.General.InitNumParticles = 100
	config.General.ParticleImage = "assets/particle.png"
	g := game{system: particles.NewSystem()}
	config.General.InitNumParticles = 0
	g.resetSystem()
	if len(g.system.Content) != 0 {
		t.Fail()
	}
}

// Vérification que extensionChanged change bien le nom de l'extension
func TestExtensionChanged (t *testing.T) {
	g := game{system: particles.NewSystem()}
	g.extensionChanged("Test")
	if g.system.ActiveExtension[0] != 84 || g.system.ActiveExtension[1] != 101 ||
	g.system.ActiveExtension[2] != 115 || g.system.ActiveExtension[3] != 116 {
		t.Fail()
	}
}

// Vérification que setConfig modifie bien la configuration
func TestSetConfig (t *testing.T) {
	setConfig("assets/particle.png", 10, true, 400, 300, 400, 300, 1, -10, 10, -10, 10, 1, 1, 1, true, true, 1, true, 100, true, 60, 10, true, true, true, true, true, true, true, true, true, true)
	if config.General.InitNumParticles != 10 || !config.General.RandomSpawn ||
	config.General.SpawnX != 400 || config.General.SpawnY != 300 ||
	config.General.SpawnXRange != 400 || config.General.SpawnYRange != 300 ||
	config.General.SpawnRate != 1 || config.General.MinXSpeed != -10 ||
	config.General.MaxXSpeed != 10 || config.General.MinYSpeed != -10 ||
	config.General.MaxYSpeed != 10 || config.General.Red != 1 ||
	config.General.Green != 1 || config.General.Blue != 1 ||
	!config.General.Multicolore || !config.General.Gravite ||
	config.General.ConstanteGravite != 1 || !config.General.RenderMarges ||
	config.General.Marges != 100 || !config.General.LifeSpan ||
	config.General.LifeSpanDuration != 60 || config.General.LifeSpanDurationRange != 10 ||
	!config.General.CollisionsBords || !config.General.CollisionsParticules ||
	!config.General.Fire || !config.General.Rain || !config.General.Rainbow ||
	!config.General.DVD || !config.General.Corail || !config.General.Circle ||
	!config.General.Fireworks || !config.General.Snake {
		t.Fail()
	}
}