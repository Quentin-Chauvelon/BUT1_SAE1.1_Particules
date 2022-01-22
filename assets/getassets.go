package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"project-particles/config"
)

// ParticleImage est une variable globale pour stocker l'image d'une particule
var ParticleImage, ExtensionsButtonBackground, ExtensionsButton, ExtensionsBackground,
ResetButton, FireExtension, RainExtension, RainbowExtension, DVDExtension, CorailExtension,
CircleExtension, FireworkExtension, SnakeExtension, CustomExtension *ebiten.Image

// Get charge en mémoire l'image de la particule. (particle.png)
// Vous pouvez changer cette image si vous le souhaitez, et même en proposer
// plusieurs, qu'on peut choisir via le fichier de configuration. Cependant
// ceci n'est pas du tout central dans le projet et ne devrait être fait que
// si vous avez déjà bien avancé sur tout le reste.
func Get() {

	if config.General.Rainbow || config.General.Circle {
		config.General.ParticleImage = "assets/particle_outline.png"
	}

	if config.General.DVD {
		config.General.ParticleImage = "assets/dvd_logo.png"
	}

	var err error
	ParticleImage, _, err = ebitenutil.NewImageFromFile(config.General.ParticleImage)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}

	ExtensionsButtonBackground, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Button_Background.png")
	if err != nil {
		log.Fatal("Problem while loading extension button background image: ", err)
	}

	ExtensionsButton, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Button.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	ExtensionsBackground, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Background.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	ResetButton, _, err = ebitenutil.NewImageFromFile("assets/Reset_Button.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	FireExtension, _, err = ebitenutil.NewImageFromFile("assets/Fire_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	RainExtension, _, err = ebitenutil.NewImageFromFile("assets/Rain_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	RainbowExtension, _, err = ebitenutil.NewImageFromFile("assets/Rainbow_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	DVDExtension, _, err = ebitenutil.NewImageFromFile("assets/DVD_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	CorailExtension, _, err = ebitenutil.NewImageFromFile("assets/Corail_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	CircleExtension, _, err = ebitenutil.NewImageFromFile("assets/Circle_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	FireworkExtension, _, err = ebitenutil.NewImageFromFile("assets/Firework_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	SnakeExtension, _, err = ebitenutil.NewImageFromFile("assets/Snake_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	CustomExtension, _, err = ebitenutil.NewImageFromFile("assets/Custom_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}
}