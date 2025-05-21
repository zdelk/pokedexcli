package pokeapi

// // Pokemon Info
// type PokeInfo struct {

// }
// Pokemon stats
type PokeStats struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy any    `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices            []any  `json:"game_indices"`
	Height                 int    `json:"height"`
	HeldItems              []any  `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      any    `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        any    `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault any `json:"front_default"`
				FrontFemale  any `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      any `json:"back_default"`
					BackGray         any `json:"back_gray"`
					BackTransparent  any `json:"back_transparent"`
					FrontDefault     any `json:"front_default"`
					FrontGray        any `json:"front_gray"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      any `json:"back_default"`
					BackGray         any `json:"back_gray"`
					BackTransparent  any `json:"back_transparent"`
					FrontDefault     any `json:"front_default"`
					FrontGray        any `json:"front_gray"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           any `json:"back_default"`
					BackShiny             any `json:"back_shiny"`
					BackShinyTransparent  any `json:"back_shiny_transparent"`
					BackTransparent       any `json:"back_transparent"`
					FrontDefault          any `json:"front_default"`
					FrontShiny            any `json:"front_shiny"`
					FrontShinyTransparent any `json:"front_shiny_transparent"`
					FrontTransparent      any `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      any `json:"back_default"`
					BackShiny        any `json:"back_shiny"`
					FrontDefault     any `json:"front_default"`
					FrontShiny       any `json:"front_shiny"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      any `json:"back_default"`
					BackShiny        any `json:"back_shiny"`
					FrontDefault     any `json:"front_default"`
					FrontShiny       any `json:"front_shiny"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  any `json:"back_default"`
					BackShiny    any `json:"back_shiny"`
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  any `json:"back_default"`
					BackShiny    any `json:"back_shiny"`
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      any `json:"back_default"`
						BackFemale       any `json:"back_female"`
						BackShiny        any `json:"back_shiny"`
						BackShinyFemale  any `json:"back_shiny_female"`
						FrontDefault     any `json:"front_default"`
						FrontFemale      any `json:"front_female"`
						FrontShiny       any `json:"front_shiny"`
						FrontShinyFemale any `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault any `json:"front_default"`
					FrontFemale  any `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault any `json:"front_default"`
					FrontFemale  any `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
