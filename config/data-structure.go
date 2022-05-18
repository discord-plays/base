package config

type DataStructure struct {
	Game    *GameJson
	Status  *StatusJson
	Credits *CreditsJson
}

func LoadBotData(loadJsonFile func(name string, out any) error) (*DataStructure, error) {
	var gameConfig GameJson
	var statusConfig StatusJson
	var creditsConfig CreditsJson
	err := loadJsonFile("game", &gameConfig)
	if err != nil {
		return nil, err
	}
	err = loadJsonFile("status", &statusConfig)
	if err != nil {
		return nil, err
	}
	err = loadJsonFile("credits", &creditsConfig)
	if err != nil {
		return nil, err
	}

	return &DataStructure{
		Game:    &gameConfig,
		Status:  &statusConfig,
		Credits: &creditsConfig,
	}, nil
}
