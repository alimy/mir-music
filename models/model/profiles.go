package model

// AppInfo indicator application info
type AppInfo struct {
	Info *Profiles `json:"info"`
}

// Profiles indicator service profile
type Profiles struct {
	Profile  string `json:"profiles"`
	Services string `json:"services"`
}
