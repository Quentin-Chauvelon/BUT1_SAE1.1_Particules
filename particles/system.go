package particles

import (
	"project-particles/config"
	"project-particles/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
)

// Couleurs pour le DVD
var dvdColors [][]float64 = [][]float64 {
	[]float64{0,38,255},
	[]float64{255,250,1},
	[]float64{255,38,0},
	[]float64{255,0,139},
	[]float64{190,0,255},
	[]float64{0,254,255},
	[]float64{255,131,0},
}

// System définit un système de particules.
type System struct {
	Content    []Particle // Particules actuelles
	NumToSpawn float64    // Nombre de particules restant à ajouter
	DVDColorsIndex int 	  // Index pour changer la couleur du DVD en utilisant le tableau
	ImageWidth float64 	  // Largeur de l'image
	ImageHeight float64   // Hauteur de l'image
	SnakeXDirection float64   // Direction du snake en X
	SnakeYDirection float64   // Direction du snake en Y
	SnakeLifeSpan int // Taille du snake
	SnakeScore int
	ExtensionsButtonPressed bool // Le bouton pour choisir les extensions a été pressé
	PauseGame bool // Appuyer sur espace pour pauser le jeu
	ActiveExtension string // Nom de l'extension active (parmi Feu, Pluie, Vagues multicolores, DVD, Corail, Cercle, Feu d'artifice, Snake et Personnalisé)
}

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
func NewSystem() System {

	var s System

	if assets.ParticleImage != nil {
		ImageWidthTmp, ImageHeightTmp := assets.ParticleImage.Size()
		s.ImageWidth, s.ImageHeight = float64(ImageWidthTmp), float64(ImageHeightTmp)
	}

	s.SnakeXDirection = 0.15
	s.SnakeLifeSpan = 42
	s.ActiveExtension = "Personnalisee (vous pouvez modifier la configuration dans le fichier config.json ou choisir une autre extension en cliquant sur la fleche a la droite de votre ecran. Appuyez sur espace pour pauser)"

	for pNum := 0; pNum < config.General.InitNumParticles; pNum++ {

		if config.General.Snake {
			s.Content = append(s.Content, snakeParticule(&s))
		} else  {
			s.addParticle()
		}
	}
	return s
}

// addParticle est une méthode qui ajoute une particule à un système de
// particules.
func (s *System) addParticle() {
	s.Content = append(s.Content, newParticle(s))
}

// changeDVDColor est une fonction qui change la couleur de la particule
// lorsque celle-ci touche le bord de l'écran et que l'extension DVD est activée
func (s *System) changeDVDColor(particle *Particle) {
	if dvdColors[s.DVDColorsIndex] != nil {
		particle.ColorRed, particle.ColorGreen, particle.ColorBlue = dvdColors[s.DVDColorsIndex][0] / 255, dvdColors[s.DVDColorsIndex][1] / 255, dvdColors[s.DVDColorsIndex][2] / 255
	}

	s.DVDColorsIndex++

	if s.DVDColorsIndex >= 7 {
		s.DVDColorsIndex = 0
	}
}

// snakeDeath est une méthode qui arrête le snake lorsqu'une particule touche
// le bord de l'écran ou touche une autre particule et que l'extension Snake est activée
func (s *System) snakeDeath() {
	config.General.Snake = false
	config.General.SpawnRate = 0
	s.SnakeXDirection, s.SnakeYDirection = 0, 0

	for i := 0; i < len(s.Content); i++ {
		if s.Content[i].ColorBlue < 0.5 {
			s.Content[i].Dead = true
		} else {
			s.Content[i].ColorRed, s.Content[i].ColorGreen, s.Content[i].ColorBlue = 0.59, 0.04, 0
			s.Content[i].LifeSpan = int(math.Inf(1))
		}
	}
}

// Update met à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
func (s *System) Update() {

	var FireSpawnX, FireSpawnY int

	if config.General.Fire {
		FireSpawnX, FireSpawnY = ebiten.CursorPosition()
		config.General.SpawnX, config.General.SpawnY = float64(FireSpawnX), float64(FireSpawnY)
	}

	if config.General.Rainbow {
		// Changer la couleur RGB de config (pour changer la couleur de base des particules lors du spawn)
		if config.General.Red >= 1 && config.General.Green <= 1 && config.General.Blue <= 0 { // Rouge
			config.General.Green += 0.01
		} else if config.General.Red >= 0 && config.General.Green >= 1 && config.General.Blue <= 0 { // Jaune
			config.General.Red -= 0.01
		} else if config.General.Red <= 0 && config.General.Green >= 1 && config.General.Blue <= 1 { // Vert
			config.General.Blue += 0.01
		} else if config.General.Red <= 0 && config.General.Green >= 0 && config.General.Blue >= 1 { // Bleu clair
			config.General.Green -= 0.01
		} else if config.General.Red <= 1 && config.General.Green <= 0 && config.General.Blue >= 1 { // Bleu foncé
			config.General.Red += 0.01
		} else if config.General.Red >= 1 && config.General.Green <= 0 && config.General.Blue >= 0 { // Violet
			config.General.Blue -= 0.01
		}
	}

	for i := 0; i < len(s.Content); i++ {

		if s.Content[i].Dead {
			continue
		}

		// Ramener les particules du feu vers le centre pour faire un effet de flamme
		if config.General.Fire {
			if s.Content[i].PositionY < float64(config.General.SpawnY - 20) {
				if s.Content[i].PositionX > float64(FireSpawnX) {
					s.Content[i].SpeedX -= 0.3
				} else {
					s.Content[i].SpeedX += 0.3
				}
			}
		}

		if config.General.CollisionsBords {
			if s.Content[i].PositionX <= 0 || s.Content[i].PositionX >= float64(config.General.WindowSizeX) - s.Content[i].ScaleX * s.ImageWidth {
				s.Content[i].SpeedX = -s.Content[i].SpeedX

				if config.General.DVD {
					s.changeDVDColor(&s.Content[i])
				}

				if config.General.Snake {
					s.snakeDeath()
				}
			}

			if s.Content[i].PositionY <= 0 || s.Content[i].PositionY >= float64(config.General.WindowSizeY) - s.Content[i].ScaleY * s.ImageHeight {
				s.Content[i].SpeedY = -s.Content[i].SpeedY

				if config.General.DVD {
					s.changeDVDColor(&s.Content[i])
				}

				// Si la particule touche le bord en bas, elle s'arrête et n'est plus update (uniquement si l'extension Corail est activée)
				if config.General.Corail {
					s.Content[i].stopCorailParticle()
					continue
				}

				if config.General.Snake {
					s.snakeDeath()
				}
			}
		}


		if config.General.CollisionsParticules {
			// Temps d'attente après une collision pour éviter de détecter plusieurs une fois la même collision avec une particule
			if s.Content[i].CollisionCooldown > 0 {
				s.Content[i].CollisionCooldown--

			} else {

				for y := 0; y < len(s.Content); y++ {

					if s.Content[i] != s.Content[y] {
						if s.Content[y].PositionX - s.Content[y].ScaleX * s.ImageWidth <= s.Content[i].PositionX &&
						s.Content[y].PositionX + s.Content[y].ScaleX * s.ImageWidth >= s.Content[i].PositionX &&
						s.Content[i].PositionY + s.Content[i].ScaleY * s.ImageHeight >= s.Content[y].PositionY &&
						s.Content[i].PositionY - s.Content[i].ScaleY * s.ImageHeight <= s.Content[y].PositionY {

							// Arrêter la particule lorsqu'elle en touche une autre par en haut et ne plus l'update
							if config.General.Corail {
								if s.Content[y].Dead {
									s.Content[i].PositionY = s.Content[y].PositionY - s.ImageHeight
									s.Content[i].stopCorailParticle()

									// Arrêter de créer des particules si elles arrivent sur l'endroit où spawn les particules
									if s.Content[i].PositionX >= float64(config.General.WindowSizeX) / 2 - 40 &&
									s.Content[i].PositionX <= float64(config.General.WindowSizeX) / 2 + 40 &&
									s.Content[i].PositionY <= 40 {
										config.General.SpawnRate = 0
									}
								}

							} else if config.General.Snake {
								if !s.Content[y].Dead {
									if s.Content[y].ColorBlue < 0.5 {
										s.Content[y].Dead = true
										s.SnakeLifeSpan += 7
										s.SnakeScore += 1
										s.Content = append(s.Content, snakeParticule(s))

									} else if s.Content[i].LifeSpan > s.SnakeLifeSpan - 1 && s.Content[y].LifeSpan < s.SnakeLifeSpan - 40 && s.Content[i].ColorBlue > 0.5 {
										s.snakeDeath()
									}
								}

							} else {
								s.Content[i].SpeedY = -s.Content[i].SpeedY
								s.Content[y].SpeedY = -s.Content[y].SpeedY

								s.Content[i].CollisionCooldown = 30
								s.Content[y].CollisionCooldown = 30
							}
						}

					} else if s.Content[i].PositionX + s.Content[i].ScaleX * s.ImageWidth >= s.Content[y].PositionX &&
					s.Content[i].PositionX + s.Content[i].ScaleX * s.ImageWidth <= s.Content[y].PositionX + s.Content[y].ScaleX * s.ImageWidth && 
					s.Content[y].PositionY - s.Content[y].ScaleY * s.ImageHeight <= s.Content[i].PositionY &&
					s.Content[y].PositionY + s.Content[y].ScaleY * s.ImageHeight >= s.Content[i].PositionY &&
					!config.General.Corail && !config.General.Snake {

						s.Content[i].SpeedX = -s.Content[i].SpeedX
						s.Content[y].SpeedX = -s.Content[y].SpeedX

						s.Content[i].CollisionCooldown = 30
						s.Content[y].CollisionCooldown = 30
					}
				}
			}
		}


		if config.General.Fireworks {
			if s.Content[i].FireworkExplosion && s.Content[i].LifeSpan > 0 {

				// Explosion du feu d'artifice
				if s.Content[i].ScaleX > 0.45 {
					s.Content[i].decreaseOpacity(0.016)

					// Ralentir la particule au fur et à mesure pour faire plus réaliste
					// s.Content[i].SpeedX, s.Content[i].SpeedY = s.Content[i].SpeedX * 0.96, s.Content[i].SpeedY * 0.96
					s.Content[i].SpeedX *= 0.96
					s.Content[i].SpeedY *= 0.96

					// Appliquer une gravité à la particule vers la fin pour faire plus réaliste
					if s.Content[i].LifeSpan < 40 {
						// s.Content[i].SpeedY = s.Content[i].SpeedY + 0.2
						s.Content[i].gravity(0.2)
					}

					s.Content = append(s.Content, fireworkExploding(20.0, s.Content[i].PositionX, s.Content[i].PositionY, 0, s.Content[i].ColorRed + 0.15, s.Content[i].ColorGreen + 0.15, s.Content[i].ColorBlue + 0.15, 8, 0.2, 1))

				// Traînée de la particule lors du feu d'artifice
				} else {
					s.Content[i].decreaseOpacity(0.125)
				}
			}

			if !s.Content[i].FireworkExplosion && s.Content[i].LifeSpan == 1 {
				var FireworkRadius float64 = 15.0 + rand.Float64() * 2

				for y := -FireworkRadius; y <= FireworkRadius; y += 1 {
					s.Content = append(s.Content, fireworkExploding(FireworkRadius, s.Content[i].PositionX, s.Content[i].PositionY, y, s.Content[i].ColorRed, s.Content[i].ColorGreen, s.Content[i].ColorBlue, 60, 0.5, 1))
					s.Content = append(s.Content, fireworkExploding(FireworkRadius, s.Content[i].PositionX, s.Content[i].PositionY, y, s.Content[i].ColorRed, s.Content[i].ColorGreen, s.Content[i].ColorBlue, 60, 0.5, -1))
				}
			}
		}

		if config.General.Snake {
			config.General.SpawnX += s.SnakeXDirection
			config.General.SpawnY += s.SnakeYDirection
		}

		s.Content[i].update()
	}

	s.NumToSpawn += config.General.SpawnRate
	for s.NumToSpawn >= 1 {
		s.addParticle()
		s.NumToSpawn--
	}
}