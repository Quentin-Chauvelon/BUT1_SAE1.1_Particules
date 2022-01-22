package particles

import (
	"project-particles/config"
	"math"
	"math/rand"
)

// Particle définit une particule.
type Particle struct {
	PositionX, PositionY            float64 // Position à l'écran
	Rotation                        float64 // Orientation
	ScaleX, ScaleY                  float64 // Taille
	SpeedX, SpeedY		            float64 // Vitesse
	ColorRed, ColorGreen, ColorBlue float64 // Couleur
	Opacity                         float64 // Transparence
	LifeSpan 						int // Duree de vie
	CollisionCooldown				int // Cooldown avant d'avoir une nouvelle collision
	FireworkExplosion				bool // Particule spéciale qui correspond à l'explosion du feu d'artifice
	Dead							bool // Particule morte (on n'agit plus dessus) (peut être tué par LifeSpan ou les marges)
}


// newParticle est une fonction qui initialise une particule et la retourne.
func newParticle(s *System) Particle {

	var ColorRed, ColorGreen, ColorBlue float64 = 1, 1, 1
	var PositionX, PositionY float64
	var ScaleX, ScaleY float64 = 1, 1
	var SpeedX float64 = config.General.MinXSpeed + rand.Float64() * (config.General.MaxXSpeed - config.General.MinXSpeed)
	var SpeedY float64 = config.General.MinYSpeed + rand.Float64() * (config.General.MaxYSpeed - config.General.MinYSpeed)
	var LifeSpan int

	if config.General.Multicolore {
		ColorRed, ColorGreen, ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
	}

	if config.General.LifeSpan {
		LifeSpan = int(float64(config.General.LifeSpanDuration - config.General.LifeSpanDurationRange) + rand.Float64() * float64((config.General.LifeSpanDuration + config.General.LifeSpanDurationRange) - (config.General.LifeSpanDuration - config.General.LifeSpanDurationRange)))
	}

	if config.General.Fire {
		ScaleX, ScaleY = 0.8, 0.8
	}

	if config.General.Rain {
		ScaleX, ScaleY = 0.1, 5
		ColorRed, ColorGreen, ColorBlue = 0.671, 0.804, 0.812
	}

	if config.General.Rainbow {
		ScaleX, ScaleY = 1, 1
		ColorRed, ColorGreen, ColorBlue = config.General.Red, config.General.Green, config.General.Blue
	}

	if config.General.DVD {
		ScaleX, ScaleY = 1, 1
		ColorRed, ColorGreen, ColorBlue = 1, 0.514, 0
	}

	if config.General.Circle {
		// Lors de la création d'un cercle, on place d'abord la particule aléatoirement entre sur l'axe X
		// puis on utilise une formule pour calculer la position correspondante en Y
		// Cependant, en faisant cela, les particules sont beaucoup plus rasssemblés près de l'axe Y que
		// près de l'axe X, ce qui crée un trou aux alentours de l'axe X.
		// Pour pallier à cela, il y a 1 chance sur 2 (rand.Float64() > 0.5) de prendre la position aléatoire en X
		// et 1 chance sur 2 de la prendre en Y, puis en calculant la position correspondante en X ou en Y
		// on obtient un cercle avec des particules parfaitement réparties sur tous les axes.
		if rand.Float64() > 0.5 {
			SpeedX = math.Sqrt(config.General.MaxXSpeed * config.General.MaxXSpeed - SpeedY * SpeedY)

			// 1 chance sur 2 que la particule soit en négatif (pour couvrir tout le cercle)
			if rand.Float64() > 0.5 {
				SpeedX = -SpeedX
			}

		} else {
			SpeedY = math.Sqrt(config.General.MaxYSpeed * config.General.MaxYSpeed - SpeedX * SpeedX)

			// 1 chance sur 2 que la particule soit en négatif (pour couvrir tout le cercle)
			if rand.Float64() > 0.5 {
				SpeedY = -SpeedY
			}
		}
	}

	if config.General.Snake {
		LifeSpan = s.SnakeLifeSpan
		ScaleX, ScaleY = 6, 6
	}

	if config.General.RandomSpawn {
		PositionX = float64(config.General.SpawnX - config.General.SpawnXRange) + rand.Float64() * (float64((config.General.SpawnX + config.General.SpawnXRange) - (config.General.SpawnX - config.General.SpawnXRange) - s.ImageWidth * ScaleX))
		PositionY = float64(config.General.SpawnY - config.General.SpawnYRange) + rand.Float64() * (float64((config.General.SpawnY + config.General.SpawnYRange) - (config.General.SpawnY - config.General.SpawnYRange) - s.ImageHeight * ScaleY))
	} else {
		PositionX, PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
	}


	var p Particle

	// Position
	p.PositionX = PositionX
	p.PositionY = PositionY

	// Orientation
	p.Rotation = 0

	// Taille
	p.ScaleX = ScaleX
	p.ScaleY = ScaleY

	// Vitesse
	p.SpeedX = SpeedX
	p.SpeedY = SpeedY

	// Couleur
	p.ColorRed = ColorRed
	p.ColorBlue = ColorBlue
	p.ColorGreen = ColorGreen

	// Transparence
	p.Opacity = 1

	// Duree de vie
	p.LifeSpan = LifeSpan

	// Cooldown avant d'avoir une nouvelle collision
	p.CollisionCooldown = 0

	// Particule spéciale qui correspond à l'explosion du feu d'artifice
	p.FireworkExplosion = false

	// Particule morte (on n'agit plus dessus) (peut être tué par LifeSpan, les Marges, le Firework ou le Snake)
	p.Dead = false

	return p
}

// fireworkExploding créer une particule avec des certaines propriétés définies (qui dépendent de chaque particule qui va créer un feu d'artifice
// Elle permet aussi de laisser une traînée derrière les particules lors d'un feu d'artifice
func fireworkExploding(Radius float64, PositionX float64, PositionY float64, SpeedX float64,
ColorRed float64, ColorGreen float64, ColorBlue float64, LifeSpan int, Scale float64, signe float64) Particle {

	var p Particle = Particle {
		PositionX: PositionX,
		PositionY: PositionY,
		Rotation: 0.1,
		ScaleX: Scale, ScaleY: Scale,
		SpeedX : SpeedX,
		SpeedY : math.Sqrt(Radius*Radius - SpeedX * SpeedX) * signe,
		ColorRed: ColorRed, ColorGreen: ColorGreen, ColorBlue: ColorBlue,
		Opacity: 1,
		LifeSpan: LifeSpan,
		CollisionCooldown: 0,
		FireworkExplosion: true,
		Dead: false,
	}

	// Traînée immobile
	if SpeedX <= 0.01 && SpeedX >= -0.01 && Scale < 0.4 {
		p.SpeedY = 0
		p.Rotation = 0
	}

	return p
}

// snakeParticule créer les particules rouges qu'il faut toucher lors de l'utilisation de l'extension Snake
func snakeParticule(s *System) Particle {

	var p Particle = Particle {
		PositionX: rand.Float64() * (float64(config.General.WindowSizeX) - s.ImageWidth * 5),
		PositionY: rand.Float64() * (float64(config.General.WindowSizeY) - s.ImageHeight * 5),
		Rotation: 0,
		ScaleX: 5, ScaleY: 5,
		SpeedX : 0,
		SpeedY : 0,
		ColorRed: 1, ColorGreen: 0, ColorBlue: 0,
		Opacity: 1,
		LifeSpan: int(math.Inf(1)),
		CollisionCooldown: 0,
		FireworkExplosion: false,
		Dead: false,
	}

	return p
}

// decreaseOpacity diminue l'opacité d'une particule
func (p *Particle) decreaseOpacity(opacity float64) {
	p.Opacity -= opacity
}

// decreaseScale diminue la taille d'un particule
func (p *Particle) decreaseScale(scale float64) {
	p.ScaleX -= scale
	p.ScaleY -= scale
}

// decreaseRed diminue la couleur rouge d'une particule
func (p *Particle) decreaseRed(red float64) {
	p.ColorRed -= red
}

// decreaseGreen diminue la couleur verte d'une particule
func (p *Particle) decreaseGreen(green float64) {
	p.ColorGreen -= green
}

// decreaseBlue diminue la couleur blueue d'une particule
func (p *Particle) decreaseBlue(blue float64) {
	p.ColorBlue -= blue
}

// rotateParticle modifie la rotation d'une particule
func (p *Particle) rotateParticle(rotation float64) {
	p.Rotation += rotation
}

// stopCorailParticle arrête une particule (uniquement si l'extension Corail est activée)
func (p *Particle) stopCorailParticle() {
	p.SpeedX, p.SpeedY = 0, 0
	p.Rotation = 0
	p.Dead = true
}

// gravity ajoute une valeur gravité à la vitesse Y de la particule pour donner un effet de gravité
func (p *Particle) gravity(gravity float64) {
	p.SpeedY += gravity
}


// Update met à jour l'état d'une particule à chaque fois qu'on l'appelle
func (p *Particle) update() {

	if config.General.Fire {
		p.decreaseOpacity(0.02)
		p.decreaseScale(0.01)

		if p.ColorBlue > 0 {
			p.decreaseBlue(0.083)
		} else if p.ColorGreen > 0 {
			p.decreaseGreen(0.083)
		} else if p.ColorRed > 0.5 {
			p.decreaseRed(0.417)
		}
	}

	if config.General.Rainbow {
		p.rotateParticle(0.05)
	}

	if !p.Dead && config.General.Corail {
		p.ColorRed = (-p.PositionY * (255 - config.General.BottomColorRed) / (float64(config.General.WindowSizeY) - 100) + 255) / 255
		p.ColorGreen = (-p.PositionY * (255 - config.General.BottomColorGreen) / (float64(config.General.WindowSizeY) - 100) + 255) / 255
		p.ColorBlue = (-p.PositionY * (255 - config.General.BottomColorBlue) / (float64(config.General.WindowSizeY) - 100) + 255) / 255
		p.rotateParticle(0.05)
	}


	if config.General.Circle {
		p.decreaseRed(0.00022)
		p.decreaseGreen(0.006)
		p.decreaseBlue(0.025)
	}

	if config.General.Gravite {
		p.gravity(-config.General.ConstanteGravite)
	}

	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY

	if config.General.RenderMarges {
		if p.PositionX < 0 - config.General.Marges ||
		p.PositionY < 0 - config.General.Marges ||
		p.PositionX > float64(config.General.WindowSizeX) + config.General.Marges ||
		p.PositionY > float64(config.General.WindowSizeY) + config.General.Marges {
			p.Dead = true
		}
	}

	p.LifeSpan -= 1

	if config.General.LifeSpan && p.LifeSpan <= 0 {
		p.Dead = true
	}
}