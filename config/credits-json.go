package config

type CreditsJson struct {
	Description []string `json:"description"`
	IdeaBy      []string `json:"idea_by"`
	Developers  []string `json:"developers"`
	Artists     []string `json:"artists"`
	ThanksTo    []string `json:"thanks_to"`
	Github      []string `json:"github"`
}
