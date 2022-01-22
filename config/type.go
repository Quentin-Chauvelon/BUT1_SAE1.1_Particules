package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           float64
	SpawnXRange, SpawnYRange float64
	SpawnRate                float64
	MinXSpeed				 float64
	MaxXSpeed				 float64
	MinYSpeed				 float64
	MaxYSpeed				 float64

	Red						 float64
	Green					 float64
	Blue					 float64

	Multicolore				 bool

	Gravite 				 bool
	ConstanteGravite		 float64

	RenderMarges			 bool
	Marges					 float64

	LifeSpan				 bool
	LifeSpanDuration		 int
	LifeSpanDurationRange	 int

	CollisionsBords			 bool
	CollisionsParticules	 bool

	Fire					 bool
	Rain					 bool
	Rainbow					 bool
	DVD						 bool

	Corail					 bool
	BottomColorRed			 float64
	BottomColorGreen		 float64
	BottomColorBlue			 float64

	Circle					 bool
	Fireworks				 bool
	Snake					 bool
}

var General Config