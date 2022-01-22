package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"testing"
)

// Vérification que les particules générées quand config.General.RandomSpawn
// vaut true sont bien toutes à l'intérieur de l'écran
func TestNewParticleInScreen(t *testing.T) {
	config.General.RandomSpawn = true
	s := NewSystem()
	for i := 0; i < 100; i++ {
		config.General.WindowSizeX = rand.Intn(600) + 200
		config.General.WindowSizeY = rand.Intn(600) + 400
		config.General.SpawnX = float64(config.General.WindowSizeX) / 2
		config.General.SpawnY = float64(config.General.WindowSizeY) / 2
		config.General.SpawnXRange = float64(config.General.WindowSizeX) / 2
		config.General.SpawnYRange = float64(config.General.WindowSizeY) / 2
		p := newParticle(&s)
		if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX) ||
			p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY) {
			t.Fail()
		}
	}
}

// Vérification que les particules générées quand config.General.RandomSpawn
// vaut false sont bien toutes à la position demandée, c'est-à-dire aux
// coordonnées (config.General.SpawnX, config.General.SpawnY)
func TestNewParticleAtSpawnPoint(t *testing.T) {
	config.General.RandomSpawn = false
	s := NewSystem()
	p := newParticle(&s)
	if math.Round(p.PositionX) != config.General.SpawnX ||
		math.Round(p.PositionY) != config.General.SpawnY {
		t.Fail()
	}
}

// Vérification que la fonction update met bien à jour la position des particules
// sur lesquelles on l'utilise
func TestParticleUpdate(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)
		goalX := p.PositionX + p.SpeedX
		goalY := p.PositionY + p.SpeedY
		p.update()
		if p.PositionX != goalX || p.PositionY != goalY {
			t.Fail()
		}
	}
}

// Vérification que la SpeedX et SpeedY des particules est bien comprise dans l'intervalle défini
func TestParticleSpeed(t *testing.T) {
	config.General.MinXSpeed = -50
	config.General.MaxXSpeed = 50
	config.General.MinYSpeed = -50
	config.General.MaxYSpeed =  50
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)
		if p.SpeedX < config.General.MinXSpeed || p.SpeedX > config.General.MaxXSpeed ||
		p.SpeedY < config.General.MinYSpeed || p.SpeedY > config.General.MaxYSpeed {
			t.Fail()
		}
	}
}

// Vérification que multicolore crée des couleurs RGB comprises entre 0 et 1
func TestParticleMulticolore(t *testing.T) {
	config.General.Multicolore = true
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)
		if p.ColorRed < 0 || p.ColorRed > 1 ||
		p.ColorGreen < 0 || p.ColorGreen > 1 ||
		p.ColorBlue < 0 || p.ColorBlue > 1 {
			t.Fail()
		}
	}
}

// Vérification que les marges tue les particules
func TestParticleMarges(t *testing.T) {
	config.General.Gravite = false
	config.General.RenderMarges = true
	config.General.Marges = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.SpawnX = 400
	config.General.SpawnY = 300
	config.General.SpawnXRange = 400
	config.General.SpawnYRange = 300
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		for p.PositionX > 0 && p.PositionX < float64(config.General.WindowSizeX) && 
		p.PositionY > 0 && p.PositionY < float64(config.General.WindowSizeY) {
			p.update()
		}

		if !p.Dead {
			t.Fail()
		}
	}
}

// Vérification que les valeurs données par LifeSpanDurationRange sont bien
// comprises dans l'intervalle donné
func TestLifeSpanDurationRange(t *testing.T) {
	config.General.RenderMarges = false
	config.General.LifeSpan = true
	config.General.LifeSpanDuration = 50
	config.General.LifeSpanDurationRange = 50
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		if p.LifeSpan < config.General.LifeSpanDuration - config.General.LifeSpanDurationRange ||
		p.LifeSpan > config.General.LifeSpanDuration + config.General.LifeSpanDurationRange {
			t.Fail()
		}
	}
}

// Vérification que le lifespan tue bien les particules après x appels
// à la fonction update
func TestParticleLifeSpan(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		var LifeSpan int = p.LifeSpan
		for j := 0; j <= LifeSpan; j++ {
			p.update()
		}

		if p.LifeSpan > 0 || !p.Dead {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseOpacity diminue l'opacité d'une particule
func TestDecreaseOpactiy(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.Opacity = rand.Float64()
		var Opacity = p.Opacity
		p.decreaseOpacity(rand.Float64())

		if p.Opacity > Opacity {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseScale diminue la taille d'une particule
func TestDecreaseScale(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.ScaleX, p.ScaleY = float64(rand.Intn(100)), float64(rand.Intn(100))
		var ScaleX, ScaleY = p.ScaleX, p.ScaleY
		p.decreaseScale(float64(rand.Intn(100)))

		if p.ScaleX > ScaleX || p.ScaleY > ScaleY {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseRed diminue la couleur rouge d'une particule
func TestDecreaseRed(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.ColorRed = rand.Float64()
		var ColorRed = p.ColorRed
		p.decreaseRed(rand.Float64())

		if p.ColorRed > ColorRed {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseGreen diminue la couleur verte d'une particule
func TestDecreaseGreen(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.ColorGreen = rand.Float64()
		var ColorGreen = p.ColorGreen
		p.decreaseGreen(rand.Float64())

		if p.ColorGreen > ColorGreen {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseBlue diminue la couleur bleue d'une particule
func TestDecreaseBlue(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.ColorBlue = rand.Float64()
		var ColorBlue = p.ColorBlue
		p.decreaseBlue(rand.Float64())

		if p.ColorBlue > ColorBlue {
			t.Fail()
		}
	}
}

// Vérifictation que rotateParticle modifie la rotation d'une particule
func TestRotateParticle(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.Rotation = rand.Float64()
		var Rotation = p.Rotation
		p.rotateParticle(rand.Float64())

		if p.Rotation < Rotation {
			t.Fail()
		}
	}
}

// Vérifictation que stopCorailParticle tue une particule et lui enlève sa rotation et sa vitesse
func TestStopCorailParticle(t *testing.T) {
	s := NewSystem()
	p := newParticle(&s)
	p.stopCorailParticle()
	if p.Rotation != 0 || p.SpeedX != 0 || p.SpeedY != 0 || !p.Dead {
		t.Fail()
	}
}

// Vérifictation que graity ajoute une valeur gravité à la vitesse Y d'une particule
func TestGravity(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		p := newParticle(&s)

		p.SpeedY = rand.Float64() * 100
		var SpeedY = p.SpeedY
		for j := 0; j < 10; j++ {
			p.gravity(-rand.Float64())
		}

		if p.SpeedY > SpeedY {
			t.Fail()
		}
	}
}