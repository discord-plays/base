package config

type StatusJson struct {
	Presence  string `json:"presence"`
	Activity  string `json:"activity"`
	Reloading string `json:"reloading"`
}
