package main

var cardlist map[string]cardInfo

func init() {
	cardlist = map[string]cardInfo{
		// basic
		"Forest": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"G"},
		},
		"Plain": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"W"},
		},
		"Swamp": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"B"},
		},
		"Island": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"U"},
		},
		"Mountain": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"R"},
		},

		// duals
		"Savannah": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "G"},
		},
		"Taiga": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "G"},
		},
		"Tundra": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "U"},
		},
		"Underground Sea": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"U", "B"},
		},
		"Badlands": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"B", "R"},
		},
		"Scrubland": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "B"},
		},
		"Volcanic Island": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "U"},
		},
		"Bayou": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "B"},
		},
		"Plateau": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "R"},
		},
		"Tropical Island": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "U"},
		},

		// slowland
		// too complicated for now to simuilate, leave empty for now

		// battle land
		"Prairie Stream": cardInfo{
			IsBasic:       false,
			ComesInTapped: true,
			TapsFor:       []string{"W", "U"},
		},
		"Smoldering Marsh": cardInfo{
			IsBasic:       false,
			ComesInTapped: true,
			TapsFor:       []string{"R", "B"},
		},
		"Sunken Hollow": cardInfo{
			IsBasic:       false,
			ComesInTapped: true,
			TapsFor:       []string{"B", "U"},
		},
		"Cinder Glade": cardInfo{
			IsBasic:       false,
			ComesInTapped: true,
			TapsFor:       []string{"R", "G"},
		},
		"Canopy Vista": cardInfo{
			IsBasic:       false,
			ComesInTapped: true,
			TapsFor:       []string{"G", "W"},
		},

		// painland
		"Hallowed Fountain": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "U"},
		},
		"Watery Grave": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"U", "B"},
		},
		"Blood Crypt": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"B", "R"},
		},
		"Stomping Ground": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "G"},
		},
		"Temple Garden": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "W"},
		},
		"Godless Shrine": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "B"},
		},
		"Steam Vents": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"U", "R"},
		},
		"Overgrown Tomb": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"B", "G"},
		},
		"Sacred Foundry": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "W"},
		},
		"Breeding Pool": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "U"},
		},
	}
}
