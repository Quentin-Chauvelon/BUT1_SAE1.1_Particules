package particles

import (
	"testing"
	"project-particles/config"
)

// Créer un système avec une valeur de InitNumParticles négative
func TestNewSystemNegatif(t *testing.T) {
	config.General.InitNumParticles = -5
	Content := NewSystem().Content
	if len(Content) != 0 {
		t.Fail()
	}
}

// Créer un système avec une valeur de InitNumParticles 0
func TestNewSystemNul(t *testing.T) {
	config.General.InitNumParticles = 0
	Content := NewSystem().Content

	if len(Content) != 0 {
		t.Fail()
	}
}

// Créer un système avec une valeur de InitNumParticles positive
func TestNewSystemPositif(t *testing.T) {
	config.General.InitNumParticles = 5
	Content := NewSystem().Content

	if len(Content) != config.General.InitNumParticles {
		t.Fail()
	}
}

// Tester le random spawn == false, les particules devraient avoir la position de spawn X et spawn Y
func TestRandomSpawnFalse(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 100

	Content := NewSystem().Content

	for i := 0; i < len(Content); i++ {
		if Content[i].PositionX != float64(config.General.SpawnX) || Content[i].PositionY != float64(config.General.SpawnY) {
			t.Fail()
		}
	}
}

// Tester le random spawn == true et vérifier que les positions sont différentes entre les différentes particules
func TestRandomSpawnTrue(t *testing.T) {
	config.General.RandomSpawn = true
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	Content := NewSystem().Content

	var SamePositionCounter int

	for i := 0; i < len(Content); i++ {
		for y := 0; y < len(Content); y++ {
			if Content[i].PositionX == Content[y].PositionX && Content[i].PositionY == Content[y].PositionY {
				SamePositionCounter++
			}
		}
	}

	// On Divise par 50 et on ajoute 1 pour avoir une marge d'erreur au cas où plusieurs particules auraient exactement la même position
	// Exemple : pour 100 particules, il faut que au au maximum 100 + (100 / 50) + 1 = 103 particules aient la même position (il y aura toujours au moins 100 particules puisque lors de la double boucle : chaque particule est comparé avec elle-même, elle aura donc la même position)
	if SamePositionCounter > config.General.InitNumParticles + config.General.InitNumParticles / 50 + 1 {
		t.Fail()
	}
}

// Vérifier que chaque particule a une speed comprise dans l'intervalle minimum et maximum de position
func TestRandomSpawnTrueIntervalle(t *testing.T) {

	Content := NewSystem().Content

	for i := 0; i < len(Content); i++ {

		if Content[i].PositionX < 0 ||
		Content[i].PositionY < 0 ||
		Content[i].PositionX > float64(config.General.WindowSizeX) ||
		Content[i].PositionY > float64(config.General.WindowSizeY) {
			t.Fail()
		}
	}
}

// Vérifier que les vitesses sont toutes différentes
func TestRandomSpeed(t *testing.T) {
	config.General.MinXSpeed = -10
	config.General.MaxXSpeed = 10
	config.General.MinYSpeed = -10
	config.General.MaxYSpeed = 10

	Content := NewSystem().Content

	var SamePositionCounter int

	for i := 0; i < len(Content); i++ {
		for y := 0; y < len(Content); y++ {
			if Content[i].SpeedX == Content[y].SpeedX && Content[i].SpeedY == Content[y].SpeedY {
				SamePositionCounter++
			}
		}
	}

	// Marge d'erreur dans le cas où plusieurs particules aient la même vitesse (même principe que pour TestRandomSpawnTrue)
	if SamePositionCounter > int(config.General.MaxXSpeed - config.General.MinXSpeed) + int(config.General.MaxXSpeed - config.General.MinXSpeed) / 10 + 1 {
		t.Fail()
	}
}

// Vérifier
func TestRandomSpeedIntervalle(t *testing.T) {

	Content := NewSystem().Content

	for i := 0; i < len(Content); i++ {
		if Content[i].SpeedX < config.General.MinXSpeed ||
		Content[i].SpeedX > config.General.MaxXSpeed ||
		Content[i].SpeedY < config.General.MinYSpeed ||
		Content[i].SpeedY > config.General.MaxYSpeed {
			t.Fail()
		}
	}
}

// Vérifier que chaque particule a une speed comprise dans l'intervalle minimum et maximum de position
func TestSpeed(t *testing.T) {
	config.General.SpawnRate = 0

	System := NewSystem()

	// Copier le tableau pour garder les positions d'origine
	var ContentCopy = make([]Particle, len(System.Content), len(System.Content))
	
	for i := 0; i < len(System.Content); i++ {	
		ContentCopy[i] = System.Content[i]
	}

	for i := 0; i < 10; i++ {
		System.Update()
	}

	for i := 0; i < len(System.Content); i++ {
		if ContentCopy[i].PositionX > 0 && System.Content[i].SpeedX > 0 && System.Content[i].PositionX < 0 {
			t.Fail()
		} else if ContentCopy[i].PositionX < 0 && System.Content[i].SpeedX < 0 && System.Content[i].PositionX > 0 {
			t.Fail()
		} else if ContentCopy[i].PositionY > 0 && System.Content[i].SpeedY < 0 && System.Content[i].PositionY < 0 {
			t.Fail()
		} else if ContentCopy[i].PositionY < 0 && System.Content[i].SpeedY > 0 && System.Content[i].PositionY > 0 {
			t.Fail()
		}
	}
}


// Tester spawn rate avec une valeur négative
func TestSpawnRateNegatif(t *testing.T) {
	config.General.SpawnRate = -10

	System := NewSystem()

	for i := 0; i < 10; i++ {
		System.Update()
	}

	if float64(len(System.Content)) != float64(config.General.InitNumParticles) {
		t.Fail()
	}
}

// Tester spawn rate avec une valeur nulle
func TestSpawnRateNul(t *testing.T) {
	config.General.SpawnRate = 0

	System := NewSystem()

	for i := 0; i < 10; i++ {
		System.Update()
	}

	if float64(len(System.Content)) != float64(config.General.InitNumParticles) {
		t.Fail()
	}
}

// Tester spawn rate avec une valeur positive comprise entre 0 et 1
func TestSpawnRateInférieur1(t *testing.T) {
	config.General.SpawnRate = 0.1

	System := NewSystem()

	for i := 0.0; i < ((1/config.General.SpawnRate) * 2); i++ {
		System.Update()
	}

	if float64(len(System.Content)) != float64(config.General.InitNumParticles) + (((1 / config.General.SpawnRate) * 2)	 * config.General.SpawnRate) {
		t.Fail()
	}
}

// Tester spawn rate avec une valeur positive supérieure à 1
func TestSpawnRateSupérieur1(t *testing.T) {
	config.General.SpawnRate = 10

	System := NewSystem()

	for i := 0; i < 10; i++ {
		System.Update()
	}

	if float64(len(System.Content)) != float64(config.General.InitNumParticles) + config.General.SpawnRate * float64(10) {
		t.Fail()
	}
}
