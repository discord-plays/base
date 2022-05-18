package config

type GameJson struct {
	ProjectName string `json:"project_name"`
	PlayAddress string `json:"play_address"`
	LogoAddress string `json:"logo_address"`
	Website     string `json:"website"`
	Github      string `json:"github"`
}
