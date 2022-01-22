package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"strconv"
	"image/color"
	"golang.org/x/image/font/basicfont"
)

// resetSystem est une méthode qui réinitialise le système (vide le Content...)
func (g *game) resetSystem() {
	assets.Get()
	g.system = particles.NewSystem()
}

// extensionChanged est une méthode qui change le nom de l'extension active
func (g *game) extensionChanged(extension string) {
	g.system.ActiveExtension = extension
}

// setConfig est une fonction qui modifie la config avec les arguments données
func setConfig(ParticleImage string, InitNumParticles int, RandomSpawn bool, SpawnX float64, SpawnY float64,
SpawnXRange float64, SpawnYRange float64, SpawnRate float64, MinXSpeed float64, MaxXSpeed float64, MinYSpeed float64,
MaxYSpeed float64, Red float64, Green float64, Blue float64, Multicolore bool, Gravite bool, ConstanteGravite float64,
RenderMarges bool, Marges float64, LifeSpan bool, LifeSpanDuration int, LifeSpanDurationRange int, CollisionsBords bool,
CollisionsParticules bool, Fire bool, Rain bool, Rainbow bool, DVD bool, Corail bool, Circle bool, Fireworks bool, Snake bool) {
	config.General.ParticleImage = ParticleImage
	config.General.InitNumParticles = InitNumParticles
	config.General.RandomSpawn = RandomSpawn
	config.General.SpawnX = SpawnX
	config.General.SpawnY = SpawnY
	config.General.SpawnXRange = SpawnXRange
	config.General.SpawnYRange = SpawnYRange
	config.General.SpawnRate = SpawnRate
	config.General.MinXSpeed = MinXSpeed
	config.General.MaxXSpeed = MaxXSpeed
	config.General.MinYSpeed = MinYSpeed
	config.General.MaxYSpeed = MaxYSpeed
	config.General.Red = Red
	config.General.Green = Green
	config.General.Blue = Blue
	config.General.Multicolore = Multicolore
	config.General.Gravite = Gravite
	config.General.ConstanteGravite = ConstanteGravite
	config.General.RenderMarges = RenderMarges
	config.General.Marges = Marges
	config.General.LifeSpan = LifeSpan
	config.General.LifeSpanDuration = LifeSpanDuration
	config.General.LifeSpanDurationRange = LifeSpanDurationRange
	config.General.CollisionsBords = CollisionsBords
	config.General.CollisionsParticules = CollisionsParticules
	config.General.Fire = Fire
	config.General.Rain = Rain
	config.General.Rainbow = Rainbow
	config.General.DVD = DVD
	config.General.Corail = Corail
	config.General.Circle = Circle
	config.General.Fireworks = Fireworks
	config.General.Snake = Snake
}

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for _, p := range g.system.Content {
		if p.Dead && !config.General.Corail {
			continue
		}

		options := ebiten.DrawImageOptions{}
		options.GeoM.Rotate(p.Rotation)
		options.GeoM.Scale(p.ScaleX, p.ScaleY)
		options.GeoM.Translate(p.PositionX, p.PositionY)
		options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
		screen.DrawImage(assets.ParticleImage, &options)
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}
	
	// Détecter un appui sur les flèches directionnels du clavier
	if config.General.Snake {
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			g.system.SnakeXDirection = -0.15
			g.system.SnakeYDirection = 0
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			g.system.SnakeXDirection = 0.15
			g.system.SnakeYDirection = 0
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			g.system.SnakeXDirection = 0
			g.system.SnakeYDirection = -0.15
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			g.system.SnakeXDirection = 0
			g.system.SnakeYDirection = 0.15
		}

		// Affiche le score en bas de l'écran
		text.Draw(screen, "Score : " + strconv.Itoa(g.system.SnakeScore), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, config.General.WindowSizeY - 13, color.RGBA{255, 255, 255, 255})
	}

	// Afficher l'rrière plan du bouton avec la flèche à droite de l'écran
	ExtensionsButtonBackgroundOptions := ebiten.DrawImageOptions{}
	ExtensionsButtonBackgroundOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 50, float64(config.General.WindowSizeY) / 2 - 125)
	screen.DrawImage(assets.ExtensionsButtonBackground, &ExtensionsButtonBackgroundOptions)

	// Afficher la flèche à droite de l'écran
	ExtensionsButtonOptions := ebiten.DrawImageOptions{}
	// Retourner la flèche lors d'un appui sur le bouton
	if g.system.ExtensionsButtonPressed {
		ExtensionsButtonOptions.GeoM.Rotate(3.14)
		ExtensionsButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 8, float64(config.General.WindowSizeY) / 2 + 33.5)
	} else {
		ExtensionsButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 42.5, float64(config.General.WindowSizeY) / 2 - 33.5)
	}
	screen.DrawImage(assets.ExtensionsButton, &ExtensionsButtonOptions)

	// Taille de l'arrière plan lorsque le menu de sélection des extensions est activé
	// Il a finalement été retiré et n'est donc pas visible
	var BackgroundSizeX = float64(config.General.WindowSizeX) * 0.8
	var BackgroundSizeY = BackgroundSizeX / 1.78


	// Appui sur le bouton gauche de la souris
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		CursorXPosition, CursorYPosition := ebiten.CursorPosition()

		// Si la souris est au dessus le bouton avec la flèche à droite de l'écran
		if float64(CursorXPosition) > float64(config.General.WindowSizeX) - 42.5 && CursorYPosition > config.General.WindowSizeY / 2 - 62 && CursorYPosition < config.General.WindowSizeY / 2 + 62 {
			g.system.ExtensionsButtonPressed = !g.system.ExtensionsButtonPressed
		}


		if g.system.ExtensionsButtonPressed {
			
			// Appui sur le bouton réinitaliser le système
			if float64(CursorXPosition) > float64(config.General.WindowSizeX) / 2 - 227.5 &&
			float64(CursorXPosition) < float64(config.General.WindowSizeX) / 2 + 227.5 &&
			float64(CursorYPosition) > float64(config.General.WindowSizeY) - 50 &&
			float64(CursorYPosition) < float64(config.General.WindowSizeY) - 10 {
				g.resetSystem()
			}

			// Appui sur le bouton pour sélectionner l'extension feu
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, false, float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY - 200), 25, 25, 1000, -5, 5, -5, 10, 1, 1, 1, false, true, 1, false, 100, true, 30, 0, false, false, true, false, false, false, false, false, false, false)
				g.resetSystem()
				g.extensionChanged("Feu")
			}

			// Appui sur le bouton pour sélectionner l'extension pluie
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, true, float64(config.General.WindowSizeX / 2), -10, float64(config.General.WindowSizeX / 2), 0, 40, 0, 0, 20, 40, 0.671, 0.804, 0.812, false, false, -1, true, 100, true, 30, 0, false, false, false, true, false, false, false, false, false, false)
				g.resetSystem()
				g.extensionChanged("Pluie")
			}
			
			// Appui sur le bouton pour sélectionner l'extension vagues multicolores
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle_outline.png", 0, true, -10, float64(config.General.WindowSizeY / 2), 0, float64(config.General.WindowSizeY / 2), 30, 5, 5, 0, 0, 1, 0, 0, false, false, -1, false, 100, false, 30, 0, false, false, false, false, true, false, false, false, false, false)
				g.resetSystem()
				g.extensionChanged("Vagues multicolore")
			}
			
			// Appui sur le bouton pour sélectionner l'extension DVD
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/dvd_logo.png", 1, true, float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY / 2), float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY / 2), 0, 5, 5, 5, 5, 1, 1, 1, false, false, -1, false, 100, false, 30, 0, true, false, false, false, false, true, false, false, false, false)
				g.resetSystem()
				g.extensionChanged("DVD")
				
			}

			// Appui sur le bouton pour sélectionner l'extension corail
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, false, float64(config.General.WindowSizeX / 2), 260, 0, 0, 5, -10, 10, -5, -10, 1, 1, 1, false, true, -0.2, false, 100, false, 60, 1, true, true, false, false, false, false, true, false, false, false)
				g.resetSystem()
				g.extensionChanged("Corail")
			}
			
			// Appui sur le bouton pour sélectionner l'extension cercle
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, false, float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY / 2), 0, 0, 200, -10, 10, -10, 10, 1, 1, 1, false, false, 0.2, false, 100, true, 40, 1, false, false, false, false, false, false, false, true, false, false)
				g.resetSystem()
				g.extensionChanged("Cercle")
			}
			
			// Appui sur le bouton pour sélectionner l'extension feu d'artifice
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, true, float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY), float64(config.General.WindowSizeX / 2), 0, 0.05, -5, 5, -10, -15, 1, 1, 1, true, false, 0.2, false, 100, true, int((0.05 * float64(config.General.WindowSizeY))), 0, false, false, false, false, false, false, false, false, true, false)
				g.resetSystem()
				g.extensionChanged("Feu d'artifice")
			}

			// Appui sur le bouton pour sélectionner l'extension snake
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 1, false, 30, float64(config.General.WindowSizeY / 2), 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, false, false, 0.2, false, 100, true, 60, 0, true, true, false, false, false, false, false, false, false, true)
				g.resetSystem()
				g.extensionChanged("Snake (Utilisez les fleches du clavier pour vous deplacer et toucher les cubes rouges")
			}
			
			// Appui sur le bouton pour sélectionner l'extension feu
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig("assets/particle.png", 0, true, float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY / 2), float64(config.General.WindowSizeX / 2), float64(config.General.WindowSizeY / 2), 1, -10, 10, -10, 10, 1, 1, 1, true, false, 0.2, true, 100, false, 60, 0, true, true, false, false, false, false, false, false, false, false)
				g.resetSystem()
				g.extensionChanged("Personnalisee")
			}
		}
	}

	// Appui sur la barre d'espace
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.system.PauseGame = !g.system.PauseGame
	}

	// Affichage du menu de sélection des extensions lors de l'appui sur la flèche à droite de l'écran
	if g.system.ExtensionsButtonPressed {

		ResetButtonOptions := ebiten.DrawImageOptions{}
		ResetButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) / 2 - 227.5, float64(config.General.WindowSizeY) - 60)
		screen.DrawImage(assets.ResetButton, &ResetButtonOptions)

		FireExtensionOptions := ebiten.DrawImageOptions{}
		FireExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		FireExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055)
		screen.DrawImage(assets.FireExtension, &FireExtensionOptions)

		RainExtensionOptions := ebiten.DrawImageOptions{}
		RainExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		RainExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055)
		screen.DrawImage(assets.RainExtension, &RainExtensionOptions)
		
		RainbowExtensionOptions := ebiten.DrawImageOptions{}
		RainbowExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		RainbowExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055)
		screen.DrawImage(assets.RainbowExtension, &RainbowExtensionOptions)
		
		DVDExtensionOptions := ebiten.DrawImageOptions{}
		DVDExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		DVDExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.DVDExtension, &DVDExtensionOptions)

		CorailExtensionOptions := ebiten.DrawImageOptions{}
		CorailExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		CorailExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.CorailExtension, &CorailExtensionOptions)
		
		CircleExtensionOptions := ebiten.DrawImageOptions{}
		CircleExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		CircleExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.CircleExtension, &CircleExtensionOptions)
		
		FireworkExtensionOptions := ebiten.DrawImageOptions{}
		FireworkExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		FireworkExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7)
		screen.DrawImage(assets.FireworkExtension, &FireworkExtensionOptions)

		SnakeExtensionOptions := ebiten.DrawImageOptions{}
		SnakeExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		SnakeExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7)
		screen.DrawImage(assets.SnakeExtension, &SnakeExtensionOptions)
		
		CustomExtensionOptions := ebiten.DrawImageOptions{}
		CustomExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		CustomExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7)
		screen.DrawImage(assets.CustomExtension, &CustomExtensionOptions)	
	}


	if ebiten.CurrentTPS() > 50 {
		text.Draw(screen, strconv.Itoa(len(g.system.Content)), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, 13, color.RGBA{0, 255, 0, 255})
	} else if ebiten.CurrentTPS() > 30 {
		text.Draw(screen, strconv.Itoa(len(g.system.Content)), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, 13, color.RGBA{255, 255, 0, 255})
	} else if ebiten.CurrentTPS() > 15 {
		text.Draw(screen, strconv.Itoa(len(g.system.Content)), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, 13, color.RGBA{255, 128, 0, 255})
	} else {
		text.Draw(screen, strconv.Itoa(len(g.system.Content)), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, 13, color.RGBA{255, 0, 0, 255})
	}

	// Affichage de l'extension en bas à gauche de l'écran
	text.Draw(screen, "Extension : " + g.system.ActiveExtension, basicfont.Face7x13, 20, config.General.WindowSizeY - 13, color.RGBA{255, 255, 255, 255})
}
