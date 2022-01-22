package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"testing"
)

// Vérification que NewSystem crée bien un système avec le nombre de particules
// initiales indiqué dans config.General.InitNumParticles
func TestNewSystemNumParticles(t *testing.T) {
	for i := 0; i < 100; i++ {
		config.General.InitNumParticles = rand.Intn(1000)

		if len(NewSystem().Content) != config.General.InitNumParticles {
			t.Fail()
		}
	}
	config.General.InitNumParticles = 0
	if len(NewSystem().Content) != 0 {
		t.Fail()
	}
}

// Vérification que Update ajoute bien le nombre de particules attendues par
// config.General.SpawnRate
func TestUpdateNumParticles(t *testing.T) {
	s := NewSystem()
	for i := 0; i < 100; i++ {
		config.General.SpawnRate = float64(rand.Intn(2)) + rand.Float64()
		numUpdate := rand.Intn(10)
		toSpawn := int(math.Floor(float64(numUpdate) * config.General.SpawnRate))
		initNum := len(s.Content)
		for j := 0; j < numUpdate; j++ {
			s.Update()
		}
		spawned := len(s.Content) - initNum
		if spawned != toSpawn {
			t.Fail()
		}
		s.NumToSpawn = 0
	}
}

// Vérification que les particules ne sortent pas de l'écran avec CollisionsBords
func TestCollisionsBords(t *testing.T) {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.LifeSpan = false
	config.General.RenderMarges = true
	config.General.Marges = 50
	config.General.CollisionsBords = true
	config.General.MinXSpeed = -10
	config.General.MaxXSpeed = 10
	config.General.MinYSpeed = -10
	config.General.MaxYSpeed = 10
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 100
	s := NewSystem()
	s.ImageWidth, s.ImageHeight = 10, 10
	for i := 0; i < 100; i++ {

		for j := 0; j < 100; j++ {
			s.Update()
		}

		if s.Content[i].Dead {
			t.Fail()
		}
	}
}

// Vérification que la couleur du DVD change bien
func TestChangeDVDColor(t *testing.T) {
	s := NewSystem()
	p := newParticle(&s)

	for i := 0; i < 7; i++ {
		var DVDColorsIndex = s.DVDColorsIndex
		s.changeDVDColor(&p)

		if math.Round(p.ColorRed * 255) != dvdColors[DVDColorsIndex][0] ||
		math.Round(p.ColorGreen * 255) != dvdColors[DVDColorsIndex][1] ||
		math.Round(p.ColorBlue * 255) != dvdColors[DVDColorsIndex][2] {
			t.Fail()
		}
	}
}

// Vérification que snakeDeath arrête bien toutes les particules
func TestSnakeDeath(t *testing.T) {
	config.General.Multicolore = false
	config.General.Snake = true
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 1
	s := NewSystem()
	for i := 0; i < 100; i++ {
		s.Update()
	}
	s.snakeDeath()
	
	if config.General.Snake || config.General.SpawnRate != 0 ||
	s.SnakeXDirection != 0 || s.SnakeYDirection != 0 {
		t.Fail()
	}
	for i := 0; i < len(s.Content); i++ {
		if s.Content[i].ColorRed != 0.59 || s.Content[i].ColorGreen != 0.04 || s.Content[i].ColorBlue != 0 {
			t.Fail()
		}
	}
}