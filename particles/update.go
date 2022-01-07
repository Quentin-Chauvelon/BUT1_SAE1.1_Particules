package particles

import (
	"project-particles/config"
)

var	CallsBeforeSpawn float64 = 0

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.	
func (s *System) Update() {

	// Mettre à jour la postion des particules en fonction de la vitesse
	for i := 0; i < len(s.Content); i++ {
		s.Content[i].PositionX += s.Content[i].SpeedX
		s.Content[i].PositionY -= s.Content[i].SpeedY
	}

	if config.General.SpawnRate > 0 {

		// Si spawn rate est inférieur à 1, on compte le nombre d'appels de la fonction et on ajoute la valeur de spawn rate à chaque fois jusqu'à ce que ce soit égal à 1
		// Exemple : pour une spawn rate de 0,1, il faut appeler 10 fois la fonction pour créer une particule (0.1 * 10 = 1)
		if config.General.SpawnRate < 1 {
			CallsBeforeSpawn += config.General.SpawnRate

			// Créer une particule et réinitialiser le compteur
			if CallsBeforeSpawn >= 0.9999 {
				CallsBeforeSpawn = 0

				var particle Particle
				s.Content = append(s.Content, particle.Add())
			}

		// Si spawn rate est supérieur à 1, on ajoute le nombre de particules correspondant
		} else {
			for i := 0; i < int(config.General.SpawnRate); i++ {
				var particle Particle
				s.Content = append(s.Content, particle.Add())
			}
		}
	}
}